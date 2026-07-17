package bingo

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type fakeStore struct {
	rounds []*Round
	saved  int
}

func (s *fakeStore) LoadRounds(context.Context) ([]*Round, error) {
	return s.rounds, nil
}

func (s *fakeStore) SaveRound(_ context.Context, round *Round) error {
	s.saved++
	s.rounds = []*Round{round}
	return nil
}

func (s *fakeStore) SaveRounds(ctx context.Context, rounds ...*Round) error {
	for _, round := range rounds {
		if err := s.SaveRound(ctx, round); err != nil {
			return err
		}
	}
	return nil
}

func TestRoundAndCardsUseUUIDs(t *testing.T) {
	round := NewRound(&Bingo{}, 75)
	if _, err := uuid.Parse(round.ID); err != nil {
		t.Fatalf("round ID is not a UUID: %v", err)
	}

	main, err := round.GetCard(0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := uuid.Parse(main.ID); err != nil {
		t.Fatalf("main card ID is not a UUID: %v", err)
	}
	if main.RoundID != round.ID {
		t.Fatalf("main card round ID = %q, want %q", main.RoundID, round.ID)
	}

	player, err := round.AddCard()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := uuid.Parse(player.ID); err != nil {
		t.Fatalf("player card ID is not a UUID: %v", err)
	}
	if player.ID == main.ID {
		t.Fatal("main and player cards have the same ID")
	}
}

func TestNilPoolCreatesNilStore(t *testing.T) {
	if store := NewPostgresStore(nil); store != nil {
		t.Fatal("nil pool created a non-nil store")
	}
}

func TestNewRoundUsesNextHighestDisplayNumber(t *testing.T) {
	game := &Bingo{Rounds: []*Round{{Round: 1}, {Round: 2}, {Round: 4}}}

	round := NewRound(game, 75)

	if round.Round != 5 {
		t.Fatalf("round display number = %d, want 5", round.Round)
	}
}

func TestAddingCardRelinksMainCard(t *testing.T) {
	round := NewRound(&Bingo{}, 75)

	for i := 0; i < 20; i++ {
		if _, err := round.AddCard(); err != nil {
			t.Fatal(err)
		}
	}

	main := &round.Cards[0]
	for i := 1; i < len(round.Cards); i++ {
		if round.Cards[i].Main != main {
			t.Fatalf("card %d is not linked to the current main card", i+1)
		}
	}
}

func TestRoundMutationRestoresFailedDraw(t *testing.T) {
	round := NewRound(&Bingo{}, 75)
	player, err := round.AddCard()
	if err != nil {
		t.Fatal(err)
	}
	player.PlayerID = uuid.NewString()
	playerID := player.PlayerID

	mutation, err := round.BeginMutation()
	if err != nil {
		t.Fatal(err)
	}
	number := round.Draw().LastNumber
	round.ToggleNumberForAll(number)

	if err := mutation.Restore(&round); err != nil {
		t.Fatal(err)
	}
	if round.Cards[0].Checked != 0 {
		t.Fatalf("main card kept %d checked numbers after rollback", round.Cards[0].Checked)
	}
	if round.Cards[0].LastNumber != 0 {
		t.Fatalf("main card kept last number %d after rollback", round.Cards[0].LastNumber)
	}
	if round.Cards[1].PlayerID != playerID {
		t.Fatal("rollback discarded anonymous player ID")
	}
}

func TestCompletedRoundDoesNotToggleFreeSpace(t *testing.T) {
	game, err := NewWithStore(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	roundValue := NewRound(game, 75)
	round := &roundValue
	player, err := round.AddCard()
	if err != nil {
		t.Fatal(err)
	}
	round.Cards[0].Checked = round.Type
	game.Rounds = append(game.Rounds, round)
	game.roundsByID[round.ID] = round
	game.roundLocks[round.ID] = &sync.Mutex{}

	main := &round.Cards[0]
	request := httptest.NewRequest("GET", "/", nil)
	request = mux.SetURLVars(request, map[string]string{
		"round": round.ID,
		"card":  main.ID,
	})

	if _, err := game.DrawHandler(request); err != nil {
		t.Fatal(err)
	}
	if !player.Numbers[2][2].Checked {
		t.Fatal("completed draw unchecked the player's free space")
	}
}

func TestCheckWebSocketOrigin(t *testing.T) {
	tests := []struct {
		name       string
		host       string
		origin     string
		forwarded  string
		acceptable bool
	}{
		{name: "missing origin", host: "example.com", acceptable: true},
		{name: "same origin", host: "example.com", origin: "http://example.com", acceptable: true},
		{name: "implicit HTTPS port", host: "example.com:443", origin: "https://example.com", forwarded: "https", acceptable: true},
		{name: "explicit HTTPS port", host: "example.com", origin: "https://example.com:443", forwarded: "https", acceptable: true},
		{name: "HTTPS proxy without scheme header", host: "example.com", origin: "https://example.com", acceptable: true},
		{name: "different default scheme", host: "example.com:80", origin: "https://example.com", forwarded: "https", acceptable: false},
		{name: "local proxy", host: "localhost:8080", origin: "http://localhost:5173", acceptable: true},
		{name: "local IP proxy", host: "127.0.0.1:8080", origin: "http://127.0.0.1:5173", acceptable: true},
		{name: "different origin", host: "example.com", origin: "https://attacker.example", acceptable: false},
		{name: "local origin against production", host: "example.com", origin: "http://localhost:5173", acceptable: false},
		{name: "non-HTTP origin", host: "localhost:8080", origin: "ftp://localhost:5173", acceptable: false},
		{name: "malformed origin", host: "example.com", origin: "://invalid", acceptable: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "http://"+test.host+"/ws", nil)
			request.Host = test.host
			if test.forwarded != "" {
				request.Header.Set("X-Forwarded-Proto", test.forwarded)
			}
			if test.origin != "" {
				request.Header.Set("Origin", test.origin)
			}

			if actual := checkWebSocketOrigin(request); actual != test.acceptable {
				t.Fatalf("checkWebSocketOrigin() = %v, want %v", actual, test.acceptable)
			}
		})
	}
}

