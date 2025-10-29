package battleship

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/inclunet/bin-go/pkg/server"
)

// Round representa uma partida de Batalha Naval
type Round struct {
	Round         int       `json:"round"`
	Phase         GamePhase `json:"phase"`
	CurrentPlayer Player    `json:"currentPlayer"`
	Winner        string    `json:"winner"`
	Next          int       `json:"next,omitempty"`

	// Estados dos jogadores
	PlayerA      string `json:"playerA"` // "A" se ocupado
	PlayerB      string `json:"playerB"` // "B" se ocupado
	PlayerAReady bool   `json:"playerAReady"`
	PlayerBReady bool   `json:"playerBReady"`

	// Tabuleiros (privados - não serializados)
	boardA Board `json:"-"`
	boardB Board `json:"-"`

	// Navios
	shipsA []Ship `json:"-"`
	shipsB []Ship `json:"-"`

	// Estatísticas
	ScoreA    int `json:"scoreA,omitempty"`
	ScoreB    int `json:"scoreB,omitempty"`
	ScoreDraw int `json:"scoreDraw,omitempty"`

	// Último tiro e último navio afundado
	LastShot     *Shot         `json:"lastShot,omitempty"`
	LastSunkShip *SunkShipInfo `json:"lastSunkShip,omitempty"`
	// Controle interno
	counted bool            `json:"-"`
	players []*BattlePlayer `json:"-"`
	lock    sync.Mutex      `json:"-"`
}

// BattlePlayer representa um jogador conectado via WebSocket
type BattlePlayer struct {
	Conn      *websocket.Conn
	Player    string // "a" | "b" | "" (espectador)
	Spectator bool
	ID        int
}

// SunkShipInfo contém dados do último navio afundado (para frontend acessível)
type SunkShipInfo struct {
	ID    int `json:"id"`
	Size  int `json:"size"`
	Cells []struct {
		Row int `json:"row"`
		Col int `json:"col"`
	} `json:"cells"`
}

// NewRound cria uma nova partida
func NewRound(roundNumber int) *Round {
	return &Round{
		Round:         roundNumber,
		Phase:         PhaseSetup,
		CurrentPlayer: PlayerA,
		boardA:        NewBoard(),
		boardB:        NewBoard(),
		shipsA:        GetDefaultShips(),
		shipsB:        GetDefaultShips(),
		players:       []*BattlePlayer{},
	}
}

// GetBoard retorna o tabuleiro de um jogador (apenas para o próprio jogador)
func (r *Round) GetBoard(player Player) Board {
	r.lock.Lock()
	defer r.lock.Unlock()

	if player == PlayerA {
		return r.boardA
	}
	return r.boardB
}

// GetEnemyBoard retorna o tabuleiro do inimigo (apenas células reveladas)
func (r *Round) GetEnemyBoard(player Player) Board {
	r.lock.Lock()
	defer r.lock.Unlock()

	var enemyBoard Board
	if player == PlayerA {
		enemyBoard = r.boardB
	} else {
		enemyBoard = r.boardA
	}

	// Criar uma versão "censurada" do tabuleiro inimigo
	var visibleBoard Board
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			cell := enemyBoard[i][j]
			switch cell.State {
			case CellEmpty, CellShip:
				// Ocultar navios não atingidos
				visibleBoard[i][j] = Cell{State: CellEmpty}
			case CellHit, CellMiss, CellSunk:
				// Mostrar resultados de tiros
				visibleBoard[i][j] = cell
			}
		}
	}

	return visibleBoard
}

// GetShips retorna os navios de um jogador
func (r *Round) GetShips(player Player) []Ship {
	r.lock.Lock()
	defer r.lock.Unlock()

	if player == PlayerA {
		return r.shipsA
	}
	return r.shipsB
}

// PlacePlayerShip posiciona um navio para um jogador
func (r *Round) PlacePlayerShip(player Player, shipID int, row, col int, orientation Orientation) error {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.Phase != PhaseSetup {
		return fmt.Errorf("não é possível posicionar navios fora da fase de setup")
	}

	var board *Board
	var ships *[]Ship

	if player == PlayerA {
		board = &r.boardA
		ships = &r.shipsA
	} else {
		board = &r.boardB
		ships = &r.shipsB
	}

	// Encontrar o navio
	shipIndex := -1
	for i, ship := range *ships {
		if ship.ID == shipID {
			shipIndex = i
			break
		}
	}

	if shipIndex == -1 {
		return fmt.Errorf("navio não encontrado")
	}

	ship := &(*ships)[shipIndex]
	ship.Row = row
	ship.Col = col
	ship.Orientation = orientation

	// Verificar se pode posicionar
	if !CanPlaceShip(*board, *ship) {
		return fmt.Errorf("não é possível posicionar o navio nesta posição")
	}

	// Remover posicionamento anterior se existir
	r.removeShipFromBoard(board, shipID)

	// Posicionar o navio
	if !PlaceShip(board, *ship) {
		return fmt.Errorf("erro ao posicionar o navio")
	}

	// Marcar como posicionado
	ship.Placed = true

	// Verificar se todos os navios foram posicionados
	r.checkPlayerReady(player)

	return nil
}

