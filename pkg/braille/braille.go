package braille

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Braille struct {
	Players []BrailleClass
}

func (b *Braille) AddPlayer() BrailleClass {
	player := NewBrailleClass(len(b.Players))
	b.Players = append(b.Players, player)
	return player
}

func (b *Braille) AddPlayerHandler(w http.ResponseWriter, r *http.Request) {
	Repply(w, b.AddPlayer())
}

func (b *Braille) CheckChallengeRepplyHandler(w http.ResponseWriter, r *http.Request) {
	player := b.GetPlayer(GetUrlIntParam(r, "player", -1))

	if player == nil {
		http.Error(w, fmt.Sprintf("Player %d not found", GetUrlIntParam(r, "player", -1)), http.StatusNotFound)
		return
	}
	fmt.Println("antes do oi")
	var challengeRepply BrailleClass

	err := json.NewDecoder(r.Body).Decode(&challengeRepply)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("oi")
	Repply(w, player.Check(challengeRepply.Challenge.Repply))
}

func (b *Braille) GetPlayer(i int) *BrailleClass {
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

	var repply BrailleClass

	err := json.NewDecoder(r.Body).Decode(&repply)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	Repply(w, b)
}

func New(routes *mux.Router) (b *Braille, err error) {
	err = LoadClass("classes.json")

	if err != nil {
		return b, err
	}

	b = &Braille{
		Players: []BrailleClass{},
	}

	if routes != nil {
		r := routes.PathPrefix("/braille").Subrouter()
		r.Methods(http.MethodGet).Path("/new").HandlerFunc(b.AddPlayerHandler)
		r.Methods(http.MethodGet).Path("/{player}").HandlerFunc(b.GetPlayerHandler)
		r.Methods(http.MethodPost).Path("/{player}").HandlerFunc(b.CheckChallengeRepplyHandler)
	}

	return b, nil
}