func TestRequestScheme(t *testing.T) {
	tests := []struct {
		name      string
		target    string
		forwarded string
		expected  string
	}{
		{name: "plain HTTP", target: "http://example.com/qr", expected: "http"},
		{name: "direct HTTPS", target: "https://example.com/qr", expected: "https"},
		{name: "HTTPS proxy", target: "http://example.com/qr", forwarded: "https", expected: "https"},
		{name: "first proxy value", target: "http://example.com/qr", forwarded: "https, http", expected: "https"},
		{name: "invalid proxy value", target: "http://example.com/qr", forwarded: "ftp", expected: "http"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", test.target, nil)
			if test.forwarded != "" {
				request.Header.Set("X-Forwarded-Proto", test.forwarded)
			}

			if actual := requestScheme(request); actual != test.expected {
				t.Fatalf("requestScheme() = %q, want %q", actual, test.expected)
			}
		})
	}
}

func TestAddRoundsHandlerReturnsExistingNextRound(t *testing.T) {
	game := New(nil)

	initialRequest := mux.SetURLVars(
		httptest.NewRequest("GET", "/api/bingo/0/new/75", nil),
		map[string]string{"round": "0", "type": "75"},
	)
	initialRequest.Header.Set("X-Bingo-Creation-ID", uuid.NewString())
	initialResponse, err := game.AddRoundsHandler(initialRequest)
	if err != nil {
		t.Fatal(err)
	}
	initialCard := initialResponse.Body.(*Card)

	newRoundVars := map[string]string{
		"round": initialCard.RoundID,
		"card":  initialCard.ID,
		"type":  "75",
	}
	newRoundRequest := mux.SetURLVars(
		httptest.NewRequest("GET", "/api/bingo/round/card/new/75", nil),
		newRoundVars,
	)
	newRoundResponse, err := game.AddRoundsHandler(newRoundRequest)
	if err != nil {
		t.Fatal(err)
	}
	newRoundCard := newRoundResponse.Body.(*Card)

	retryRequest := mux.SetURLVars(
		httptest.NewRequest("GET", "/api/bingo/round/card/new/75", nil),
		newRoundVars,
	)
	retryResponse, err := game.AddRoundsHandler(retryRequest)
	if err != nil {
		t.Fatal(err)
	}
	retryCard := retryResponse.Body.(*Card)

	if retryCard.ID != newRoundCard.ID || retryCard.RoundID != newRoundCard.RoundID {
		t.Fatalf("retry returned card %s/%s, want %s/%s", retryCard.RoundID, retryCard.ID, newRoundCard.RoundID, newRoundCard.ID)
	}
	if len(game.Rounds) != 2 {
		t.Fatalf("retry created %d rounds, want 2", len(game.Rounds))
	}
}

