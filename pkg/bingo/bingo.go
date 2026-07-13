package bingo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/server"
)

type Bingo struct {
	Rounds     []*Round
	store      Store
	roundsByID map[string]*Round
	roundLocks map[string]*sync.Mutex
	mu         sync.RWMutex
}

type PublicRound struct {
	ID    string
	Round int
	Type  int
}

func (b *Bingo) persistRound(ctx context.Context, round *Round) error {
	if b.store == nil {
		return nil
	}

	return b.store.SaveRound(ctx, round)
}

func (b *Bingo) persistRounds(ctx context.Context, rounds ...*Round) error {
	if b.store == nil {
		return nil
	}

	return b.store.SaveRounds(ctx, rounds...)
}

func (b *Bingo) getRoundAndLock(roundID string) (*Round, *sync.Mutex, error) {
	b.mu.RLock()
	round := b.roundsByID[roundID]
	roundLock := b.roundLocks[roundID]
	b.mu.RUnlock()

	if round == nil || roundLock == nil {
		return nil, nil, fmt.Errorf("round %s not found", roundID)
	}

	roundLock.Lock()
	return round, roundLock, nil
}

func (b *Bingo) AddCardsHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.AddCard()

	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("new card cannot be added"))
	}
	if err := b.persistRound(r.Context(), round); err != nil {
		round.Cards = round.Cards[:len(round.Cards)-1]
		round.RelinkCards()
		return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("new card cannot be saved: %w", err))
	}

	b.Log("Add Bingo Card", card)

	return server.NewResponse(card)
}

func (b *Bingo) AddRoundsHandler(r *http.Request) (*server.Response, error) {
	b.mu.Lock()
	defer b.mu.Unlock()

	newRound := NewRound(b, server.GetURLParamHasInt(r, "type"))

	round := &newRound
	b.Rounds = append(b.Rounds, round)
	b.roundsByID[round.ID] = round
	b.roundLocks[round.ID] = &sync.Mutex{}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	old := b.roundsByID[server.GetURLParam(r, "round")]

	if old != nil {
		oldLock := b.roundLocks[old.ID]
		oldLock.Lock()
		defer oldLock.Unlock()

		previousRoundID := old.NextRoundID
		previousNextRounds := make([]int, len(old.Cards))
		previousNextRoundIDs := make([]string, len(old.Cards))
		for i := range old.Cards {
			previousNextRounds[i] = old.Cards[i].NextRound
			previousNextRoundIDs[i] = old.Cards[i].NextRoundID
		}

		players := old.SetNextRoundForAll(round)
		if err := b.persistRounds(r.Context(), round, old); err != nil {
			old.NextRoundID = previousRoundID
			for i := range old.Cards {
				old.Cards[i].NextRound = previousNextRounds[i]
				old.Cards[i].NextRoundID = previousNextRoundIDs[i]
			}
			b.Rounds = b.Rounds[:len(b.Rounds)-1]
			delete(b.roundsByID, round.ID)
			delete(b.roundLocks, round.ID)
			return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("rounds cannot be saved: %w", err))
		}
		for i := range old.Cards {
			_ = old.Cards[i].UpdateCard()
		}
		b.Log("Redirect Old Players to the New Bingo Round", card, "from", old.Round, "to", round.Round, "players", players)
	} else if err := b.persistRound(r.Context(), round); err != nil {
		b.Rounds = b.Rounds[:len(b.Rounds)-1]
		delete(b.roundsByID, round.ID)
		delete(b.roundLocks, round.ID)
		return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("round cannot be saved: %w", err))
	}

	b.Log("Add Bingo Round", card)

	return server.NewResponse(card)
}

func (b *Bingo) CancelAlertHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.GetCardByID(server.GetURLParam(r, "card"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	mutation, err := round.BeginMutation()
	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, err)
	}
	card.CancelAlert()
	if err := b.persistRound(r.Context(), round); err != nil {
		_ = mutation.Restore(round)
		return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("card cannot be saved: %w", err))
	}
	mutation.Publish(round)

	b.Log("Cancel Bingo Alert", card)

	return server.NewResponse(card)
}

func (b *Bingo) DrawHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}
	if card.ID != server.GetURLParam(r, "card") {
		return server.NewResponseError(http.StatusForbidden, errors.New("only the main card can draw numbers"))
	}

	mutation, err := round.BeginMutation()
	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, err)
	}
	number := card.Draw()

	checked, Unchecked := round.ToggleNumberForAll(number)
	if err := b.persistRound(r.Context(), round); err != nil {
		_ = mutation.Restore(round)
		return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("draw cannot be saved: %w", err))
	}
	mutation.Publish(round)

	b.Log("Draw new Random Bingo Number", card, "number", number, "checked", checked, "unchecked", Unchecked)

	return server.NewResponse(card)
}

func (b *Bingo) GetCardsHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.GetCardByID(server.GetURLParam(r, "card"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	b.Log("Get Bingo Card", card)

	return server.NewResponse(card)
}

func (b *Bingo) GetCardsQRHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	qr := server.NewQRCode(fmt.Sprintf("https://%s/bingo/%s/new", r.Host, round.ID))

	b.Log("Get Bingo Round QR", card, "qr", qr.Content)

	return server.NewResponse(qr)
}

func (b *Bingo) GetRound(round int) (*Round, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if round < 0 || round >= len(b.Rounds) || len(b.Rounds) == 0 {
		return nil, fmt.Errorf("Round %d not found", round)
	}

	return b.Rounds[round], nil
}

