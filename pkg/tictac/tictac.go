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
	Rounds []*Round `json:"rounds"`
}

type Round struct {
	Round   int      `json:"round"`
	Board   [3][3]string `json:"board"`
	PlayerX string   `json:"playerX"`
	PlayerO string   `json:"playerO"`
	Turn    string   `json:"turn"` // "X" ou "O"
	Winner  string   `json:"winner"`
	Next    int      `json:"next,omitempty"`
	players []*Player       `json:"-"`
	lock    sync.Mutex      `json:"-"`
}

type Player struct {
	Conn      *websocket.Conn
	Piece     string // "X" | "O" | "" (espectador)
	Spectator bool
}

func NewGame() *Game { return &Game{Rounds: []*Round{}} }

func (g *Game) newRoundHandler(r *http.Request) (*server.Response, error) {
	current := server.GetURLParamHasInt(r, "round")
	newR := &Round{Round: len(g.Rounds) + 1, Turn: "X"}
	g.Rounds = append(g.Rounds, newR)
	if current > 0 {
		if old, err := g.getRound(current - 1); err == nil {
			old.Next = newR.Round
			old.broadcastLocked("redirect")
		}
	}
	server.Logger.Info("Add TicTac Round", "round", newR.Round)
	return server.NewResponse(newR)
}

func (g *Game) getRound(idx int) (*Round, error) {
	if idx < 0 || idx >= len(g.Rounds) || len(g.Rounds) == 0 { return nil, fmt.Errorf("round %d not found", idx+1) }
	return g.Rounds[idx], nil
}

func (g *Game) getRoundHandler(r *http.Request) (*server.Response, error) {
	rd, err := g.getRound(server.GetURLParamHasInt(r, "round")-1)
	if err != nil { return server.NewResponseError(http.StatusNotFound, errors.New("round not found")) }
	return server.NewResponse(rd)
}

func (g *Game) moveHandler(r *http.Request) (*server.Response, error) {
	idx := server.GetURLParamHasInt(r, "round") - 1
	rd, err := g.getRound(idx)
	if err != nil { return server.NewResponseError(http.StatusNotFound, errors.New("round not found")) }
	rd.lock.Lock()
	defer rd.lock.Unlock()
	if rd.Winner != "" { return server.NewResponse(rd) }
	row := server.GetURLParamHasInt(r, "row")
	col := server.GetURLParamHasInt(r, "col")
	player := server.GetURLParamWithDefault(r, "player", "x")
	piece := "X"; if player == "o" { piece = "O" }
	// garantir que o jogador que tenta jogar é realmente o dono da peça registrada
	if piece == "X" && rd.PlayerX == "" { rd.PlayerX = "X" } // fallback caso primeiro movimento venha via REST antes do WS
	if piece == "O" && rd.PlayerO == "" { rd.PlayerO = "O" }
	if (piece == "X" && rd.PlayerX != "X") || (piece == "O" && rd.PlayerO != "O") {
		return server.NewResponseError(http.StatusForbidden, errors.New("player not registered for this piece"))
	}
	if rd.Turn != piece { return server.NewResponse(rd) }
	if row < 1 || row > 3 || col < 1 || col > 3 { return server.NewResponseError(http.StatusBadRequest, errors.New("invalid position")) }
	if rd.Board[row-1][col-1] != "" { return server.NewResponse(rd) }
	rd.Board[row-1][col-1] = piece
	checkWinner(rd)
	if rd.Winner == "" { toggleTurn(rd) } else { server.Logger.Info("TicTac Winner", "round", rd.Round, "winner", rd.Winner) }
	server.Logger.Info("TicTac Move", "round", rd.Round, "player", piece, "row", row, "col", col, "turn", rd.Turn, "winner", rd.Winner)
	rd.broadcastLocked("state")
	return server.NewResponse(rd)
}

func toggleTurn(r *Round) { if r.Turn == "X" { r.Turn = "O" } else { r.Turn = "X" } }