func TestAddRoundsHandlerReturnsExistingInitialRound(t *testing.T) {
	game := New(nil)
	creationID := uuid.NewString()

	createRound := func() *Card {
		t.Helper()
		request := mux.SetURLVars(
			httptest.NewRequest("GET", "/api/bingo/0/new/75", nil),
			map[string]string{"round": "0", "type": "75"},
		)
		request.Header.Set("X-Bingo-Creation-ID", creationID)
		response, err := game.AddRoundsHandler(request)
		if err != nil {
			t.Fatal(err)
		}
		return response.Body.(*Card)
	}

	first := createRound()
	second := createRound()
	if first.RoundID != second.RoundID || first.ID != second.ID {
		t.Fatal("same creation request produced different rounds")
	}
	if len(game.Rounds) != 1 {
		t.Fatalf("game has %d rounds, want 1", len(game.Rounds))
	}
}

func TestAddRoundsHandlerRejectsUnknownParentRound(t *testing.T) {
	game := New(nil)
	request := mux.SetURLVars(
		httptest.NewRequest("GET", "/api/bingo/missing/card/new/75", nil),
		map[string]string{
			"round": uuid.NewString(),
			"card":  uuid.NewString(),
			"type":  "75",
		},
	)
	if _, err := game.AddRoundsHandler(request); err == nil {
		t.Fatal("unknown parent round created an orphan round")
	}
	if len(game.Rounds) != 0 {
		t.Fatalf("game has %d orphan rounds", len(game.Rounds))
	}
}

func TestAddCardsHandlerReturnsExistingCardForPlayer(t *testing.T) {
	game := New(nil)
	roundRequest := mux.SetURLVars(
		httptest.NewRequest("GET", "/api/bingo/0/new/75", nil),
		map[string]string{"round": "0", "type": "75"},
	)
	roundRequest.Header.Set("X-Bingo-Creation-ID", uuid.NewString())
	roundResponse, err := game.AddRoundsHandler(roundRequest)
	if err != nil {
		t.Fatal(err)
	}
	roundID := roundResponse.Body.(*Card).RoundID
	playerID := uuid.NewString()

	addCard := func() *Card {
		t.Helper()
		request := mux.SetURLVars(
			httptest.NewRequest("GET", "/api/bingo/round/0", nil),
			map[string]string{"round": roundID},
		)
		request.Header.Set("X-Bingo-Player-ID", playerID)
		response, err := game.AddCardsHandler(request)
		if err != nil {
			t.Fatal(err)
		}
		return response.Body.(*Card)
	}

	first := addCard()
	second := addCard()
	if first.ID != second.ID {
		t.Fatalf("same player received cards %s and %s", first.ID, second.ID)
	}
	round, err := game.GetRoundByID(roundID)
	if err != nil {
		t.Fatal(err)
	}
	if len(round.Cards) != 2 {
		t.Fatalf("round has %d cards, want organizer and one player", len(round.Cards))
	}
}

