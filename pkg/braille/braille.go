package braille

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/server"
)

type Braille struct {
	Players []BrailleClass
}

func (b *Braille) AddPlayerHandler(r *http.Request) (*server.Response, error) {
	player := NewBrailleClass(len(b.Players))
	b.Players = append(b.Players, player)

	return server.NewResponse(player)
}

func (b *Braille) CheckChallengeRepplyHandler(r *http.Request) (*server.Response, error) {
	player := b.GetPlayer(GetUrlIntParam(r, "player", -1))

	if player == nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("player not found"))
	}

	var repply BrailleClass

	err := json.NewDecoder(r.Body).Decode(&repply)

	if err != nil {
		return server.NewResponseError(http.StatusBadRequest, err)
	}

	return server.NewResponse(player.Check(repply.Challenge.Repply))
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

func (b *Braille) GetPlayerHandler(r *http.Request) (*server.Response, error) {
	player := b.GetPlayer(GetUrlIntParam(r, "player", -1))

	if player == nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("player not found"))
	}

	return server.NewResponse(player)
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
		r.Methods(http.MethodGet).Path("/new").Handler(server.SendJson(b.AddPlayerHandler))
		r.Methods(http.MethodGet).Path("/{player}").Handler(server.SendJson(b.GetPlayerHandler))
		r.Methods(http.MethodPost).Path("/{player}").Handler(server.SendJson(b.CheckChallengeRepplyHandler))
	}

	return b, nil
}