func checkWinner(r *Round) {
	b := r.Board
	lines := [][]string{
		{b[0][0], b[0][1], b[0][2]}, {b[1][0], b[1][1], b[1][2]}, {b[2][0], b[2][1], b[2][2]},
		{b[0][0], b[1][0], b[2][0]}, {b[0][1], b[1][1], b[2][1]}, {b[0][2], b[1][2], b[2][2]},
		{b[0][0], b[1][1], b[2][2]}, {b[0][2], b[1][1], b[2][0]},
	}
	for _, ln := range lines { if ln[0] != "" && ln[0] == ln[1] && ln[1] == ln[2] { r.Winner = ln[0]; return } }
	full := true
	for i:=0;i<3;i++ { for j:=0;j<3;j++ { if b[i][j]=="" { full=false; break } } }
	if full { r.Winner = "Empate" }
}

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func (r *Round) addConnection(conn *websocket.Conn, requested string) {
	r.lock.Lock()
	defer r.lock.Unlock()
	piece := ""
	spectator := false
	if requested == "x" || requested == "X" {
		if r.PlayerX == "" { r.PlayerX = "X"; piece = "X" } else { spectator = true }
	} else if requested == "o" || requested == "O" {
		if r.PlayerO == "" { r.PlayerO = "O"; piece = "O" } else { spectator = true }
	} else {
		spectator = true
	}
	if piece == "" && !spectator { spectator = true }
	r.players = append(r.players, &Player{Conn: conn, Piece: piece, Spectator: spectator})
	server.Logger.Info("TicTac Connect", "round", r.Round, "piece", piece, "spectator", spectator, "players", len(r.players))
}

func (r *Round) compactPlayersLocked() {
	cleaned := make([]*Player, 0, len(r.players))
	for _, p := range r.players { if p != nil && p.Conn != nil { cleaned = append(cleaned, p) } }
	r.players = cleaned
}

func (r *Round) broadcastLocked(event string) {
	state, _ := json.Marshal(r)
	for _, p := range r.players {
		if p == nil || p.Conn == nil { continue }
		if err := p.Conn.WriteMessage(websocket.TextMessage, state); err != nil {
			p.Conn.Close(); p.Conn = nil
		}
	}
	r.compactPlayersLocked()
}

func (g *Game) liveHandler(w http.ResponseWriter, r *http.Request) {
	rd, err := g.getRound(server.GetURLParamHasInt(r, "round")-1)
	if err != nil { http.Error(w, "round not found", http.StatusNotFound); return }
	piece := server.GetURLParamWithDefault(r, "player", "")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil { http.Error(w, "upgrade error", http.StatusInternalServerError); return }
	rd.addConnection(conn, piece)
	// envia estado atual apenas ao novo cliente primeiro
	if state, mErr := json.Marshal(rd); mErr == nil { conn.WriteMessage(websocket.TextMessage, state) }
	// depois broadcast geral para sincronizar espectadores
	rd.lock.Lock(); rd.broadcastLocked("join"); rd.lock.Unlock()
	go func(c *websocket.Conn, round *Round) {
		for {
			if _, _, err := c.ReadMessage(); err != nil { c.Close(); return }
		}
	}(conn, rd)
}

func New(routes *mux.Router) *Game {
	g := NewGame()
	if routes != nil {
		r := routes.PathPrefix("/tictac").Subrouter()
		r.Methods(http.MethodGet).Path("/{round}/new").Handler(server.SendJson(g.newRoundHandler))
		r.Methods(http.MethodGet).Path("/{round}").Handler(server.SendJson(g.getRoundHandler))
		r.Methods(http.MethodGet).Path("/{round}/{player}/{row}/{col}").Handler(server.SendJson(g.moveHandler))
	}
	return g
}

func (g *Game) AddWsRoutes(routes *mux.Router) *Game {
	if routes != nil {
		r := routes.PathPrefix("/tictac").Subrouter()
		r.Methods(http.MethodGet).Path("/{round}/{player}").HandlerFunc(g.liveHandler)
	}
	return g
}
