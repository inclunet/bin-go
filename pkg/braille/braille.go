package braille

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/inclunet/bin-go/pkg/classes"
)

type Braille struct {
	Players []classes.BrailleClass
}

func (b *Braille) AddPlayer() classes.BrailleClass {
	player := classes.NewBrailleClass(len(b.Players))
	b.Players = append(b.Players, player)
	return player
}

func (b *Braille) AddPlayerHandler(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.AddPlayer())
}

func (b *Braille) GetPlayer(i int) *classes.BrailleClass {
	if i < 0 {
		return nil
	}

	if i >= len(b.Players) {
		return nil
	}

	return &b.Players[i]
}

func (b *Braille) GetPlayerHandler(w http.ResponseWriter, r *http.Request) {
	player := b.GetPlayer(GetUrlIntParam(r, "player", -1))

	if player == nil {
		http.Error(w, fmt.Sprintf("Player %d not found", GetUrlIntParam(r, "player", -1)), http.StatusNotFound)
		return
	}

	Repply(w, player)
}

func (b *Braille) SolveChallengeHandler(w http.ResponseWriter, r *http.Request) {
	player := b.GetPlayer(GetUrlIntParam(r, "player", -1))

	if player == nil {
		http.Error(w, fmt.Sprintf("Player %d not found", GetUrlIntParam(r, "player", -1)), http.StatusNotFound)
		return
	}

	var repply classes.BrailleClass

	err := json.NewDecoder(r.Body).Decode(&repply)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Repply(w, b)
}

func New(classesFileName string) (*Braille, error) {
	err := classes.LoadClass(classesFileName)

	if err != nil {
		return nil, err
	}

	b := Braille{
		Players: []classes.BrailleClass{},
	}

	return &b, nil
}
