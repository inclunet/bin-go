package battleship

import (
	"testing"
)

func TestNewBoard(t *testing.T) {
	board := NewBoard()

	// Verificar se o tabuleiro está vazio
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if board[i][j].State != CellEmpty {
				t.Errorf("Expected empty cell at (%d,%d), got %s", i, j, board[i][j].State)
			}
		}
	}
}

func TestIsValidPosition(t *testing.T) {
	tests := []struct {
		row, col int
		expected bool
	}{
		{0, 0, true},
		{9, 9, true},
		{5, 5, true},
		{-1, 0, false},
		{0, -1, false},
		{10, 0, false},
		{0, 10, false},
		{-1, -1, false},
		{10, 10, false},
	}

	for _, test := range tests {
		result := IsValidPosition(test.row, test.col)
		if result != test.expected {
			t.Errorf("IsValidPosition(%d,%d) = %v, expected %v", test.row, test.col, result, test.expected)
		}
	}
}

func TestGetShipCells(t *testing.T) {
	// Teste navio horizontal
	ship := Ship{
		ID:          1,
		Size:        3,
		Row:         2,
		Col:         3,
		Orientation: OrientationHorizontal,
	}

	cells := GetShipCells(ship)
	expected := []struct{ Row, Col int }{
		{2, 3}, {2, 4}, {2, 5},
	}

	if len(cells) != len(expected) {
		t.Errorf("Expected %d cells, got %d", len(expected), len(cells))
	}

	for i, cell := range cells {
		if cell.Row != expected[i].Row || cell.Col != expected[i].Col {
			t.Errorf("Cell %d: expected (%d,%d), got (%d,%d)", i, expected[i].Row, expected[i].Col, cell.Row, cell.Col)
		}
	}

	// Teste navio vertical
	ship.Orientation = OrientationVertical
	cells = GetShipCells(ship)
	expectedVertical := []struct{ Row, Col int }{
		{2, 3}, {3, 3}, {4, 3},
	}

	for i, cell := range cells {
		if cell.Row != expectedVertical[i].Row || cell.Col != expectedVertical[i].Col {
			t.Errorf("Vertical Cell %d: expected (%d,%d), got (%d,%d)", i, expectedVertical[i].Row, expectedVertical[i].Col, cell.Row, cell.Col)
		}
	}
}

func TestCanPlaceShip(t *testing.T) {
	board := NewBoard()

	// Teste posicionamento válido
	ship := Ship{
		ID:          1,
		Size:        3,
		Row:         2,
		Col:         3,
		Orientation: OrientationHorizontal,
	}

	if !CanPlaceShip(board, ship) {
		t.Error("Should be able to place ship in empty area")
	}

	// Posicionar o navio
	PlaceShip(&board, ship)

	// Teste posicionamento sobreposto
	ship2 := Ship{
		ID:          2,
		Size:        2,
		Row:         2,
		Col:         4,
		Orientation: OrientationHorizontal,
	}

	if CanPlaceShip(board, ship2) {
		t.Error("Should not be able to place ship over existing ship")
	}

	// Teste posicionamento fora do tabuleiro
	ship3 := Ship{
		ID:          3,
		Size:        3,
		Row:         9,
		Col:         8,
		Orientation: OrientationHorizontal,
	}

	if CanPlaceShip(board, ship3) {
		t.Error("Should not be able to place ship outside board")
	}
}

func TestPlaceShip(t *testing.T) {
	board := NewBoard()

	ship := Ship{
		ID:          1,
		Size:        3,
		Row:         2,
		Col:         3,
		Orientation: OrientationHorizontal,
	}

	if !PlaceShip(&board, ship) {
		t.Error("Should be able to place valid ship")
	}

	// Verificar se as células foram marcadas corretamente
	expectedCells := GetShipCells(ship)
	for _, cell := range expectedCells {
		boardCell := board[cell.Row][cell.Col]
		if boardCell.State != CellShip {
			t.Errorf("Cell (%d,%d) should be CellShip, got %s", cell.Row, cell.Col, boardCell.State)
		}
		if boardCell.ShipID != ship.ID {
			t.Errorf("Cell (%d,%d) should have ShipID %d, got %d", cell.Row, cell.Col, ship.ID, boardCell.ShipID)
		}
	}
}

func TestShoot(t *testing.T) {
	board := NewBoard()
	ships := GetDefaultShips()

	// Posicionar um navio
	ship := Ship{
		ID:          1,
		Name:        "Test Ship",
		Size:        3,
		Row:         2,
		Col:         3,
		Orientation: OrientationHorizontal,
	}
	ships[0] = ship
	PlaceShip(&board, ship)

	// Teste tiro na água
	shot := Shoot(&board, &ships, 0, 0)
	if shot.Result != "miss" {
		t.Errorf("Expected miss, got %s", shot.Result)
	}
	if board[0][0].State != CellMiss {
		t.Errorf("Cell should be marked as miss")
	}

	// Teste acerto
	shot = Shoot(&board, &ships, 2, 3)
	if shot.Result != "hit" {
		t.Errorf("Expected hit, got %s", shot.Result)
	}
	if board[2][3].State != CellHit {
		t.Errorf("Cell should be marked as hit")
	}
	if ships[0].Hits != 1 {
		t.Errorf("Ship should have 1 hit, got %d", ships[0].Hits)
	}

	// Mais acertos até afundar
	Shoot(&board, &ships, 2, 4)
	shot = Shoot(&board, &ships, 2, 5) // Último tiro para afundar

	if shot.Result != "sunk" {
		t.Errorf("Expected sunk, got %s", shot.Result)
	}
	if !ships[0].Sunk {
		t.Error("Ship should be marked as sunk")
	}

	// Verificar se todas as células do navio foram marcadas como afundado
	for col := 3; col <= 5; col++ {
		if board[2][col].State != CellSunk {
			t.Errorf("Cell (%d,%d) should be marked as sunk", 2, col)
		}
	}

	// Teste tiro em posição já atacada
	shot = Shoot(&board, &ships, 2, 3)
	if shot.Result != "already_shot" {
		t.Errorf("Expected already_shot, got %s", shot.Result)
	}
}

func TestAllShipsSunk(t *testing.T) {
	ships := []Ship{
		{ID: 1, Sunk: true},
		{ID: 2, Sunk: true},
		{ID: 3, Sunk: false},
	}

	if AllShipsSunk(ships) {
		t.Error("Not all ships are sunk")
	}

	ships[2].Sunk = true
	if !AllShipsSunk(ships) {
		t.Error("All ships are sunk")
	}
}

func TestCountRemainingShips(t *testing.T) {
	ships := []Ship{
		{ID: 1, Sunk: true},
		{ID: 2, Sunk: false},
		{ID: 3, Sunk: false},
	}

	count := CountRemainingShips(ships)
	if count != 2 {
		t.Errorf("Expected 2 remaining ships, got %d", count)
	}

	ships[1].Sunk = true
	count = CountRemainingShips(ships)
	if count != 1 {
		t.Errorf("Expected 1 remaining ship, got %d", count)
	}
}
