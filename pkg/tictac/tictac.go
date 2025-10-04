package tictac

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/inclunet/bin-go/pkg/server"
)

type Game struct {
	Rounds        []*Round          `json:"rounds"`
	openWatchers  []*websocket.Conn `json:"-"`
	openWatchLock sync.Mutex        `json:"-"`
}

type Round struct {
	Round     int          `json:"round"`
	Board     [3][3]string `json:"board"`
	PlayerX   string       `json:"playerX"`
	PlayerO   string       `json:"playerO"`
	Turn      string       `json:"turn"` // "X" ou "O"
	Winner    string       `json:"winner"`
	Next      int          `json:"next,omitempty"`
	ScoreX    int          `json:"scoreX,omitempty"`
	ScoreO    int          `json:"scoreO,omitempty"`
	ScoreDraw int          `json:"scoreDraw,omitempty"`
	counted   bool         `json:"-"`
	players   []*Player    `json:"-"`
	lock      sync.Mutex   `json:"-"`
}

type Player struct {
	Conn      *websocket.Conn
	Piece     string // "X" | "O" | "" (espectador)
	Spectator bool
}

func NewGame() *Game { return &Game{Rounds: []*Round{}} }

func (g *Game) newRoundHandler(r *http.Request) (*server.Response, error) {
	current := server.GetURLParamHasInt(r, "round")
	// Caso base: nenhuma rodada ainda
	if len(g.Rounds) == 0 {
		if current != 1 {
			return server.NewResponseError(http.StatusBadRequest, errors.New("invalid round number"))
		}
		first := &Round{Round: 1, Turn: ""}
		g.Rounds = append(g.Rounds, first)
		server.Logger.Info("Add TicTac Round", "round", 1, "first", true)
		g.broadcastOpenRounds()
		return server.NewResponse(first)
	}

	nextNumber := len(g.Rounds) + 1
	if current == nextNumber {
		// Nova cadeia independente (como bingo): scoreboard zerado, turno adiado
		newR := &Round{Round: nextNumber, Turn: "", ScoreX: 0, ScoreO: 0, ScoreDraw: 0}
		g.Rounds = append(g.Rounds, newR)
		server.Logger.Info("Add TicTac Round (new sequence)", "round", newR.Round)
		g.broadcastOpenRounds()
		return server.NewResponse(newR)
	}

	// Continuação a partir de uma rodada específica já concluída (source = current)
	sourceIdx := current - 1
	if sourceIdx < 0 || sourceIdx >= len(g.Rounds) {
		return server.NewResponseError(http.StatusNotFound, errors.New("source round not found"))
	}
	source := g.Rounds[sourceIdx]
	if source.Winner == "" {
		return server.NewResponseError(http.StatusPreconditionFailed, errors.New("source round not finished"))
	}
	if source.Next != 0 {
		return server.NewResponseError(http.StatusConflict, errors.New("round already continued"))
	}
	startTurn := startingTurnFrom(source)
	cont := &Round{Round: nextNumber, Turn: startTurn, ScoreX: source.ScoreX, ScoreO: source.ScoreO, ScoreDraw: source.ScoreDraw}
	g.Rounds = append(g.Rounds, cont)
	source.Next = cont.Round
	source.broadcastLocked()
	server.Logger.Info("Add TicTac Round (continue)", "round", cont.Round, "from", source.Round)
	g.broadcastOpenRounds()
	return server.NewResponse(cont)
}

func (g *Game) getRound(idx int) (*Round, error) {
	if idx < 0 || idx >= len(g.Rounds) || len(g.Rounds) == 0 {
		return nil, fmt.Errorf("round %d not found", idx+1)
	}
	return g.Rounds[idx], nil
}

func (g *Game) getRoundHandler(r *http.Request) (*server.Response, error) {
	rd, err := g.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	return server.NewResponse(rd)
}

