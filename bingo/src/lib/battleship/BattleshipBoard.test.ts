import { render, screen, cleanup } from '@testing-library/svelte';
import { tick } from 'svelte';
import { describe, it, expect, afterEach } from 'vitest';
import BattleshipBoard from '$lib/battleship/BattleshipBoard.svelte';

function emptyBoard() {
  return Array.from({ length: 10 }, () => Array.from({ length: 10 }, () => ({ state: 'empty' })));
}

function withCell(board: any[][], r: number, c: number, cell: any) {
  const clone = board.map(row => row.map(col => ({ ...col })));
  clone[r][c] = cell;
  return clone;
}

afterEach(() => cleanup());

describe('BattleshipBoard labels acessíveis (gridcell)', () => {
  it('ship', () => {
    let board = emptyBoard();
    board = withCell(board, 0, 0, { state: 'ship', shipId: 1, name: 'Porta-aviões' });
    render(BattleshipBoard, { board, mode: 'setup', isMyTurn: true, row:0, col:0, roundParam:'1', shipNames:{1:'Porta-aviões'} });
    const cell = screen.getByRole('gridcell', { name: 'Porta-aviões, A1' });
    expect(cell).toBeTruthy();
  });
  it('hit', () => {
    let board = emptyBoard();
    board = withCell(board, 1, 1, { state: 'hit' });
    render(BattleshipBoard, { board, mode: 'playing', isMyTurn: true, row:1, col:1, roundParam:'1', shipNames:{} });
    const cell = screen.getByRole('gridcell', { name: 'Acerto, B2' });
    expect(cell).toBeTruthy();
  });
  it('miss', () => {
    let board = emptyBoard();
    board = withCell(board, 2, 2, { state: 'miss' });
    render(BattleshipBoard, { board, mode: 'playing', isMyTurn: true, row:2, col:2, roundParam:'1', shipNames:{} });
    const cell = screen.getByRole('gridcell', { name: 'Água, C3' });
    expect(cell).toBeTruthy();
  });
  it('sunk', () => {
    let board = emptyBoard();
    board = withCell(board, 3, 3, { state: 'sunk' });
    render(BattleshipBoard, { board, mode: 'playing', isMyTurn: true, row:3, col:3, roundParam:'1', shipNames:{} });
    const cell = screen.getByRole('gridcell', { name: 'Afundado, D4' });
    expect(cell).toBeTruthy();
  });
  it('empty', () => {
    let board = emptyBoard();
    render(BattleshipBoard, { board, mode: 'setup', isMyTurn: true, row:4, col:4, roundParam:'1', shipNames:{} });
    const cell = screen.getByRole('gridcell', { name: 'Vazio, E5' });
    expect(cell).toBeTruthy();
  });
});

describe('BattleshipBoard atualização dinâmica de shipNames', () => {
  it('altera label quando shipNames muda', async () => {
    let board = emptyBoard();
    board = withCell(board,0,0,{ shipId:1 }); // sem state => deriveState => ship
    const { component } = render(BattleshipBoard, { board, mode:'setup', isMyTurn:true, row:0, col:0, roundParam:'1', shipNames:{1:'Alpha'} });
    // label inicial
    const initialCell = screen.getByRole('gridcell', { name: 'Alpha, A1' });
    expect(initialCell).toBeTruthy();
    component.$set({ shipNames: {1:'Porta-aviões'} });
    await tick();
    const updatedCell = screen.getByRole('gridcell', { name: 'Porta-aviões, A1' });
    expect(updatedCell).toBeTruthy();
  });
});
