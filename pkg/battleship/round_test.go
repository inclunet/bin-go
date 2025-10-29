package battleship

import (
	"testing"
)

// ---- Novos testes para lastSunkShip e transição de fase sem fallback ----

// helper para posicionar todos os navios de um jogador em linhas separadas
func placeAllShips(t *testing.T, r *Round, p Player) {
	row := 0
	for _, s := range GetDefaultShips() {
		if err := r.PlacePlayerShip(p, s.ID, row, 0, OrientationHorizontal); err != nil {
			t.Fatalf("falha ao posicionar navio %d: %v", s.ID, err)
		}
		row++
	}
}

func TestPhasePlayingWhenBothReady_NoClientFallback(t *testing.T) {
	r := NewRound(99)
	placeAllShips(t, r, PlayerA)
	if r.Phase != PhaseSetup {
		t.Fatalf("esperava setup depois de só A pronto, got %s", r.Phase)
	}
	placeAllShips(t, r, PlayerB)
	if r.Phase != PhasePlaying {
		t.Fatalf("esperava playing após ambos prontos, got %s", r.Phase)
	}
}

func TestLastSunkShipInfoProvided(t *testing.T) {
	r := NewRound(77)
	placeAllShips(t, r, PlayerA)
	placeAllShips(t, r, PlayerB)
	if r.Phase != PhasePlaying {
		t.Fatalf("fase não mudou para playing")
	}

	// localizar destroyer (size 2) de B
	var destroyer Ship
	for _, s := range r.shipsB {
		if s.Size == 2 {
			destroyer = s
			break
		}
	}
	if destroyer.ID == 0 {
		t.Fatalf("destroyer não encontrado")
	}

	cells := GetShipCells(destroyer)
	if len(cells) != destroyer.Size {
		t.Fatalf("esperava %d células, got %d", destroyer.Size, len(cells))
	}

	shot1, err := r.Shoot(PlayerA, cells[0].Row, cells[0].Col)
	if err != nil {
		t.Fatalf("erro primeiro tiro: %v", err)
	}
	if shot1.Result != "hit" {
		t.Fatalf("esperava hit, got %s", shot1.Result)
	}
	if r.LastSunkShip != nil {
		t.Fatalf("não deveria haver LastSunkShip ainda")
	}

	shot2, err := r.Shoot(PlayerA, cells[1].Row, cells[1].Col)
	if err != nil {
		t.Fatalf("erro segundo tiro: %v", err)
	}
	if shot2.Result != "sunk" {
		t.Fatalf("esperava sunk, got %s", shot2.Result)
	}
	if r.LastSunkShip == nil {
		t.Fatalf("LastSunkShip deveria estar preenchido")
	}
	if r.LastSunkShip.Size != destroyer.Size {
		t.Fatalf("size esperado %d got %d", destroyer.Size, r.LastSunkShip.Size)
	}
	if len(r.LastSunkShip.Cells) != destroyer.Size {
		t.Fatalf("cells esperadas %d got %d", destroyer.Size, len(r.LastSunkShip.Cells))
	}
	for _, c := range r.LastSunkShip.Cells {
		if r.boardB[c.Row][c.Col].State != CellSunk {
			t.Fatalf("célula %d,%d deveria estar sunk", c.Row, c.Col)
		}
	}
}

func TestNewRound(t *testing.T) {
	round := NewRound(1)

	if round.Round != 1 {
		t.Errorf("Expected round number 1, got %d", round.Round)
	}

	if round.Phase != PhaseSetup {
		t.Errorf("Expected setup phase, got %s", round.Phase)
	}

	if round.CurrentPlayer != PlayerA {
		t.Errorf("Expected player A to start, got %s", round.CurrentPlayer)
	}

	if len(round.shipsA) != 5 || len(round.shipsB) != 5 {
		t.Error("Should have 5 ships for each player")
	}

	// Verificar tabuleiros vazios
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if round.boardA[i][j].State != CellEmpty {
				t.Errorf("Board A should be empty at (%d,%d)", i, j)
			}
			if round.boardB[i][j].State != CellEmpty {
				t.Errorf("Board B should be empty at (%d,%d)", i, j)
			}
		}
	}
}

