package battleship

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"github.com/inclunet/bin-go/pkg/server"
)

// Battleship representa o controlador principal do jogo
type Battleship struct {
	Rounds        []*Round          `json:"rounds"`
	openWatchers  []*websocket.Conn `json:"-"`
	openWatchLock sync.Mutex        `json:"-"`
}

// New cria uma nova instância do jogo Battleship
func New(routes *mux.Router) *Battleship {
	b := &Battleship{
		Rounds: []*Round{},
	}

	if routes != nil {
		r := routes.PathPrefix("/battleship").Subrouter()

		// Rotas REST (ações place/shoot removidas - agora via WS)
		r.Methods(http.MethodGet).Path("/{round}/new").Handler(server.SendJson(b.newRoundHandler))
		r.Methods(http.MethodGet).Path("/{round}").Handler(server.SendJson(b.getRoundHandler))
		r.Methods(http.MethodGet).Path("/{round}/qr").Handler(server.SendQRCode(b.qrRoundHandler))
		r.Methods(http.MethodGet).Path("/open").Handler(server.SendJson(b.openRoundsHandler))
		r.Methods(http.MethodGet).Path("/{round}/{player}/board").Handler(server.SendJson(b.getBoardHandler))
		r.Methods(http.MethodGet).Path("/{round}/{player}/ships").Handler(server.SendJson(b.getShipsHandler))
	}

	return b
}

// AddWsRoutes adiciona as rotas WebSocket
func (b *Battleship) AddWsRoutes(routes *mux.Router) *Battleship {
	// 'routes' já deve ser o subrouter com prefixo /ws definido em main.go
	// Evitar duplicar /ws no caminho final. Caminho alvo: /ws/battleship/...
	if routes != nil {
		r := routes.PathPrefix("/battleship").Subrouter()
		r.Methods(http.MethodGet).Path("/open").HandlerFunc(b.openLiveHandler)
		r.Methods(http.MethodGet).Path("/{round}/{player}").HandlerFunc(b.liveHandler)
	}
	return b
}

// getRound obtém uma partida pelo índice
func (b *Battleship) getRound(idx int) (*Round, error) {
	if idx < 0 || idx >= len(b.Rounds) || len(b.Rounds) == 0 {
		return nil, fmt.Errorf("round %d not found", idx+1)
	}
	return b.Rounds[idx], nil
}

// newRoundHandler cria uma nova partida
func (b *Battleship) newRoundHandler(r *http.Request) (*server.Response, error) {
	// Padrão de criação/continuação de rodadas:
	// Mantém a mesma semântica já usada em TicTac (ver pkg/tictac/tictac.go:newRoundHandler) e análoga à de Bingo.
	// GET /{round}/new segue regras:
	// 1) Se não existe nenhuma rodada ainda e round == 1 -> cria primeira rodada (placares zerados).
	// 2) Se round == len(Rounds)+1 -> cria nova sequência independente (placar herdado = zerado, aqui ScoreA/B/Draw já começam 0).
	// 3) Se round aponta para uma rodada existente já concluída (Winner definido) -> cria continuação: nova rodada com número len(Rounds)+1
	//    e herda placar acumulado (ScoreA, ScoreB, ScoreDraw). Amarra campo Next na rodada anterior.
	// 4) Se round aponta para uma rodada não finalizada -> erro (PreconditionFailed).
	// 5) Se a rodada fonte já tem Next definido -> erro (Conflict) evitando múltiplas continuações.
	// Esta lógica evita endpoints extras (como POST /rounds) e mantém padronização simples via GET semelhante aos outros jogos.
	// Qualquer alteração estrutural futura deve também atualizar TicTac para preservar consistência cross-jogos.
	current := server.GetURLParamHasInt(r, "round")

	// Caso base: nenhuma rodada ainda
	if len(b.Rounds) == 0 {
		if current != 1 {
			return server.NewResponseError(http.StatusBadRequest, errors.New("invalid round number"))
		}
		first := NewRound(1)
		b.Rounds = append(b.Rounds, first)
		server.Logger.Info("Add Battleship Round", "round", 1, "first", true)
		b.broadcastOpenRounds()
		return server.NewResponse(first)
	}

	nextNumber := len(b.Rounds) + 1
	if current == nextNumber {
		// Nova partida independente
		newR := NewRound(nextNumber)
		b.Rounds = append(b.Rounds, newR)
		server.Logger.Info("Add Battleship Round (new)", "round", newR.Round)
		b.broadcastOpenRounds()
		return server.NewResponse(newR)
	}

	// Continuação a partir de uma rodada específica já concluída
	sourceIdx := current - 1
	if sourceIdx < 0 || sourceIdx >= len(b.Rounds) {
		return server.NewResponseError(http.StatusNotFound, errors.New("source round not found"))
	}
	source := b.Rounds[sourceIdx]
	if source.Winner == "" {
		return server.NewResponseError(http.StatusPreconditionFailed, errors.New("source round not finished"))
	}
	if source.Next != 0 {
		return server.NewResponseError(http.StatusConflict, errors.New("round already continued"))
	}

	cont := NewRound(nextNumber)
	cont.ScoreA = source.ScoreA
	cont.ScoreB = source.ScoreB
	cont.ScoreDraw = source.ScoreDraw

	b.Rounds = append(b.Rounds, cont)
	source.Next = cont.Round
	source.BroadcastLocked()

	server.Logger.Info("Add Battleship Round (continue)", "round", cont.Round, "from", source.Round)
	b.broadcastOpenRounds()
	return server.NewResponse(cont)
}