func TestQueueUpdateKeepsOnlyLatestPendingSend(t *testing.T) {
	card := Card{runtime: &cardRuntime{}}
	started := make(chan struct{})
	release := make(chan struct{})
	done := make(chan struct{})
	var executedMu sync.Mutex
	executed := []int{}

	card.QueueUpdate(func() error {
		executedMu.Lock()
		executed = append(executed, 0)
		executedMu.Unlock()
		close(started)
		<-release
		return nil
	})
	<-started

	for i := 1; i <= 10; i++ {
		value := i
		card.QueueUpdate(func() error {
			executedMu.Lock()
			executed = append(executed, value)
			executedMu.Unlock()
			if value == 10 {
				close(done)
			}
			return nil
		})
	}
	close(release)

	select {
	case <-done:
	case <-time.After(time.Second):
		t.Fatal("latest queued update was not sent")
	}

	executedMu.Lock()
	defer executedMu.Unlock()
	if len(executed) != 2 || executed[0] != 0 || executed[1] != 10 {
		t.Fatalf("executed updates %v, want [0 10]", executed)
	}
}

func TestUpdateCardWithoutConnectionDoesNotStartWriter(t *testing.T) {
	card := Card{runtime: &cardRuntime{}}
	send, err := card.PrepareUpdate()
	if err != nil {
		t.Fatal(err)
	}
	if send != nil {
		t.Fatal("PrepareUpdate returned a send function without a connection")
	}
	if err := card.UpdateCard(); err != nil {
		t.Fatal(err)
	}

	card.runtime.queueMu.Lock()
	defer card.runtime.queueMu.Unlock()
	if card.runtime.writerRunning || card.runtime.pendingSend != nil {
		t.Fatal("UpdateCard started a writer without a connection")
	}
}

func TestNewCardResponseSnapshotsMutableState(t *testing.T) {
	round := NewRound(&Bingo{}, 75)
	card, err := round.GetCard(0)
	if err != nil {
		t.Fatal(err)
	}

	response, err := newCardResponse(card)
	if err != nil {
		t.Fatal(err)
	}
	snapshot := response.Body.(*Card)
	originalRound := snapshot.Round
	originalChecked := snapshot.Numbers[0][0].Checked

	card.Round++
	card.Numbers[0][0].Checked = !card.Numbers[0][0].Checked

	if snapshot.Round != originalRound {
		t.Fatalf("snapshot round changed to %d, want %d", snapshot.Round, originalRound)
	}
	if snapshot.Numbers[0][0].Checked != originalChecked {
		t.Fatal("snapshot number changed with live card")
	}
}

func TestCardJSONDoesNotPersistRuntimeMainPointer(t *testing.T) {
	round := NewRound(&Bingo{}, 75)
	card, err := round.AddCard()
	if err != nil {
		t.Fatal(err)
	}

	encoded, err := json.Marshal(card)
	if err != nil {
		t.Fatal(err)
	}

	var state map[string]interface{}
	if err := json.Unmarshal(encoded, &state); err != nil {
		t.Fatal(err)
	}
	if _, exists := state["Main"]; exists {
		t.Fatal("runtime Main pointer was serialized")
	}
	if _, exists := state["PlayerID"]; exists {
		t.Fatal("anonymous player ID was serialized")
	}
}

func TestNewWithStoreRestoresPersistedRounds(t *testing.T) {
	round := NewRound(&Bingo{}, 75)
	if _, err := round.AddCard(); err != nil {
		t.Fatal(err)
	}

	store := &fakeStore{rounds: []*Round{&round}}
	game, err := NewWithStore(nil, store)
	if err != nil {
		t.Fatal(err)
	}

	loaded, err := game.GetRoundByID(round.ID)
	if err != nil {
		t.Fatal(err)
	}
	if len(loaded.Cards) != 2 {
		t.Fatalf("loaded %d cards, want 2", len(loaded.Cards))
	}
}