func TestPlacePlayerShip(t *testing.T) {
	round := NewRound(1)

	// Posicionar navio válido
	err := round.PlacePlayerShip(PlayerA, 1, 2, 3, OrientationHorizontal)
	if err != nil {
		t.Errorf("Should be able to place ship: %v", err)
	}

	// Verificar se foi posicionado corretamente
	ship := round.shipsA[0]
	if ship.Row != 2 || ship.Col != 3 || ship.Orientation != OrientationHorizontal {
		t.Error("Ship not positioned correctly")
	}

	// Verificar se o tabuleiro foi atualizado
	for col := 3; col <= 7; col++ { // Porta-aviões tem tamanho 5
		cell := round.boardA[2][col]
		if cell.State != CellShip || cell.ShipID != 1 {
			t.Errorf("Board cell (%d,%d) not marked correctly", 2, col)
		}
	}

	// Tentar posicionar navio sobreposto
	err = round.PlacePlayerShip(PlayerA, 2, 2, 4, OrientationHorizontal)
	if err == nil {
		t.Error("Should not allow overlapping ships")
	}

	// Tentar posicionar fora da fase de setup
	round.Phase = PhasePlaying
	err = round.PlacePlayerShip(PlayerA, 3, 0, 0, OrientationHorizontal)
	if err == nil {
		t.Error("Should not allow placing ships outside setup phase")
	}
}

func TestCheckPlayerReady(t *testing.T) {
	round := NewRound(1)

	// Inicialmente não está pronto
	if round.PlayerAReady {
		t.Error("Player A should not be ready initially")
	}

	// Posicionar todos os navios
	ships := []struct {
		id          int
		row, col    int
		orientation Orientation
	}{
		{1, 0, 0, OrientationHorizontal}, // Porta-aviões (5)
		{2, 1, 0, OrientationHorizontal}, // Encouraçado (4)
		{3, 2, 0, OrientationHorizontal}, // Cruzador (3)
		{4, 3, 0, OrientationHorizontal}, // Submarino (3)
		{5, 4, 0, OrientationHorizontal}, // Destroyer (2)
	}

	for _, ship := range ships {
		err := round.PlacePlayerShip(PlayerA, ship.id, ship.row, ship.col, ship.orientation)
		if err != nil {
			t.Errorf("Error placing ship %d: %v", ship.id, err)
		}
	}

	// Agora deve estar pronto
	if !round.PlayerAReady {
		t.Error("Player A should be ready after placing all ships")
	}

	// Mas o jogo não deve começar ainda (só jogador A pronto)
	if round.Phase != PhaseSetup {
		t.Error("Game should still be in setup phase")
	}

	// Posicionar navios do jogador B
	for _, ship := range ships {
		err := round.PlacePlayerShip(PlayerB, ship.id, ship.row, ship.col+5, ship.orientation)
		if err != nil {
			t.Errorf("Error placing ship %d for player B: %v", ship.id, err)
		}
	}

	// Agora o jogo deve começar
	if round.Phase != PhasePlaying {
		t.Error("Game should be in playing phase after both players ready")
	}

	if round.CurrentPlayer != PlayerA {
		t.Error("Player A should start the game")
	}
}

func TestRoundShoot(t *testing.T) {
	round := NewRound(1)

	// Preparar jogo: posicionar navios e mudar para fase de jogo
	round.PlacePlayerShip(PlayerA, 1, 0, 0, OrientationHorizontal)
	round.PlacePlayerShip(PlayerB, 1, 5, 5, OrientationHorizontal)
	round.Phase = PhasePlaying
	round.PlayerAReady = true
	round.PlayerBReady = true
	round.CurrentPlayer = PlayerA

	// Teste tiro fora da fase de jogo
	round.Phase = PhaseSetup
	_, err := round.Shoot(PlayerA, 1, 1)
	if err == nil {
		t.Error("Should not allow shooting outside playing phase")
	}
	round.Phase = PhasePlaying

	// Teste tiro fora de turno
	_, err = round.Shoot(PlayerB, 1, 1)
	if err == nil {
		t.Error("Should not allow shooting out of turn")
	}

	// Teste tiro válido (erro)
	shot, err := round.Shoot(PlayerA, 1, 1)
	if err != nil {
		t.Errorf("Valid shot should succeed: %v", err)
	}
	if shot.Result != "miss" {
		t.Errorf("Expected miss, got %s", shot.Result)
	}

	// Turno deve ter mudado após erro
	if round.CurrentPlayer != PlayerB {
		t.Error("Turn should change after miss")
	}

	// Teste acerto (mantém turno)
	shot, err = round.Shoot(PlayerB, 0, 0)
	if err != nil {
		t.Errorf("Valid shot should succeed: %v", err)
	}
	if shot.Result != "hit" {
		t.Errorf("Expected hit, got %s", shot.Result)
	}

	// Turno NÃO deve mudar após acerto
	if round.CurrentPlayer != PlayerB {
		t.Error("Turn should not change after hit")
	}

	// Teste tiro em posição já atacada
	_, err = round.Shoot(PlayerB, 0, 0)
	if err == nil {
		t.Error("Should not allow shooting same position twice")
	}
}

