// Pequeno harness para testes de broadcast sem depender do componente +page.svelte
// Fornece função pura que aplica parte do payload do servidor sobre estado local
// Uso focado em testes de navio afundado e lastShot/lastSunkShip

export interface Cell { state: string; shipId?: number; name?: string }
export type Board = Cell[][];

export interface BroadcastPayload {
  enemyBoard?: Board;
  myBoard?: Board;
  lastShot?: { row: number; col: number; result: string };
  lastSunkShip?: { id?: number; size?: number; cells: { row: number; col: number }[] };
}

export interface LocalState {
  enemyBoard: Board;
  myBoard: Board;
  lastAction: string;
  lastSunkShip: any;
}

export function cloneBoard(b: Board): Board {
  return b.map(r => r.map(c => ({ ...c })));
}

export function simulateBroadcast(prev: LocalState, payload: BroadcastPayload): LocalState {
  const next: LocalState = {
    enemyBoard: prev.enemyBoard,
    myBoard: prev.myBoard,
    lastAction: prev.lastAction,
    lastSunkShip: prev.lastSunkShip
  };

  if (payload.enemyBoard) next.enemyBoard = cloneBoard(payload.enemyBoard);
  if (payload.myBoard) next.myBoard = cloneBoard(payload.myBoard);

  if (payload.lastShot) {
    const { row, col, result } = payload.lastShot;
    next.lastAction = `${result} em ${coord(row, col)}`;
  }
  if (payload.lastSunkShip) {
    next.lastSunkShip = payload.lastSunkShip;
    if (!payload.lastShot) {
      // ainda registrar uma ação genérica caso não venha lastShot
      next.lastAction = `sunk em (${payload.lastSunkShip.cells.length} células)`;
    }
  }

  return next;
}

export function coord(row: number, col: number): string {
  return String.fromCharCode(65 + col) + (row + 1);
}
