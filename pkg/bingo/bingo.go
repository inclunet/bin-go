package bingo

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/server"
)

type Bingo struct {
	Rounds []Round
}

func (b *Bingo) AddCardsHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.AddCard()

	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("new card cannot be added"))
	}

	main, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	server.Logger.Info("Added Bingo Card", "round", card.Round, "card", card.Card, "checked", card.CheckDrawedNumbers(main))

	return server.NewResponse(card)
}

func (b *Bingo) AddRoundsHandler(r *http.Request) (*server.Response, error) {
	newRound := NewRound(b, GetUrlIntParam(r, "type"))

	b.Rounds = append(b.Rounds, newRound)

	round, err := b.GetRound(newRound.Round - 1)

	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("round cannot be added"))
	}

	old, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err == nil {
		server.Logger.Info("Redirect players", "from", old.Round, "to", round.Round, "players", old.SetNextRound(round.Round))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	server.Logger.Info("Added Bingo Round", "round", round.Round, "card", card.Card)

	return server.NewResponse(card)
}

func (b *Bingo) DrawHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	number := card.Draw()

	server.Logger.Info("Draw", "round", card.Round, "card", card.Card, "number", number, "checked", round.CheckNumberForAll(number))

	return server.NewResponse(card)
}

func (b *Bingo) GetCardsHandler(r *http.Request) (*server.Response, error) {
	Round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := Round.GetCard(GetUrlIntParam(r, "card") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	server.Logger.Info("Get Bingo Card", "round", card.Round, "card", card.Card, "checked", card.Checked)

	return server.NewResponse(card)
}

func (b *Bingo) GetCardsQRHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	qr := server.NewQRCode(fmt.Sprintf("https://%s/api/bingo/%d/new", r.Host, round.Round))

	server.Logger.Info("Get Bingo QR", "round", round.Round, "qr", qr.Content)

	return server.NewResponse(qr)
}

func (b *Bingo) GetRound(round int) (*Round, error) {
	if round < 0 || round >= len(b.Rounds) || len(b.Rounds) == 0 {
		return nil, fmt.Errorf("Round %d not found", round)
	}

	return &b.Rounds[round], nil
}

func (b *Bingo) GetRoundsHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	server.Logger.Info("Get Bingo Round", "round", round.Round, "cards", len(round.Cards))

	return server.NewResponse(round)
}

func (b *Bingo) LiveHandler(w http.ResponseWriter, r *http.Request) {
	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		response, err := server.NewResponseError(http.StatusNotFound, fmt.Errorf("round not found"))
		server.Logger.Error(err.Error())
		response.SendHasJson(w)
		return
	}

	card, err := round.GetCard(GetUrlIntParam(r, "card") - 1)

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

func (b *Bingo) ToggleCardsAutoplayHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	card, err := round.GetCard(GetUrlIntParam(r, "card") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	card.ToggleAutoplay()

	server.Logger.Info("Toggle Autoplay", "round", card.Round, "card", card.Card, "checked", card.Checked, "autoplay", card.Autoplay, "bingo", card.Bingo)

	return server.NewResponse(card)
}

func (b *Bingo) ToggleNumbersHandler(r *http.Request) (*server.Response, error) {
	round, err := b.GetRound(GetUrlIntParam(r, "round") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	mainCard, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	card, err := round.GetCard(GetUrlIntParam(r, "card") - 1)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	if card.ToggleNumber(mainCard, GetUrlIntParam(r, "number")) && card.Card > 1 {
		server.Logger.Info("Toggle Number", "round", card.Round, "card", card.Card, "number", GetUrlIntParam(r, "number"), "checked", card.Checked, "bingo", card.Bingo)
	}

	if card.Card == 1 {
		checked, unchecked := round.ToggleNumberForAll(GetUrlIntParam(r, "number"))

		server.Logger.Info("Toggle Number", "round", card.Round, "card", card.Card, "number", GetUrlIntParam(r, "number"), "checked", checked, "unchecked", unchecked)
	}

	return server.NewResponse(card)
}

func GetUrlIntParam(r *http.Request, param string) int {
	if intParam, err := strconv.Atoi(GetUrlStringParam(r, param)); err == nil {
		return intParam
	}

	return 0
}

func GetQueryString(r *http.Request, key string, value string) string {
	query := r.URL.Query()

	if value, ok := query[key]; ok {
		return value[0]
	}

	return value
}

func GetUrlStringParam(r *http.Request, param string) string {
	vars := mux.Vars(r)
	return vars[param]
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
		r.Methods(http.MethodGet).Path("/{round}/{card}/autoplay").Handler(server.SendJson(b.ToggleCardsAutoplayHandler))
		r.Methods(http.MethodGet).Path("/{round}/1/0").Handler(server.SendJson(b.DrawHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/{number}").Handler(server.SendJson(b.ToggleNumbersHandler))
	}

	return b
}