func TestGameCompletion(t *testing.T) {
	round := NewRound(1)

	// Preparar jogo com apenas um navio pequeno para cada jogador para facilitar o teste
	round.shipsA = []Ship{{ID: 1, Name: "Test", Size: 1, Row: 0, Col: 0, Orientation: OrientationHorizontal}}
	round.shipsB = []Ship{{ID: 1, Name: "Test", Size: 1, Row: 5, Col: 5, Orientation: OrientationHorizontal}}

	// Posicionar navios nos tabuleiros
	PlaceShip(&round.boardA, round.shipsA[0])
	PlaceShip(&round.boardB, round.shipsB[0])

	round.Phase = PhasePlaying
	round.PlayerAReady = true
	round.PlayerBReady = true
	round.CurrentPlayer = PlayerA

	// Jogador A atira e acerta o único navio de B
	shot, err := round.Shoot(PlayerA, 5, 5)
	if err != nil {
		t.Errorf("Final shot should succeed: %v", err)
	}

	if shot.Result != "sunk" {
		t.Errorf("Expected sunk, got %s", shot.Result)
	}

	// Jogo deve ter terminado
	if round.Winner != string(PlayerA) {
		t.Errorf("Player A should be winner, got %s", round.Winner)
	}

	if round.Phase != PhaseFinished {
		t.Error("Game should be finished")
	}

	// Score deve ter sido incrementado
	if round.ScoreA != 1 {
		t.Errorf("Player A score should be 1, got %d", round.ScoreA)
	}
}

func TestGetEnemyBoard(t *testing.T) {
	round := NewRound(1)

	// Posicionar navio do jogador B
	round.PlacePlayerShip(PlayerB, 1, 2, 3, OrientationHorizontal)

	// Jogador A deve ver tabuleiro "censurado" (sem navios não atingidos)
	enemyBoard := round.GetEnemyBoard(PlayerA)

	// Verificar que navios não atingidos estão ocultos
	for col := 3; col <= 7; col++ {
		if enemyBoard[2][col].State != CellEmpty {
			t.Errorf("Unshot ship should be hidden at (%d,%d)", 2, col)
		}
	}

	// Simular um tiro que acerta
	round.boardB[2][3].State = CellHit
	enemyBoard = round.GetEnemyBoard(PlayerA)

	// Agora deve mostrar o acerto
	if enemyBoard[2][3].State != CellHit {
		t.Error("Hit should be visible on enemy board")
	}

	// Mas partes não atingidas ainda devem estar ocultas
	if enemyBoard[2][4].State != CellEmpty {
		t.Error("Unshot parts should still be hidden")
	}
}

func TestEnemyBoardShowsEntireSunkShip(t *testing.T) {
	r := NewRound(3)
	// Colocar um navio pequeno para B e entrar em playing
	// Para simplificar, reposicionar manualmente só destroyer (size 2)
	// Substituir frota B por um único navio size 2
	r.shipsB = []Ship{{ID: 99, Name: "Duo", Size: 2, Row: 5, Col: 5, Orientation: OrientationHorizontal, Placed: true}}
	PlaceShip(&r.boardB, r.shipsB[0])
	// Jogador A precisa de qualquer navio posicionado para marcar pronto
	r.shipsA = []Ship{{ID: 1, Name: "Solo", Size: 1, Row: 0, Col: 0, Orientation: OrientationHorizontal, Placed: true}}
	PlaceShip(&r.boardA, r.shipsA[0])
	r.PlayerAReady = true
	r.PlayerBReady = true
	r.Phase = PhasePlaying
	r.CurrentPlayer = PlayerA

	cells := GetShipCells(r.shipsB[0])
	// Primeiro tiro (hit)
	shot1, err := r.Shoot(PlayerA, cells[0].Row, cells[0].Col)
	if err != nil {
		t.Fatalf("erro tiro 1: %v", err)
	}
	if shot1.Result != "hit" {
		t.Fatalf("esperava hit got %s", shot1.Result)
	}
	// Segundo tiro (sunk)
	shot2, err := r.Shoot(PlayerA, cells[1].Row, cells[1].Col)
	if err != nil {
		t.Fatalf("erro tiro 2: %v", err)
	}
	if shot2.Result != "sunk" {
		t.Fatalf("esperava sunk got %s", shot2.Result)
	}
	if r.LastSunkShip == nil {
		t.Fatalf("LastSunkShip não preenchido")
	}

	// Visão de A do tabuleiro inimigo deve mostrar ambas as células como sunk
	enemyView := r.GetEnemyBoard(PlayerA)
	for _, c := range cells {
		if enemyView[c.Row][c.Col].State != CellSunk {
			t.Fatalf("enemy view deveria ter sunk em %d,%d", c.Row, c.Col)
		}
	}
}
