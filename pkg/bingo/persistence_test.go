package bingo

import (
	"context"
	"encoding/json"
	"net/http/httptest"
	"sync"
	"testing"

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
	if _, err := round.AddCard(); err != nil {
		t.Fatal(err)
	}

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
		acceptable bool
	}{
		{name: "missing origin", host: "example.com", acceptable: true},
		{name: "same origin", host: "example.com", origin: "https://example.com", acceptable: true},
		{name: "local proxy", host: "localhost:8080", origin: "http://localhost:5173", acceptable: true},
		{name: "local IP proxy", host: "127.0.0.1:8080", origin: "http://127.0.0.1:5173", acceptable: true},
		{name: "different origin", host: "example.com", origin: "https://attacker.example", acceptable: false},
		{name: "local origin against production", host: "example.com", origin: "http://localhost:5173", acceptable: false},
		{name: "malformed origin", host: "example.com", origin: "://invalid", acceptable: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest("GET", "http://"+test.host+"/ws", nil)
			request.Host = test.host
			if test.origin != "" {
				request.Header.Set("Origin", test.origin)
			}

			if actual := checkWebSocketOrigin(request); actual != test.acceptable {
				t.Fatalf("checkWebSocketOrigin() = %v, want %v", actual, test.acceptable)
			}
		})
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
