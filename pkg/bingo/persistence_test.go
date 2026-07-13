package bingo

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
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