// getRoundHandler obtém informações de uma partida
func (b *Battleship) getRoundHandler(r *http.Request) (*server.Response, error) {
	rd, err := b.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}
	return server.NewResponse(rd)
}

// qrRoundHandler gera QR Code para uma partida
func (b *Battleship) qrRoundHandler(r *http.Request) (*server.Response, error) {
	roundNum := server.GetURLParamHasInt(r, "round")
	if roundNum <= 0 || roundNum > len(b.Rounds) {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	host := r.Host
	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	url := fmt.Sprintf("%s://%s/battleship/%d", scheme, host, roundNum)
	qrb := server.NewQRCode(url)
	qrb.SetSize(340, 340)
	return &server.Response{StatusCode: http.StatusOK, Body: qrb}, nil
}

// openRoundsHandler retorna partidas em aberto
func (b *Battleship) openRoundsHandler(r *http.Request) (*server.Response, error) {
	type openRound struct {
		Round   int  `json:"round"`
		PlayerA bool `json:"hasPlayerA"`
		PlayerB bool `json:"hasPlayerB"`
	}

	res := []openRound{}
	for _, rd := range b.Rounds {
		if rd == nil || rd.Winner != "" {
			continue
		}
		// Partida em aberto se ao menos um jogador não está definido
		if rd.PlayerA == "" || rd.PlayerB == "" {
			res = append(res, openRound{
				Round:   rd.Round,
				PlayerA: rd.PlayerA != "",
				PlayerB: rd.PlayerB != "",
			})
		}
	}
	return server.NewResponse(res)
}

// placeShipHandler posiciona um navio
func (b *Battleship) placeShipHandler(r *http.Request) (*server.Response, error) {
	rd, err := b.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	playerStr := server.GetURLParamWithDefault(r, "player", "a")
	var player Player
	if playerStr == "b" {
		player = PlayerB
	} else {
		player = PlayerA
	}

	// Parse do JSON
	var req struct {
		ShipID      int    `json:"shipId"`
		Row         int    `json:"row"`
		Col         int    `json:"col"`
		Orientation string `json:"orientation"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return server.NewResponseError(http.StatusBadRequest, errors.New("invalid JSON"))
	}

	orientation := OrientationHorizontal
	if req.Orientation == "vertical" {
		orientation = OrientationVertical
	}

	if err := rd.PlacePlayerShip(player, req.ShipID, req.Row, req.Col, orientation); err != nil {
		return server.NewResponseError(http.StatusBadRequest, err)
	}

	rd.BroadcastLocked()
	server.Logger.Info("Battleship PlaceShip", "round", rd.Round, "player", player, "shipId", req.ShipID, "row", req.Row, "col", req.Col, "orientation", orientation)

	return server.NewResponse(map[string]interface{}{
		"success": true,
		"phase":   rd.Phase,
		"ready": map[string]bool{
			"a": rd.PlayerAReady,
			"b": rd.PlayerBReady,
		},
	})
}

// shootHandler executa um tiro
func (b *Battleship) shootHandler(r *http.Request) (*server.Response, error) {
	rd, err := b.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	playerStr := server.GetURLParamWithDefault(r, "player", "a")
	var player Player
	if playerStr == "b" {
		player = PlayerB
	} else {
		player = PlayerA
	}

	// Parse do JSON
	var req struct {
		Row int `json:"row"`
		Col int `json:"col"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		return server.NewResponseError(http.StatusBadRequest, errors.New("invalid JSON"))
	}

	shot, err := rd.Shoot(player, req.Row, req.Col)
	if err != nil {
		return server.NewResponseError(http.StatusBadRequest, err)
	}

	rd.BroadcastLocked()
	lastSunkID := 0
	lastSunkSize := 0
	if rd.LastSunkShip != nil {
		lastSunkID = rd.LastSunkShip.ID
		lastSunkSize = rd.LastSunkShip.Size
	}
	server.Logger.Info("Battleship Shoot",
		"round", rd.Round,
		"player", player,
		"row", req.Row,
		"col", req.Col,
		"result", shot.Result,
		"winner", rd.Winner,
		"phase", rd.Phase,
		"lastSunkShipId", lastSunkID,
		"lastSunkShipSize", lastSunkSize,
		"shipsRemainingA", CountRemainingShips(rd.shipsA),
		"shipsRemainingB", CountRemainingShips(rd.shipsB),
	)

	// Atualizar lista de partidas abertas se jogo terminou
	if rd.Winner != "" {
		b.broadcastOpenRounds()
	}

	// Construir estado completo para o jogador requisitante (mascarado conforme regras)
	state := rd.createPlayerState(player)

	return server.NewResponse(map[string]interface{}{
		"shot":         shot,
		"winner":       rd.Winner,
		"phase":        rd.Phase,
		"lastSunkShip": rd.LastSunkShip,
		"state":        state,
	})
}

// getBoardHandler retorna o tabuleiro de um jogador
func (b *Battleship) getBoardHandler(r *http.Request) (*server.Response, error) {
	rd, err := b.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	playerStr := server.GetURLParamWithDefault(r, "player", "a")
	var player Player
	if playerStr == "b" {
		player = PlayerB
	} else {
		player = PlayerA
	}

	enemy := r.URL.Query().Get("enemy") == "true"

	var board Board
	if enemy {
		board = rd.GetEnemyBoard(player)
	} else {
		board = rd.GetBoard(player)
	}

	return server.NewResponse(map[string]interface{}{
		"board": board,
		"phase": rd.Phase,
	})
}

// getShipsHandler retorna os navios de um jogador
func (b *Battleship) getShipsHandler(r *http.Request) (*server.Response, error) {
	rd, err := b.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		return server.NewResponseError(http.StatusNotFound, errors.New("round not found"))
	}

	playerStr := server.GetURLParamWithDefault(r, "player", "a")
	var player Player
	if playerStr == "b" {
		player = PlayerB
	} else {
		player = PlayerA
	}

	ships := rd.GetShips(player)

	return server.NewResponse(map[string]interface{}{
		"ships": ships,
		"phase": rd.Phase,
	})
}

