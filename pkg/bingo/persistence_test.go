package bingo

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/google/uuid"
)

type fakeStore struct {
	rounds []Round
	saved  int
}

func (s *fakeStore) LoadRounds(context.Context) ([]Round, error) {
	return s.rounds, nil
}

func (s *fakeStore) SaveRound(_ context.Context, round *Round) error {
	s.saved++
	s.rounds = []Round{*round}
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

	store := &fakeStore{rounds: []Round{round}}
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
