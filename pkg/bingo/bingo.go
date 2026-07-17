package bingo

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/inclunet/bin-go/pkg/server"
)

type Bingo struct {
	Rounds     []*Round
	store      Store
	roundsByID map[string]*Round
	roundLocks map[string]*sync.Mutex
	creationMu sync.Mutex
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

func persistenceResponseError(message string, err error) (*server.Response, error) {
	server.Logger.Error(message, "error", err)
	return server.NewResponseError(http.StatusInternalServerError, errors.New(message))
}

func newCardResponse(card *Card) (*server.Response, error) {
	payload, err := json.Marshal(card)
	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("card response cannot be created"))
	}

	var snapshot Card
	if err := json.Unmarshal(payload, &snapshot); err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("card response cannot be created"))
	}
	return server.NewResponse(&snapshot)
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

	playerID := strings.TrimSpace(r.Header.Get("X-Bingo-Player-ID"))
	if _, err := uuid.Parse(playerID); err != nil {
		return server.NewResponseError(http.StatusBadRequest, errors.New("valid player id is required"))
	}
	for i := range round.Cards {
		if round.Cards[i].PlayerID == playerID {
			return newCardResponse(&round.Cards[i])
		}
	}

	card, err := round.AddCard()

	if err != nil {
		return server.NewResponseError(http.StatusInternalServerError, errors.New("new card cannot be added"))
	}
	card.PlayerID = playerID
	if err := b.persistRound(r.Context(), round); err != nil {
		round.Cards = round.Cards[:len(round.Cards)-1]
		round.RelinkCards()
		return persistenceResponseError("new card cannot be saved", err)
	}

	b.Log("Add Bingo Card", card)

	return newCardResponse(card)
}

func (b *Bingo) AddRoundsHandler(r *http.Request) (*server.Response, error) {
	b.creationMu.Lock()
	defer b.creationMu.Unlock()

	roundID := server.GetURLParam(r, "round")
	b.mu.RLock()
	old := b.roundsByID[roundID]
	oldLock := b.roundLocks[roundID]
	b.mu.RUnlock()

	players := 0
	if old != nil {
		oldLock.Lock()
		defer oldLock.Unlock()

		organizer, err := old.GetCardByID(server.GetURLParam(r, "card"))
		if err != nil || organizer.Card != 1 {
			return server.NewResponseError(http.StatusNotFound, errors.New("organizer card not found"))
		}

		if old.NextRoundID != "" {
			next, nextLock, err := b.getRoundAndLock(old.NextRoundID)
			if err != nil {
				return server.NewResponseError(http.StatusInternalServerError, errors.New("next round not found"))
			}
			defer nextLock.Unlock()

			card, err := next.GetCard(0)
			if err != nil {
				return server.NewResponseError(http.StatusInternalServerError, errors.New("main card not found"))
			}
			return newCardResponse(card)
		}
	} else {
		if roundID != "0" {
			return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
		}
		creationID := strings.TrimSpace(r.Header.Get("X-Bingo-Creation-ID"))
		if _, err := uuid.Parse(creationID); err != nil {
			return server.NewResponseError(http.StatusBadRequest, errors.New("valid creation id is required"))
		}

		var existingRoundID string
		b.mu.RLock()
		for _, existing := range b.Rounds {
			if existing.CreationID == creationID {
				existingRoundID = existing.ID
				break
			}
		}
		b.mu.RUnlock()
		if existingRoundID != "" {
			existing, existingLock, err := b.getRoundAndLock(existingRoundID)
			if err != nil {
				return server.NewResponseError(http.StatusInternalServerError, errors.New("existing round not found"))
			}
			defer existingLock.Unlock()
			card, err := existing.GetCard(0)
			if err != nil {
				return server.NewResponseError(http.StatusInternalServerError, errors.New("main card not found"))
			}
			return newCardResponse(card)
		}
	}

	b.mu.Lock()
	newRound := NewRound(b, server.GetURLParamHasInt(r, "type"))
	b.mu.Unlock()

	round := &newRound
	if old == nil {
		round.CreationID = strings.TrimSpace(r.Header.Get("X-Bingo-Creation-ID"))
	}
	card, err := round.GetCard(0)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("main card not found"))
	}

	if old != nil {
		previousRoundID := old.NextRoundID
		previousNextRounds := make([]int, len(old.Cards))
		previousNextRoundIDs := make([]string, len(old.Cards))
		for i := range old.Cards {
			previousNextRounds[i] = old.Cards[i].NextRound
			previousNextRoundIDs[i] = old.Cards[i].NextRoundID
		}

		players = old.SetNextRoundForAll(round)
		if err := b.persistRounds(r.Context(), round, old); err != nil {
			old.NextRoundID = previousRoundID
			for i := range old.Cards {
				old.Cards[i].NextRound = previousNextRounds[i]
				old.Cards[i].NextRoundID = previousNextRoundIDs[i]
			}
			return persistenceResponseError("rounds cannot be saved", err)
		}
	} else if err := b.persistRound(r.Context(), round); err != nil {
		return persistenceResponseError("round cannot be saved", err)
	}

	b.mu.Lock()
	b.Rounds = append(b.Rounds, round)
	b.roundsByID[round.ID] = round
	b.roundLocks[round.ID] = &sync.Mutex{}
	b.mu.Unlock()

	if old != nil {
		old.Publish()
		b.Log("Redirect Old Players to the New Bingo Round", card, "from", old.Round, "to", round.Round, "players", players)
	}

	b.Log("Add Bingo Round", card)

	return newCardResponse(card)
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
		return persistenceResponseError("card cannot be saved", err)
	}
	mutation.Publish(round)

	b.Log("Cancel Bingo Alert", card)

	return newCardResponse(card)
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
	if number == 0 {
		mutation.Publish(round)
		return newCardResponse(card)
	}

	checked, Unchecked := round.ToggleNumberForAll(number)
	if err := b.persistRound(r.Context(), round); err != nil {
		_ = mutation.Restore(round)
		return persistenceResponseError("draw cannot be saved", err)
	}
	mutation.Publish(round)

	b.Log("Draw new Random Bingo Number", card, "number", number, "checked", checked, "unchecked", Unchecked)

	return newCardResponse(card)
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

	return newCardResponse(card)
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

	qr := server.NewQRCode(fmt.Sprintf("%s://%s/bingo/%s/new", requestScheme(r), r.Host, round.ID))

	b.Log("Get Bingo Round QR", card, "qr", qr.Content)

	return server.NewResponse(qr)
}