// WebSocket handlers

var upgrader = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

// liveHandler gerencia conexões WebSocket para uma partida específica
func (b *Battleship) liveHandler(w http.ResponseWriter, r *http.Request) {
	rd, err := b.getRound(server.GetURLParamHasInt(r, "round") - 1)
	if err != nil {
		http.Error(w, "round not found", http.StatusNotFound)
		return
	}

	player := server.GetURLParamWithDefault(r, "player", "")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "upgrade error", http.StatusInternalServerError)
		return
	}

	bp := rd.AddConnection(conn, player)

	// Se usuário pediu A/B mas slot ocupado -> rejeitar ao invés de rebaixar silenciosamente para spectator
	if (player == "a" || player == "b") && (bp.Player == "" && !bp.Spectator) {
		msg := map[string]interface{}{"type": "error", "action": "connect", "message": "slot ocupado", "requested": player}
		if data, _ := json.Marshal(msg); data != nil {
			_ = conn.WriteMessage(websocket.TextMessage, data)
		}
		conn.Close()
		return
	}

	// Basear estado inicial no papel realmente atribuído (bp) e não apenas no parâmetro solicitado
	var state interface{}
	if bp.Spectator {
		state = rd.createSpectatorState()
	} else if bp.Player == "a" {
		state = rd.createPlayerState(PlayerA)
	} else if bp.Player == "b" {
		state = rd.createPlayerState(PlayerB)
	} else { // fallback teórico
		state = rd.createSpectatorState()
	}
	if m, ok := state.(map[string]interface{}); ok {
		m["type"] = "state"
		if data, mErr := json.Marshal(m); mErr == nil {
			_ = conn.WriteMessage(websocket.TextMessage, data)
		}
	} else if data, mErr := json.Marshal(state); mErr == nil {
		_ = conn.WriteMessage(websocket.TextMessage, data)
	}
	// Broadcast global para alinhar todos (sem segurar lock externo evitando deadlock reentrante)
	rd.BroadcastLocked()

	// Atualizar lista de partidas abertas
	b.broadcastOpenRounds()

	// Goroutine para ler comandos
	go func(c *websocket.Conn, round *Round, bp *BattlePlayer) {
		defer func() {
			c.Close()
			// marcar desconexão
			bp.Conn = nil
			round.BroadcastLocked()
		}()

		// util local
		send := func(v interface{}) {
			if data, err := json.Marshal(v); err == nil {
				_ = c.WriteMessage(websocket.TextMessage, data)
			}
		}

		for {
			_, raw, err := c.ReadMessage()
			if err != nil {
				return
			}
			if len(raw) == 0 {
				continue
			}
			var msg map[string]interface{}
			if err := json.Unmarshal(raw, &msg); err != nil {
				continue
			}
			typeStr, _ := msg["type"].(string)
			id, _ := msg["id"].(string)
			if typeStr == "ping" {
				send(map[string]interface{}{"type": "pong", "id": id, "ts": time.Now().UnixMilli()})
				continue
			}
			// Determinar jogador desta conexão
			var p Player
			if bp.Player == "a" {
				p = PlayerA
			} else if bp.Player == "b" {
				p = PlayerB
			}
			switch typeStr {
			case "shoot":
				if p == Player("") {
					send(map[string]interface{}{"type": "error", "id": id, "action": "shoot", "message": "apenas jogador pode atirar"})
					continue
				}
				rowF, _ := msg["row"].(float64)
				colF, _ := msg["col"].(float64)
				shot, err := round.Shoot(p, int(rowF), int(colF))
				if err != nil {
					send(map[string]interface{}{"type": "error", "id": id, "action": "shoot", "message": err.Error()})
					continue
				}
				state := round.createPlayerState(p)
				send(map[string]interface{}{"type": "shootResult", "id": id, "ok": true, "shot": shot, "state": state})
				round.BroadcastLocked()
			case "placeShip":
				if p == Player("") {
					send(map[string]interface{}{"type": "error", "id": id, "action": "placeShip", "message": "apenas jogador pode posicionar"})
					continue
				}
				shipIdF, _ := msg["shipId"].(float64)
				rowF, _ := msg["row"].(float64)
				colF, _ := msg["col"].(float64)
				orientation, _ := msg["orientation"].(string)
				orient := OrientationHorizontal
				if orientation == "vertical" {
					orient = OrientationVertical
				}
				if err := round.PlacePlayerShip(p, int(shipIdF), int(rowF), int(colF), orient); err != nil {
					send(map[string]interface{}{"type": "error", "id": id, "action": "placeShip", "message": err.Error()})
					continue
				}
				state := round.createPlayerState(p)
				send(map[string]interface{}{"type": "placeResult", "id": id, "ok": true, "state": state, "ready": map[string]bool{"a": round.PlayerAReady, "b": round.PlayerBReady}})
				round.BroadcastLocked()
			default:
				// ignorar silenciosamente tipos desconhecidos
			}
		}
	}(conn, rd, bp)
}