// removeShipFromBoard remove um navio específico do tabuleiro
func (r *Round) removeShipFromBoard(board *Board, shipID int) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board[i][j].ShipID == shipID {
				board[i][j] = Cell{State: CellEmpty}
			}
		}
	}
}

// checkPlayerReady verifica se o jogador terminou de posicionar todos os navios
func (r *Round) checkPlayerReady(player Player) {
	var ships []Ship
	if player == PlayerA {
		ships = r.shipsA
	} else {
		ships = r.shipsB
	}

	placedCount := 0
	for _, ship := range ships {
		if ship.Placed {
			placedCount++
		}
	}
	allPlaced := placedCount == len(ships) && len(ships) > 0

	if player == PlayerA {
		r.PlayerAReady = allPlaced
	} else {
		r.PlayerBReady = allPlaced
	}

	prePhase := r.Phase
	if r.PlayerAReady && r.PlayerBReady && r.Phase == PhaseSetup {
		r.Phase = PhasePlaying
		r.CurrentPlayer = PlayerA // A começa
	}
	// Logging detalhado de readiness e possível transição
	server.Logger.Info("Battleship ReadyCheck",
		"round", r.Round,
		"player", player,
		"placedCount", placedCount,
		"totalShips", len(ships),
		"playerAReady", r.PlayerAReady,
		"playerBReady", r.PlayerBReady,
		"phaseBefore", prePhase,
		"phaseAfter", r.Phase,
	)
}

// Shoot executa um tiro
func (r *Round) Shoot(attacker Player, row, col int) (*Shot, error) {
	r.lock.Lock()
	defer r.lock.Unlock()

	if r.Phase != PhasePlaying {
		return nil, fmt.Errorf("não é possível atirar fora da fase de jogo")
	}

	if r.Winner != "" {
		return nil, fmt.Errorf("jogo já terminou")
	}

	if r.CurrentPlayer != attacker {
		return nil, fmt.Errorf("não é sua vez de atirar")
	}

	// Determinar tabuleiro e navios alvo
	var targetBoard *Board
	var targetShips *[]Ship

	if attacker == PlayerA {
		targetBoard = &r.boardB
		targetShips = &r.shipsB
	} else {
		targetBoard = &r.boardA
		targetShips = &r.shipsA
	}

	// Executar o tiro
	shot := Shoot(targetBoard, targetShips, row, col)
	r.LastShot = &shot
	// limpar referência anterior; será reatribuída se este tiro afundar algo
	r.LastSunkShip = nil

	if shot.Result == "invalid" || shot.Result == "already_shot" {
		return &shot, fmt.Errorf("tiro inválido: %s", shot.Result)
	}

	// Se afundou, identificar navio e registrar células
	if shot.Result == "sunk" {
		for i := range *targetShips {
			ship := &(*targetShips)[i]
			if ship.Sunk { // pode haver múltiplos já afundados; encontrar o que contém a célula do tiro
				cells := GetShipCells(*ship)
				contains := false
				for _, c := range cells {
					if c.Row == row && c.Col == col {
						contains = true
						break
					}
				}
				if contains {
					info := &SunkShipInfo{ID: ship.ID, Size: ship.Size}
					for _, c := range cells {
						info.Cells = append(info.Cells, struct {
							Row int `json:"row"`
							Col int `json:"col"`
						}{Row: c.Row, Col: c.Col})
					}
					r.LastSunkShip = info
					break
				}
			}
		}
	}

	// Verificar vitória
	if AllShipsSunk(*targetShips) {
		r.Winner = string(attacker)
		r.Phase = PhaseFinished

		if !r.counted {
			// NÃO limpar LastSunkShip aqui: precisamos comunicar ao cliente qual foi o navio final afundado
			r.counted = true
			if attacker == PlayerA {
				r.ScoreA++
			} else {
				r.ScoreB++
			}
		}

		server.Logger.Info("Battleship Winner", "round", r.Round, "winner", r.Winner)
	} else if shot.Result == "miss" {
		// Apenas alterna o turno se errou o tiro
		if r.CurrentPlayer == PlayerA {
			r.CurrentPlayer = PlayerB
		} else {
			r.CurrentPlayer = PlayerA
		}
	}
	// Se acertou (hit ou sunk), mantém o turno

	return &shot, nil
}