func requestScheme(r *http.Request) string {
	if forwarded := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-Proto"), ",")[0]); forwarded != "" {
		if strings.EqualFold(forwarded, "https") {
			return "https"
		}
		if strings.EqualFold(forwarded, "http") {
			return "http"
		}
	}
	if r.TLS != nil {
		return "https"
	}
	return "http"
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

	card, err := round.GetCardByID(server.GetURLParam(r, "card"))

	if err != nil {
		roundLock.Unlock()
		response, err := server.NewResponseError(http.StatusNotFound, fmt.Errorf("card not found"))
		server.Logger.Error(err.Error())
		response.SendHasJson(w)
		return
	}

	upgrader := round.upgrader
	roundLock.Unlock()

	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		server.Logger.Error("websocket upgrade failed", "error", err)
		return
	}

	round, roundLock, err = b.getRoundAndLock(server.GetURLParam(r, "round"))
	if err != nil {
		_ = conn.Close()
		return
	}
	card, err = round.GetCardByID(server.GetURLParam(r, "card"))
	if err != nil {
		roundLock.Unlock()
		_ = conn.Close()
		return
	}
	if !card.SetConn(conn) {
		roundLock.Unlock()
		_ = conn.Close()
		return
	}
	sendUpdate, err := card.PrepareUpdate()
	if err != nil || sendUpdate == nil {
		roundLock.Unlock()
		_ = conn.Close()
		return
	}
	roundLock.Unlock()
	if err := sendUpdate(); err != nil {
		_ = conn.Close()
	}
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
		return persistenceResponseError("completions cannot be saved", err)
	}
	mutation.Publish(round)

	b.Log("Set new Completions Values for the bingo Round", card, "counter", counter, "completions", card.Completions)

	return newCardResponse(card)
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
		return persistenceResponseError("autoplay cannot be saved", err)
	}
	mutation.Publish(round)

	b.Log("Toggle Bingo Card Autoplay", card)

	return newCardResponse(card)
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
		return persistenceResponseError("card number cannot be saved", err)
	}
	mutation.Publish(round)

	return newCardResponse(card)
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
		r.Methods(http.MethodGet).Path("/{round}/{card}/new/{type}").Handler(server.SendJson(b.AddRoundsHandler))
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