// Gera QR com link absoluto (ou relativo) para a página da rodada de seleção de símbolo.
func (g *Game) qrRoundHandler(r *http.Request) (*server.Response, error) {
	roundNum := server.GetURLParamHasInt(r, "round")
	if roundNum <= 0 || roundNum > len(g.Rounds) {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	// Tentar montar origem (host) se cabecalhos presentes; fallback relativo.
	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	url := fmt.Sprintf("%s://%s/tictac/%d", scheme, host, roundNum)
	qrb := server.NewQRCode(url)
	qrb.SetSize(340, 340)
	return &server.Response{StatusCode: http.StatusOK, Body: qrb}, nil
}

// openRoundsHandler retorna rodadas que ainda aguardam jogadores (ao menos um símbolo não escolhido e não finalizada)
func (g *Game) openRoundsHandler(r *http.Request) (*server.Response, error) {
	type openRound struct {
		Round   int  `json:"round"`
		PlayerX bool `json:"hasPlayerX"`
		PlayerO bool `json:"hasPlayerO"`
	}
	res := []openRound{}
	for _, rd := range g.Rounds {
		if rd == nil {
			continue
		}
		if rd.Winner != "" {
			continue
		}
		// aguardando jogadores se ao menos um símbolo não está definido
		if rd.PlayerX == "" || rd.PlayerO == "" {
			res = append(res, openRound{Round: rd.Round, PlayerX: rd.PlayerX != "", PlayerO: rd.PlayerO != ""})
		}
	}
	return server.NewResponse(res)
}

func (g *Game) moveHandler(r *http.Request) (*server.Response, error) {
	idx := server.GetURLParamHasInt(r, "round") - 1
	rd, err := g.getRound(idx)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	rd.lock.Lock()
	defer rd.lock.Unlock()
	if rd.Winner != "" {
		return server.NewResponse(rd)
	}
	row := server.GetURLParamHasInt(r, "row")
	col := server.GetURLParamHasInt(r, "col")
	player := server.GetURLParamWithDefault(r, "player", "x")
	piece := "X"
	if player == "o" {
		piece = "O"
	}
	// Se turno ainda não definido (primeiro round antes de qualquer conexão atribuir) assume peça do primeiro movimento
	if rd.Turn == "" {
		rd.Turn = piece
	}
	// garantir que o jogador que tenta jogar é realmente o dono da peça registrada
	if piece == "X" && rd.PlayerX == "" {
		rd.PlayerX = "X"
	} // fallback caso primeiro movimento venha via REST antes do WS
	if piece == "O" && rd.PlayerO == "" {
		rd.PlayerO = "O"
	}
	// Bloqueia jogadas até ambos jogadores terem entrado
	if rd.PlayerX == "" || rd.PlayerO == "" {
		return server.NewResponseError(http.StatusPreconditionFailed, errors.New("waiting opponent"))
	}
	if (piece == "X" && rd.PlayerX != "X") || (piece == "O" && rd.PlayerO != "O") {
		return server.NewResponseError(http.StatusForbidden, errors.New("player not registered for this piece"))
	}
	if rd.Turn != piece {
		return server.NewResponse(rd)
	}
	if row < 1 || row > 3 || col < 1 || col > 3 {
		return server.NewResponseError(http.StatusBadRequest, errors.New("invalid position"))
	}
	if rd.Board[row-1][col-1] != "" {
		return server.NewResponse(rd)
	}
	rd.Board[row-1][col-1] = piece
	checkWinner(rd)
	switch rd.Winner {
	case "":
		// ninguém venceu ainda, alterna turno
		toggleTurn(rd)
	case "X", "O", "Empate":
		if !rd.counted { // incrementa somente uma vez
			rd.counted = true
			switch rd.Winner {
			case "X":
				rd.ScoreX += 1
			case "O":
				rd.ScoreO += 1
			case "Empate":
				rd.ScoreDraw += 1
			}
		}
		server.Logger.Info("TicTac Winner", "round", rd.Round, "winner", rd.Winner)
	}
	// Log de jogada detalhado removido para reduzir ruído; vencedor continua logado acima.
	rd.broadcastLocked()
	// Caso a rodada tenha deixado de estar aberta (ambos jogadores) ou tenha terminado, atualizar lista
	g.broadcastOpenRounds()
	return server.NewResponse(rd)
}

func toggleTurn(r *Round) {
	if r.Turn == "X" {
		r.Turn = "O"
	} else {
		r.Turn = "X"
	}
}

func checkWinner(r *Round) {
	b := r.Board
	lines := [][]string{
		{b[0][0], b[0][1], b[0][2]}, {b[1][0], b[1][1], b[1][2]}, {b[2][0], b[2][1], b[2][2]},
		{b[0][0], b[1][0], b[2][0]}, {b[0][1], b[1][1], b[2][1]}, {b[0][2], b[1][2], b[2][2]},
		{b[0][0], b[1][1], b[2][2]}, {b[0][2], b[1][1], b[2][0]},
	}
	for _, ln := range lines {
		if ln[0] != "" && ln[0] == ln[1] && ln[1] == ln[2] {
			r.Winner = ln[0]
			return
		}
	}
	full := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == "" {
				full = false
				break
			}
		}
	}
	if full {
		r.Winner = "Empate"
	}
}

