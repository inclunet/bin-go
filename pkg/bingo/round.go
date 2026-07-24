package bingo

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Round struct {
	ID          string
	Cards       []Card
	Round       int
	Type        int
	NextRoundID string
	CreationID  string `json:"-"`
	upgrader    websocket.Upgrader
}

type RoundMutation struct {
	snapshot    []byte
	connections []*websocket.Conn
	runtimes    []*cardRuntime
	playerIDs   []string
	creationID  string
	upgrader    websocket.Upgrader
}

func (r *Round) AddCard() (*Card, error) {
	card := NewCard(r)

	r.Cards = append(r.Cards, card)
	r.RelinkCards()

	return r.GetCard(card.Card - 1)
}

func (r Round) Draw() *Card {
	mainCard, err := r.GetCard(0)

	if err != nil {
		return nil
	}

	mainCard.Draw()

	return mainCard
}

func (r *Round) GetCard(card int) (*Card, error) {
	if card < 0 || card >= len(r.Cards) || len(r.Cards) == 0 {
		return nil, fmt.Errorf("Card %d not found", card)
	}

	return &r.Cards[card], nil
}

func (r *Round) GetCardByID(cardID string) (*Card, error) {
	for i := range r.Cards {
		if r.Cards[i].ID == cardID {
			return &r.Cards[i], nil
		}
	}

	return nil, fmt.Errorf("card %s not found", cardID)
}

func (r *Round) SetCompletionsForAll(completions *Completions) (int, error) {
	counter := 0

	if len(r.Cards) == 0 {
		return counter, fmt.Errorf("no cards found")
	}

	for i := range r.Cards {
		if err := r.Cards[i].SetCompletions(completions); err != nil {
			return counter, err
		}

		counter++
	}

	return counter, nil
}

func (r *Round) SetNextRoundForAll(nextRound *Round) int {
	count := 0

	if nextRound == nil {
		return count
	}

	for card := range r.Cards {
		if r.Cards[card].SetNextRound(nextRound.Round, nextRound.ID) {
			count++
		}
	}

	r.NextRoundID = nextRound.ID

	return count
}

func (r *Round) SetRoundForAll(round *Round) bool {
	for card := range r.Cards {
		r.Cards[card].SetNextRound(round.Round, round.ID)
	}

	return false
}

func (r *Round) ToggleAutoplay(card int) *Card {
	currentCard, err := r.GetCard(card)

	if err != nil {
		return nil
	}

	// currentCard.ToggleAutoplay().CheckDrawedNumbers()

	return currentCard
}

func (r *Round) ToggleNumberForAll(number int) (int, int) {
	checkCounter := 0
	uncheckCounter := 0

	for i, card := range r.Cards {
		if i > 0 {
			if card.Autoplay && !card.IsChecked(number) {
				if r.Cards[i].CheckNumber(number) {
					checkCounter++
				}
			} else {
				if r.Cards[i].UncheckNumber(number) {
					uncheckCounter++
				}
			}
		}
	}

	return checkCounter, uncheckCounter
}

func (r *Round) UncheckNumberForAll(number int) *Round {
	for i := range r.Cards {
		if i > 0 {
			r.Cards[i].UncheckNumber(number)
		}
	}

	return r
}

func (r *Round) RestoreRuntime() {
	r.upgrader = newWebsocketUpgrader()

	for i := range r.Cards {
		r.Cards[i].RoundID = r.ID
		r.Cards[i].Round = r.Round
		r.Cards[i].Type = r.Type
		r.Cards[i].runtime = &cardRuntime{}
	}

	r.RelinkCards()
}

func (r *Round) RelinkCards() {
	if len(r.Cards) == 0 {
		return
	}

	r.Cards[0].Main = nil
	main := &r.Cards[0]
	for i := 1; i < len(r.Cards); i++ {
		r.Cards[i].Main = main
	}
}

func (r *Round) BeginMutation() (*RoundMutation, error) {
	snapshot, err := json.Marshal(r)
	if err != nil {
		return nil, fmt.Errorf("snapshot bingo round: %w", err)
	}

	connections := make([]*websocket.Conn, len(r.Cards))
	runtimes := make([]*cardRuntime, len(r.Cards))
	playerIDs := make([]string, len(r.Cards))
	for i := range r.Cards {
		playerIDs[i] = r.Cards[i].PlayerID
		runtime := r.Cards[i].getRuntime()
		runtime.writeMu.Lock()
		runtime.updateSeq.Add(1)
		runtime.queueMu.Lock()
		runtime.pendingSend = nil
		runtime.queueMu.Unlock()
		runtime.connMu.Lock()
		connections[i] = runtime.conn
		runtimes[i] = runtime
		runtime.conn = nil
		runtime.connMu.Unlock()
		runtime.writeMu.Unlock()
	}

	return &RoundMutation{
		snapshot:    snapshot,
		connections: connections,
		runtimes:    runtimes,
		playerIDs:   playerIDs,
		creationID:  r.CreationID,
		upgrader:    r.upgrader,
	}, nil
}