// openLiveHandler gerencia WebSocket para lista de partidas abertas
func (b *Battleship) openLiveHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "upgrade error", http.StatusInternalServerError)
		return
	}

	b.openWatchLock.Lock()
	b.openWatchers = append(b.openWatchers, conn)
	b.openWatchLock.Unlock()

	// Enviar lista inicial
	b.broadcastOpenRounds()

	go func(c *websocket.Conn) {
		defer c.Close()
		for {
			if _, _, err := c.ReadMessage(); err != nil {
				return
			}
		}
	}(conn)
}

// collectOpenRounds coleta partidas em aberto
func (b *Battleship) collectOpenRounds() []map[string]interface{} {
	res := []map[string]interface{}{}
	for _, rd := range b.Rounds {
		if rd == nil || rd.Winner != "" {
			continue
		}
		if rd.PlayerA == "" || rd.PlayerB == "" {
			res = append(res, map[string]interface{}{
				"round":      rd.Round,
				"hasPlayerA": rd.PlayerA != "",
				"hasPlayerB": rd.PlayerB != "",
			})
		}
	}
	return res
}

// broadcastOpenRounds transmite lista de partidas abertas
func (b *Battleship) broadcastOpenRounds() {
	b.openWatchLock.Lock()
	defer b.openWatchLock.Unlock()

	if len(b.openWatchers) == 0 {
		return
	}

	payload := map[string]interface{}{"rounds": b.collectOpenRounds()}
	data, _ := json.Marshal(payload)

	alive := make([]*websocket.Conn, 0, len(b.openWatchers))
	for _, c := range b.openWatchers {
		if c == nil {
			continue
		}
		if err := c.WriteMessage(websocket.TextMessage, data); err != nil {
			c.Close()
			continue
		}
		alive = append(alive, c)
	}
	b.openWatchers = alive
}

// Helper function para parsing seguro de inteiros
func parseIntParam(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if parsed, err := strconv.Atoi(value); err == nil {
		return parsed
	}
	return defaultValue
}