// AddConnection adiciona uma conexão WebSocket
func (r *Round) AddConnection(conn *websocket.Conn, requestedPlayer string) *BattlePlayer {
	r.lock.Lock()
	defer r.lock.Unlock()

	// Antes de alocar slot, tentar liberar conexões zumbis para evitar fallback indevido para spectator
	r.compactPlayersLocked()

	player := ""
	spectator := false

	switch requestedPlayer {
	case "a", "A":
		if r.PlayerA == "" { // slot livre
			r.PlayerA = "A"
			player = "a"
		} else {
			// NÃO força spectator silencioso; mantém player vazio para sinalizar rejeição (decisão no handler)
		}
	case "b", "B":
		if r.PlayerB == "" {
			r.PlayerB = "B"
			player = "b"
		} else {
			// idem acima: não rebaixa automaticamente para spectator
		}
	default:
		spectator = true
	}

	// Se o usuário pediu explicitamente A ou B e o slot está ocupado, deixamos player vazio e spectator=false.
	// O handler decidirá encerrar a conexão com erro para evitar experiência confusa de spectator oculto.

	id := len(r.players) + 1 // id simples incremental (não reutilizado)
	bp := &BattlePlayer{Conn: conn, Player: player, Spectator: spectator, ID: id}
	r.players = append(r.players, bp)

	server.Logger.Info("Battleship Connect", "round", r.Round, "player", player, "spectator", spectator, "connId", id, "players", len(r.players))
	return bp
}

// BroadcastLocked envia o estado atual para todos os jogadores conectados
func (r *Round) BroadcastLocked() {
	// Estratégia: snapshot atômico sob lock -> envia fora do lock minimizando contenção e evitando reentrância.
	r.lock.Lock()
	players := make([]*BattlePlayer, len(r.players))
	copy(players, r.players)

	// Funções utilitárias internas (sem novo lock)
	maskEnemy := func(enemy Board) Board {
		var vis Board
		for i := 0; i < 10; i++ {
			for j := 0; j < 10; j++ {
				c := enemy[i][j]
				switch c.State {
				case CellHit, CellMiss, CellSunk:
					vis[i][j] = c
				default:
					vis[i][j] = Cell{State: CellEmpty}
				}
			}
		}
		return vis
	}

	build := func(player Player, bp *BattlePlayer) map[string]interface{} {
		// Base comum
		st := map[string]interface{}{
			"round":         r.Round,
			"phase":         r.Phase,
			"currentPlayer": r.CurrentPlayer,
			"winner":        r.Winner,
			"playerA":       r.PlayerA,
			"playerB":       r.PlayerB,
			"playerAReady":  r.PlayerAReady,
			"playerBReady":  r.PlayerBReady,
			"scoreA":        r.ScoreA,
			"scoreB":        r.ScoreB,
			"scoreDraw":     r.ScoreDraw,
			"lastShot":      r.LastShot,
			"lastSunkShip":  r.LastSunkShip,
			"shipsRemaining": map[string]int{
				"a": CountRemainingShips(r.shipsA),
				"b": CountRemainingShips(r.shipsB),
			},
		}
		if bp != nil && bp.Spectator {
			st["spectator"] = true
			return st
		}
		// Player específico
		var myBoard, enemyBoard Board
		var myShips []Ship
		if player == PlayerA {
			myBoard = r.boardA
			enemyBoard = maskEnemy(r.boardB)
			myShips = append(myShips, r.shipsA...)
		} else if player == PlayerB {
			myBoard = r.boardB
			enemyBoard = maskEnemy(r.boardA)
			myShips = append(myShips, r.shipsB...)
		}
		st["myBoard"] = myBoard
		st["enemyBoard"] = enemyBoard
		st["myShips"] = myShips
		return st
	}

	// Pré-montar payloads
	server.Logger.Info("Battleship Broadcast", "round", r.Round, "phase", r.Phase, "players", len(players))
	type outMsg struct {
		p    *BattlePlayer
		data []byte
	}
	messages := []outMsg{}
	for _, bp := range players {
		if bp == nil || bp.Conn == nil {
			continue
		}
		var payload map[string]interface{}
		switch bp.Player {
		case "a":
			payload = build(PlayerA, bp)
		case "b":
			payload = build(PlayerB, bp)
		default:
			payload = build(Player(""), &BattlePlayer{Spectator: true})
		}
		// Tipar broadcast
		payload["type"] = "state"
		if encoded, err := json.Marshal(payload); err == nil {
			messages = append(messages, outMsg{p: bp, data: encoded})
			server.Logger.Info("Battleship BroadcastPlayer", "round", r.Round, "connId", bp.ID, "player", bp.Player, "spectator", bp.Spectator, "phase", r.Phase, "bytes", len(encoded))
		} else {
			server.Logger.Error("Battleship BroadcastMarshalError", "round", r.Round, "connId", bp.ID, "err", err)
		}
	}
	r.lock.Unlock()

	// Enviar fora do lock
	for _, m := range messages {
		// Capturar referência local para evitar race se m.p.Conn for setado para nil entre iterações
		conn := m.p.Conn
		if conn == nil {
			continue
		}
		err := conn.WriteMessage(websocket.TextMessage, m.data)
		if err != nil {
			server.Logger.Error("Battleship BroadcastWriteError", "round", r.Round, "connId", m.p.ID, "player", m.p.Player, "err", err)
			_ = conn.Close()
			m.p.Conn = nil
		} else {
			server.Logger.Info("Battleship BroadcastWriteOk", "round", r.Round, "connId", m.p.ID, "player", m.p.Player)
		}
	}

	// Recompactar após envios (sob lock breve)
	r.lock.Lock()
	r.compactPlayersLocked()
	r.lock.Unlock()
}

