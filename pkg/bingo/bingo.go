package bingo

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/server"
)

type Bingo struct {
	Rounds []Round
}

func (b *Bingo) AddCardsHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.AddCard()

	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("new card cannot be added"))
	}

	b.Log("Add Bingo Card", card)

	return server.NewResponse(card)
}

func (b *Bingo) AddRoundsHandler(r *http.Request) (*server.Response, error) {
	newRound := NewRound(b, server.GetURLParamHasInt(r, "type"))

	b.Rounds = append(b.Rounds, newRound)

	round, err := b.GetRound(newRound.Round - 1)

	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("round cannot be added"))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	old, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err == nil {
		b.Log("Redirect Old Players to the New Bingo Round", card, "from", old.Round, "to", round.Round, "players", old.SetNextRoundForAll(round.Round))
	}

	b.Log("Add Bingo Round", card)

	return server.NewResponse(card)
}

func (b *Bingo) CancelAlertHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(server.GetURLParamHasInt(r, "card") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	card.CancelAlert()

	b.Log("Cancel Bingo Alert", card)

	return server.NewResponse(card)
}

func (b *Bingo) DrawHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	number := card.Draw()

	checked, Unchecked := round.ToggleNumberForAll(number)

	b.Log("Draw new Random Bingo Number", card, "number", number, "checked", checked, "unchecked", Unchecked)

	return server.NewResponse(card)
}

func (b *Bingo) GetCardsHandler(r *http.Request) (*server.Response, error) {
	Round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := Round.GetCard(server.GetURLParamHasInt(r, "card") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	b.Log("Get Bingo Card", card)

	return server.NewResponse(card)
}

func (b *Bingo) GetCardsQRHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	qr := server.NewQRCode(fmt.Sprintf("https://%s/api/bingo/%d/new", r.Host, round.Round))

	b.Log("Get Bingo Round QR", card, "qr", qr.Content)

	return server.NewResponse(qr)
}

func (b *Bingo) GetRound(round int) (*Round, error) {
	if round < 0 || round >= len(b.Rounds) || len(b.Rounds) == 0 {
		return nil, fmt.Errorf("Round %d not found", round)
	}

	return &b.Rounds[round], nil
}

func (b *Bingo) GetRoundsHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	b.Log("Get Bingo Round", card, "cards", len(round.Cards))

	return server.NewResponse(round)
}

func (b *Bingo) LiveHandler(w http.ResponseWriter, r *http.Request) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		response, err := server.NewResponseError(http.StatusNotFound, fmt.Errorf("round not found"))
		server.Logger.Error(err.Error())
		response.SendHasJson(w)
		return
	}

	card, err := round.GetCard(server.GetURLParamHasInt(r, "card") - 1)

	if err != nil {
		response, err := server.NewResponseError(http.StatusNotFound, fmt.Errorf("card not found"))
		server.Logger.Error(err.Error())
		response.SendHasJson(w)
		return
	}

	round.upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	conn, err := round.upgrader.Upgrade(w, r, nil)

	if err != nil {
		response, err := server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("websocket connection error: %v", err))
		server.Logger.Error(err.Error())
		response.SendHasJson(w)
		return
	}

	if !card.SetConn(conn) {
		response, err := server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("websocket connection error: %v", err))
		server.Logger.Error(err.Error())
		response.SendHasJson(w)
		return

	}

	card.UpdateCard()
}

func (b *Bingo) Log(msg string, card *Card, complement ...any) {
	info := []any{
		"round", card.Round,
		"card", card.Card,
		"checked", card.Checked,
		"lastnumber", card.LastNumber,
		"autoplay", card.Autoplay,
		"bingo", card.Bingo,
		"lastcompletion", card.LastCompletion,
		"type", card.Type,
	}

	if len(complement) > 1 && len(complement)%2 == 0 {
		info = append(info, complement...)
	}

	server.Logger.Info(msg, info...)
}

func (b *Bingo) SetCompletionsHandler(h *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(h, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	completions := NewDefaultCompletions()

	err = json.NewDecoder(h.Body).Decode(&completions)

	if err != nil {
		return server.NewResponseError(http.StatusBadRequest, errors.New("completions cannot be decoded"))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	counter, err := round.SetCompletionsForAll(completions)

	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, err)
	}

	b.Log("Set new Completions Values for the bingo Round", card, "counter", counter, "completions", card.Completions)

	return server.NewResponse(card)
}

func (b *Bingo) ToggleCardsAutoplayHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(server.GetURLParamHasInt(r, "card") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	card.ToggleAutoplay()

	b.Log("Toggle Bingo Card Autoplay", card)

	return server.NewResponse(card)
}

func (b *Bingo) ToggleNumbersHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(server.GetURLParamHasInt(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(server.GetURLParamHasInt(r, "card") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	if card.ToggleNumber(server.GetURLParamHasInt(r, "number")) && card.Card > 1 {
		b.Log("Toggle Bingo Card Number", card, "number", server.GetURLParamHasInt(r, "number"))
	}

	if card.Card == 1 {
		checked, unchecked := round.ToggleNumberForAll(server.GetURLParamHasInt(r, "number"))

		b.Log("Toggle Bingo Card Number for All", card, "number", server.GetURLParamHasInt(r, "number"), "checked", checked, "unchecked", unchecked)
	}

	return server.NewResponse(card)
}

func (b *Bingo) AddQrRoutes(routes *mux.Router) *Bingo {
	if routes != nil {
		r := routes.PathPrefix("/bingo").Subrouter()
		r.Methods(http.MethodGet).Path("/{round}").Handler(server.SendQRCode(b.GetCardsQRHandler))
	}

	return b
}

func (b *Bingo) AddWsRoutes(routes *mux.Router) *Bingo {
	if routes != nil {
		r := routes.PathPrefix("/bingo").Subrouter()
		r.Methods(http.MethodGet).Path("/{round}/{card}").HandlerFunc(b.LiveHandler)
	}

	return b
}

func New(routes *mux.Router) (b *Bingo) {
	b = &Bingo{
		Rounds: []Round{},
	}

	if routes != nil {
		r := routes.PathPrefix("/bingo").Subrouter()
		r.Methods(http.MethodGet).Path("/{round}/new/{type}").Handler(server.SendJson(b.AddRoundsHandler))
		r.Methods(http.MethodGet).Path("/{round}").Handler(server.SendJson(b.GetRoundsHandler))
		r.Methods(http.MethodGet).Path("/{round}/0").Handler(server.SendJson(b.AddCardsHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}").Handler(server.SendJson(b.GetCardsHandler))
		r.Methods(http.MethodPost).Path("/{round}/1/completions").Handler(server.SendJson(b.SetCompletionsHandler))
		r.Methods(http.MethodGet).Path("/{round}/1/0").Handler(server.SendJson(b.DrawHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/autoplay").Handler(server.SendJson(b.ToggleCardsAutoplayHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/cancel").Handler(server.SendJson(b.CancelAlertHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/{number}").Handler(server.SendJson(b.ToggleNumbersHandler))
	}

	return b
}