func (b *Bingo) GetRoundByID(roundID string) (*Round, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	round := b.roundsByID[roundID]
	if round != nil {
		return round, nil
	}

	return nil, fmt.Errorf("round %s not found", roundID)
}

func (b *Bingo) GetRoundsHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	b.Log("Get Bingo Round", card, "cards", len(round.Cards))

	return server.NewResponse(PublicRound{
		ID:    round.ID,
		Round: round.Round,
		Type:  round.Type,
	})
}

func (b *Bingo) LiveHandler(w http.ResponseWriter, r *http.Request) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		response, err := server.NewResponseError(http.StatusNotFound, fmt.Errorf("round not found"))
		server.Logger.Error(err.Error())
		response.SendHasJson(w)
		return
	}
	defer roundLock.Unlock()

	card, err := round.GetCardByID(server.GetURLParam(r, "card"))

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
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(h, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	completions := NewDefaultCompletions()

	err = json.NewDecoder(h.Body).Decode(&completions)

	if err != nil {
		return server.NewResponseError(http.StatusBadRequest, errors.New("completions cannot be decoded"))
	}

	card, err := round.GetCard(0)

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}
	if card.ID != server.GetURLParam(h, "card") {
		return server.NewResponseError(http.StatusForbidden, errors.New("only the main card can configure completions"))
	}

	mutation, err := round.BeginMutation()
	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, err)
	}
	counter, err := round.SetCompletionsForAll(completions)

	if err != nil {
		_ = mutation.Restore(round)
		return server.NewResponseError(http.StatusInternalServerError, err)
	}
	if err := b.persistRound(h.Context(), round); err != nil {
		_ = mutation.Restore(round)
		return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("completions cannot be saved: %w", err))
	}
	mutation.Publish(round)

	b.Log("Set new Completions Values for the bingo Round", card, "counter", counter, "completions", card.Completions)

	return server.NewResponse(card)
}

func (b *Bingo) ToggleCardsAutoplayHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.GetCardByID(server.GetURLParam(r, "card"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	mutation, err := round.BeginMutation()
	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, err)
	}
	card.ToggleAutoplay()
	if err := b.persistRound(r.Context(), round); err != nil {
		_ = mutation.Restore(round)
		return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("autoplay cannot be saved: %w", err))
	}
	mutation.Publish(round)

	b.Log("Toggle Bingo Card Autoplay", card)

	return server.NewResponse(card)
}

func (b *Bingo) ToggleNumbersHandler(r *http.Request) (*server.Response, error) {
	round, roundLock, err := b.getRoundAndLock(server.GetURLParam(r, "round"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	defer roundLock.Unlock()

	card, err := round.GetCardByID(server.GetURLParam(r, "card"))

	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("card not found"))
	}

	mutation, err := round.BeginMutation()
	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, err)
	}
	if card.ToggleNumber(server.GetURLParamHasInt(r, "number")) && card.Card > 1 {
		b.Log("Toggle Bingo Card Number", card, "number", server.GetURLParamHasInt(r, "number"))
	}

	if card.Card == 1 {
		checked, unchecked := round.ToggleNumberForAll(server.GetURLParamHasInt(r, "number"))

		b.Log("Toggle Bingo Card Number for All", card, "number", server.GetURLParamHasInt(r, "number"), "checked", checked, "unchecked", unchecked)
	}
	if err := b.persistRound(r.Context(), round); err != nil {
		_ = mutation.Restore(round)
		return server.NewResponseError(http.StatusInternalServerError, fmt.Errorf("card number cannot be saved: %w", err))
	}
	mutation.Publish(round)

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

func New(routes *mux.Router) *Bingo {
	b, err := NewWithStore(routes, nil)
	if err != nil {
		panic(err)
	}

	return b
}

func NewWithStore(routes *mux.Router, store Store) (b *Bingo, err error) {
	b = &Bingo{
		Rounds:     []*Round{},
		store:      store,
		roundsByID: make(map[string]*Round),
		roundLocks: make(map[string]*sync.Mutex),
	}

	if store != nil {
		b.Rounds, err = store.LoadRounds(context.Background())
		if err != nil {
			return nil, fmt.Errorf("load persisted bingo state: %w", err)
		}
		for i := range b.Rounds {
			b.Rounds[i].RestoreRuntime()
			b.roundsByID[b.Rounds[i].ID] = b.Rounds[i]
			b.roundLocks[b.Rounds[i].ID] = &sync.Mutex{}
		}
	}

	if routes != nil {
		r := routes.PathPrefix("/bingo").Subrouter()
		r.Methods(http.MethodGet).Path("/{round}/new/{type}").Handler(server.SendJson(b.AddRoundsHandler))
		r.Methods(http.MethodGet).Path("/{round}").Handler(server.SendJson(b.GetRoundsHandler))
		r.Methods(http.MethodGet).Path("/{round}/0").Handler(server.SendJson(b.AddCardsHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}").Handler(server.SendJson(b.GetCardsHandler))
		r.Methods(http.MethodPost).Path("/{round}/{card}/completions").Handler(server.SendJson(b.SetCompletionsHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/0").Handler(server.SendJson(b.DrawHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/autoplay").Handler(server.SendJson(b.ToggleCardsAutoplayHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/cancel").Handler(server.SendJson(b.CancelAlertHandler))
		r.Methods(http.MethodGet).Path("/{round}/{card}/{number}").Handler(server.SendJson(b.ToggleNumbersHandler))
	}

	return b, nil
}
