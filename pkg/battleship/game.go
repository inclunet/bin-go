package battleship

// CellState representa o estado de uma célula no tabuleiro
type CellState string

const (
	CellEmpty CellState = "empty" // Célula vazia
	CellShip  CellState = "ship"  // Navio posicionado
	CellHit   CellState = "hit"   // Navio atingido
	CellMiss  CellState = "miss"  // Tiro na água
	CellSunk  CellState = "sunk"  // Navio afundado
)

// GamePhase representa a fase atual do jogo
type GamePhase string

const (
	PhaseSetup    GamePhase = "setup"    // Posicionamento de navios
	PhasePlaying  GamePhase = "playing"  // Jogando
	PhaseFinished GamePhase = "finished" // Jogo terminado
)

// Player representa um jogador
type Player string

const (
	PlayerA Player = "a"
	PlayerB Player = "b"
)

// Orientation representa a orientação de um navio
type Orientation string

const (
	OrientationHorizontal Orientation = "horizontal"
	OrientationVertical   Orientation = "vertical"
)

// Cell representa uma célula no tabuleiro
type Cell struct {
	State  CellState `json:"state"`
	ShipID int       `json:"shipId,omitempty"`
}

// Board representa um tabuleiro 10x10
type Board [10][10]Cell

// Ship representa um navio
type Ship struct {
	ID          int         `json:"id"`
	Name        string      `json:"name"`
	Size        int         `json:"size"`
	Row         int         `json:"row"`
	Col         int         `json:"col"`
	Orientation Orientation `json:"orientation"`
	Hits        int         `json:"hits"`
	Sunk        bool        `json:"sunk"`
	Placed      bool        `json:"placed"`
}

// Shot representa um tiro
type Shot struct {
	Row    int    `json:"row"`
	Col    int    `json:"col"`
	Result string `json:"result"` // "hit", "miss", "sunk"
}

// GameState representa o estado completo do jogo
type GameState struct {
	BoardA         Board     `json:"boardA"`
	BoardB         Board     `json:"boardB"`
	ShipsA         []Ship    `json:"shipsA"`
	ShipsB         []Ship    `json:"shipsB"`
	CurrentPlayer  Player    `json:"currentPlayer"`
	Phase          GamePhase `json:"phase"`
	Winner         string    `json:"winner"`
	PlayerAReady   bool      `json:"playerAReady"`
	PlayerBReady   bool      `json:"playerBReady"`
	LastShot       *Shot     `json:"lastShot,omitempty"`
	ShotsRemaining int       `json:"shotsRemaining"`
}

// GetDefaultShips retorna a configuração padrão de navios
func GetDefaultShips() []Ship {
	return []Ship{
		{ID: 1, Name: "Porta-aviões", Size: 5},
		{ID: 2, Name: "Encouraçado", Size: 4},
		{ID: 3, Name: "Cruzador", Size: 3},
		{ID: 4, Name: "Submarino", Size: 3},
		{ID: 5, Name: "Destroyer", Size: 2},
	}
}

// NewBoard cria um novo tabuleiro vazio
func NewBoard() Board {
	var board Board
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			board[i][j] = Cell{State: CellEmpty}
		}
	}
	return board
}

// IsValidPosition verifica se uma posição é válida no tabuleiro
func IsValidPosition(row, col int) bool {
	return row >= 0 && row < 10 && col >= 0 && col < 10
}

// GetShipCells retorna as células ocupadas por um navio
func GetShipCells(ship Ship) []struct{ Row, Col int } {
	var cells []struct{ Row, Col int }

	for i := 0; i < ship.Size; i++ {
		if ship.Orientation == OrientationHorizontal {
			cells = append(cells, struct{ Row, Col int }{ship.Row, ship.Col + i})
		} else {
			cells = append(cells, struct{ Row, Col int }{ship.Row + i, ship.Col})
		}
	}

	return cells
}

// CanPlaceShip verifica se um navio pode ser posicionado
func CanPlaceShip(board Board, ship Ship) bool {
	cells := GetShipCells(ship)

	for _, cell := range cells {
		// Verifica se está dentro do tabuleiro
		if !IsValidPosition(cell.Row, cell.Col) {
			return false
		}

		// Verifica se a célula está vazia
		if board[cell.Row][cell.Col].State != CellEmpty {
			return false
		}
	}

	return true
}

// PlaceShip posiciona um navio no tabuleiro
func PlaceShip(board *Board, ship Ship) bool {
	if !CanPlaceShip(*board, ship) {
		return false
	}

	cells := GetShipCells(ship)
	for _, cell := range cells {
		board[cell.Row][cell.Col] = Cell{
			State:  CellShip,
			ShipID: ship.ID,
		}
	}

	return true
}

// Shoot executa um tiro no tabuleiro
func Shoot(board *Board, ships *[]Ship, row, col int) Shot {
	shot := Shot{Row: row, Col: col}

	if !IsValidPosition(row, col) {
		shot.Result = "invalid"
		return shot
	}

	cell := &board[row][col]

	// Se já foi atacado antes
	if cell.State == CellHit || cell.State == CellMiss || cell.State == CellSunk {
		shot.Result = "already_shot"
		return shot
	}

	if cell.State == CellShip {
		// Acerto!
		cell.State = CellHit
		shot.Result = "hit"

		// Verificar se o navio afundou
		for i := range *ships {
			ship := &(*ships)[i]
			if ship.ID == cell.ShipID {
				ship.Hits++
				if ship.Hits >= ship.Size {
					ship.Sunk = true
					shot.Result = "sunk"

					// Marcar todas as células do navio como afundado
					cells := GetShipCells(*ship)
					for _, shipCell := range cells {
						board[shipCell.Row][shipCell.Col].State = CellSunk
					}
				}
				break
			}
		}
	} else {
		// Erro
		cell.State = CellMiss
		shot.Result = "miss"
	}

	return shot
}

// AllShipsSunk verifica se todos os navios foram afundados
func AllShipsSunk(ships []Ship) bool {
	for _, ship := range ships {
		if !ship.Sunk {
			return false
		}
	}
	return true
}

// CountRemainingShips conta quantos navios ainda estão ativos
func CountRemainingShips(ships []Ship) int {
	count := 0
	for _, ship := range ships {
		if !ship.Sunk {
			count++
		}
	}
	return count
}
