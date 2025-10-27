import { describe, it, expect } from 'vitest';
import { simulateBroadcast, coord, type LocalState, type Board } from '../../battleship/broadcastHarness';

function emptyBoard(): Board {
  return Array.from({ length: 10 }, () => Array.from({ length: 10 }, () => ({ state: 'empty' })));
}

describe('simulateBroadcast - navio afundado', () => {
  it('marca lastAction baseado em lastShot e preserva células sunk do payload', () => {
    // Estado inicial com duas células já hit de um navio de 3 (coordenadas (2,2) (2,3)) e terceira ainda hit no payload final
    const initial: LocalState = {
      enemyBoard: emptyBoard(),
      myBoard: emptyBoard(),
      lastAction: '',
      lastSunkShip: null
    };
    initial.enemyBoard[2][2] = { state: 'hit' };
    initial.enemyBoard[2][3] = { state: 'hit' };

    // Payload de broadcast do servidor: navio afundado nas três células (2,2),(2,3),(2,4)
    const payload = {
      enemyBoard: (() => {
        const b = emptyBoard();
        b[2][2] = { state: 'sunk' };
        b[2][3] = { state: 'sunk' };
        b[2][4] = { state: 'sunk' };
        return b;
      })(),
      lastShot: { row: 2, col: 4, result: 'sunk' },
      lastSunkShip: { id: 7, size: 3, cells: [ { row:2, col:2 }, { row:2, col:3 }, { row:2, col:4 } ] }
    };

    const next = simulateBroadcast(initial, payload);

    // Verificações
    expect(next.lastAction).toBe(`sunk em ${coord(2,4)}`);
    expect(next.lastSunkShip).toBeTruthy();
    expect(next.lastSunkShip.size).toBe(3);
    // Todas as células do navio vieram como sunk no payload e foram clonadas
    for (const c of payload.lastSunkShip.cells) {
      expect(next.enemyBoard[c.row][c.col].state).toBe('sunk');
    }
  });

  it('lastSunkShip sem lastShot ainda registra ação genérica', () => {
    const initial: LocalState = {
      enemyBoard: emptyBoard(),
      myBoard: emptyBoard(),
      lastAction: '',
      lastSunkShip: null
    };

    const payload = {
      enemyBoard: emptyBoard(),
      lastSunkShip: { id: 9, size: 2, cells: [ { row:0, col:0 }, { row:0, col:1 } ] }
    };

    const next = simulateBroadcast(initial, payload);
    expect(next.lastAction).toContain('sunk');
    expect(next.lastSunkShip.size).toBe(2);
  });
});