func (m *RoundMutation) Restore(round *Round) error {
	if err := json.Unmarshal(m.snapshot, round); err != nil {
		return fmt.Errorf("restore bingo round: %w", err)
	}

	round.CreationID = m.creationID
	round.upgrader = m.upgrader
	for i := range round.Cards {
		if i < len(m.playerIDs) {
			round.Cards[i].PlayerID = m.playerIDs[i]
		}
	}
	round.RelinkCards()
	m.attachConnections(round)
	return nil
}

func (m *RoundMutation) Publish(round *Round) {
	m.attachConnections(round)
	round.Publish()
}

func (r *Round) Publish() {
	for i := range r.Cards {
		send, err := r.Cards[i].PrepareUpdate()
		if err == nil && send != nil {
			r.Cards[i].QueueUpdate(send)
		}
	}
}

func (m *RoundMutation) attachConnections(round *Round) {
	for i := range round.Cards {
		if i < len(m.runtimes) && m.runtimes[i] != nil {
			round.Cards[i].runtime = m.runtimes[i]
		} else {
			round.Cards[i].runtime = &cardRuntime{}
		}
		runtime := round.Cards[i].runtime
		runtime.connMu.Lock()
		if i < len(m.connections) {
			runtime.conn = m.connections[i]
		}
		runtime.connMu.Unlock()
	}
}

func nextRoundNumber(bingo *Bingo) int {
	next := 1
	for _, round := range bingo.Rounds {
		if round.Round >= next {
			next = round.Round + 1
		}
	}
	return next
}

func newWebsocketUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkWebSocketOrigin,
	}
}

func checkWebSocketOrigin(r *http.Request) bool {
	origin := r.Header.Get("Origin")
	if origin == "" {
		return true
	}

	originURL, err := url.Parse(origin)
	if err != nil || originURL.Host == "" {
		return false
	}
	if !strings.EqualFold(originURL.Scheme, "http") && !strings.EqualFold(originURL.Scheme, "https") {
		return false
	}
	if sameWebSocketOrigin(r, originURL) {
		return true
	}

	requestURL, err := url.Parse("//" + r.Host)
	if err != nil {
		return false
	}
	return isLoopbackHost(originURL.Hostname()) && isLoopbackHost(requestURL.Hostname())
}

func sameWebSocketOrigin(r *http.Request, originURL *url.URL) bool {
	if !strings.EqualFold(originURL.Scheme, "http") && !strings.EqualFold(originURL.Scheme, "https") {
		return false
	}

	requestURL, err := url.Parse("//" + r.Host)
	if err != nil || !strings.EqualFold(originURL.Hostname(), requestURL.Hostname()) {
		return false
	}

	if requestURL.Port() == "" && originURL.Port() == "" && r.TLS == nil &&
		strings.TrimSpace(r.Header.Get("X-Forwarded-Proto")) == "" {
		// A TLS-terminating proxy may omit the external scheme. With matching
		// hostnames and implicit ports, the browser Origin remains authoritative.
		return true
	}

	return normalizedPort(originURL.Scheme, originURL.Port()) ==
		normalizedPort(requestScheme(r), requestURL.Port())
}

func normalizedPort(scheme, port string) string {
	if port != "" {
		return port
	}
	if strings.EqualFold(scheme, "https") || strings.EqualFold(scheme, "wss") {
		return "443"
	}
	if strings.EqualFold(scheme, "http") || strings.EqualFold(scheme, "ws") {
		return "80"
	}
	return ""
}

func isLoopbackHost(host string) bool {
	if strings.EqualFold(host, "localhost") ||
		strings.EqualFold(host, "localhost.localdomain") {
		return true
	}
	ip := net.ParseIP(host)
	return ip != nil && ip.IsLoopback()
}

func NewRound(bingo *Bingo, roundType int) Round {
	round := Round{
		ID:       uuid.NewString(),
		Round:    nextRoundNumber(bingo),
		Type:     roundType,
		upgrader: newWebsocketUpgrader(),
		Cards:    []Card{},
	}

	round.AddCard()

	return round
}