// startingTurnFrom decide o turno inicial da próxima rodada baseado no resultado da anterior
// Regras:
// - Vitória de X => X inicia
// - Vitória de O => O inicia
// - Empate => inicia quem jogaria depois (ou seja, o oposto de quem fez a última jogada)
func startingTurnFrom(last *Round) string {
	if last == nil {
		return "X"
	}
	if last.Winner == "X" {
		return "X"
	}
	if last.Winner == "O" {
		return "O"
	}
	if last.Winner == "Empate" {
		// Quem jogaria depois? Se último turno registrado foi de X, então próximo seria O, e vice-versa.
		if last.Turn == "X" {
			return "O"
		}
		return "X"
	}
	// fallback
	return "X"
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (r *Round) addConnection(conn *websocket.Conn, requested string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	piece := ""
	spectator := false
	switch requested {
	case "x", "X":
		if r.PlayerX == "" {
			r.PlayerX = "X"
			piece = "X"
		} else {
			spectator = true
		}
	case "o", "O":
		if r.PlayerO == "" {
			r.PlayerO = "O"
			piece = "O"
		} else {
			spectator = true
		}
	default:
		spectator = true
	}
	if piece == "" && !spectator {
		spectator = true
	}
	r.players = append(r.players, &Player{Conn: conn, Piece: piece, Spectator: spectator})
	if r.Turn == "" && piece != "" {
		r.Turn = piece
	}
	server.Logger.Info("TicTac Connect", "round", r.Round, "piece", piece, "spectator", spectator, "players", len(r.players))
}

func (r *Round) compactPlayersLocked() {
	cleaned := make([]*Player, 0, len(r.players))
	for _, p := range r.players {
		if p != nil && p.Conn != nil {
			cleaned = append(cleaned, p)
		}
	}
	r.players = cleaned
}

func (r *Round) broadcastLocked() {
	state, _ := json.Marshal(r)
	for _, p := range r.players {
		if p == nil || p.Conn == nil {
			continue
		}
		if err := p.Conn.WriteMessage(websocket.TextMessage, state); err != nil {
			p.Conn.Close()
			p.Conn = nil
		}
	}
	r.compactPlayersLocked()
}

func (g *Game) liveHandler(w http.ResponseWriter, r *http.Request) {
	rd, err := g.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		http.Error(w, "round not found", http.StatusNotFound)
		return
	}
	piece := server.GetURLParamWithDefault(r, "player", "")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "upgrade error", http.StatusInternalServerError)
		return
	}
	rd.addConnection(conn, piece)
	// envia estado atual apenas ao novo cliente primeiro
	if state, mErr := json.Marshal(rd); mErr == nil {
		conn.WriteMessage(websocket.TextMessage, state)
	}
	// depois broadcast geral para sincronizar espectadores
	rd.lock.Lock()
	rd.broadcastLocked()
	rd.lock.Unlock()
	// Pode mudar status de aberta (segundo jogador entrou)
	g.broadcastOpenRounds()
	go func(c *websocket.Conn, round *Round) {
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				c.Close()
				return
			}
		}
	}(conn, rd)
}

// --- Open rounds (aguardando jogadores) ---
func (g *Game) collectOpenRounds() []map[string]interface{} {
	res := []map[string]interface{}{}
	for _, rd := range g.Rounds {
		if rd == nil {
			continue
		}
		if rd.Winner != "" {
			continue
		}
		if rd.PlayerX == "" || rd.PlayerO == "" { // ainda aberta
			res = append(res, map[string]interface{}{
				"round":      rd.Round,
				"hasPlayerX": rd.PlayerX != "",
				"hasPlayerO": rd.PlayerO != "",
			})
		}
	}
	return res
}

func (g *Game) broadcastOpenRounds() {
	g.openWatchLock.Lock()
	defer g.openWatchLock.Unlock()
	if len(g.openWatchers) == 0 {
		return
	}
	payload := map[string]interface{}{"rounds": g.collectOpenRounds()}
	data, _ := json.Marshal(payload)
	alive := make([]*websocket.Conn, 0, len(g.openWatchers))
	for _, c := range g.openWatchers {
		if c == nil {
			continue
		}
		if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
			c.Close()
			continue
		}
		alive = append(alive, c)
	}
	g.openWatchers = alive
}

func (g *Game) openLiveHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "upgrade error", http.StatusInternalServerError)
		return
	}
	g.openWatchLock.Lock()
	g.openWatchers = append(g.openWatchers, conn)
	g.openWatchLock.Unlock()
	// Envia lista inicial
	g.broadcastOpenRounds()
	go func(c *websocket.Conn) {
		for { // ignorar mensagens de entrada
			if _, _, err := c.ReadMessage(); err != nil {
				c.Close()
				return
			}
		}
	}(conn)
}

func New(routes *mux.Router) *Game {
	g := NewGame()
	if routes != nil {
		r := routes.PathPrefix("/tictac").Subrouter()
		r.Methods(http.MethodGet).Path("/{round}/new").Handler(server.SendJson(g.newRoundHandler))
		r.Methods(http.MethodGet).Path("/{round}").Handler(server.SendJson(g.getRoundHandler))
		r.Methods(http.MethodGet).Path("/{round}/qr").Handler(server.SendQRCode(g.qrRoundHandler))
		r.Methods(http.MethodGet).Path("/open").Handler(server.SendJson(g.openRoundsHandler))
		r.Methods(http.MethodGet).Path("/{round}/{player}/{row}/{col}").Handler(server.SendJson(g.moveHandler))
	}
	return g
}

func (g *Game) AddWsRoutes(routes *mux.Router) *Game {
	if routes != nil {
		r := routes.PathPrefix("/tictac").Subrouter()
		r.Methods(http.MethodGet).Path("/open").HandlerFunc(g.openLiveHandler)
		r.Methods(http.MethodGet).Path("/{round}/{player}").HandlerFunc(g.liveHandler)
	}
	return g
}