// createPlayerState cria o estado específico para um jogador
func (r *Round) createPlayerState(player Player) map[string]interface{} {
	state := map[string]interface{}{
		"round":         r.Round,
		"phase":         r.Phase,
		"currentPlayer": r.CurrentPlayer,
		"winner":        r.Winner,
		"playerA":       r.PlayerA,
		"playerB":       r.PlayerB,
		"playerAReady":  r.PlayerAReady,
		"playerBReady":  r.PlayerBReady,
		"scoreA":        r.ScoreA,
		"scoreB":        r.ScoreB,
		"scoreDraw":     r.ScoreDraw,
		"lastShot":      r.LastShot,
		"lastSunkShip":  r.LastSunkShip,
		"myBoard":       r.GetBoard(player),
		"enemyBoard":    r.GetEnemyBoard(player),
		"myShips":       r.GetShips(player),
		"shipsRemaining": map[string]int{
			"a": CountRemainingShips(r.shipsA),
			"b": CountRemainingShips(r.shipsB),
		},
	}

	return state
}

// createSpectatorState cria o estado para espectadores
func (r *Round) createSpectatorState() map[string]interface{} {
	return map[string]interface{}{
		"round":         r.Round,
		"phase":         r.Phase,
		"currentPlayer": r.CurrentPlayer,
		"winner":        r.Winner,
		"playerA":       r.PlayerA,
		"playerB":       r.PlayerB,
		"playerAReady":  r.PlayerAReady,
		"playerBReady":  r.PlayerBReady,
		"scoreA":        r.ScoreA,
		"scoreB":        r.ScoreB,
		"scoreDraw":     r.ScoreDraw,
		"lastShot":      r.LastShot,
		"lastSunkShip":  r.LastSunkShip,
		"spectator":     true,
		"shipsRemaining": map[string]int{
			"a": CountRemainingShips(r.shipsA),
			"b": CountRemainingShips(r.shipsB),
		},
	}
}

// compactPlayersLocked remove jogadores desconectados
func (r *Round) compactPlayersLocked() {
	cleaned := make([]*BattlePlayer, 0, len(r.players))
	for _, p := range r.players {
		if p != nil && p.Conn != nil {
			cleaned = append(cleaned, p)
		}
	}
	// Atualiza lista compactada
	r.players = cleaned

	// Verificar se ainda existem conexões ativas para cada slot de jogador.
	// Caso não exista mais conexão para 'a' ou 'b', liberamos o slot para permitir reconexão.
	hasA := false
	hasB := false
	for _, p := range r.players {
		if p == nil {
			continue
		}
		if p.Player == "a" {
			hasA = true
		}
		if p.Player == "b" {
			hasB = true
		}
	}
	if !hasA && r.PlayerA != "" {
		r.PlayerA = "" // libera slot A
		server.Logger.Info("Battleship SlotCleared", "round", r.Round, "player", "a")
	}
	if !hasB && r.PlayerB != "" {
		r.PlayerB = "" // libera slot B
		server.Logger.Info("Battleship SlotCleared", "round", r.Round, "player", "b")
	}
}
