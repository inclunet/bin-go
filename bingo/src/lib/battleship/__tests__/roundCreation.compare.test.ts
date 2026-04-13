import { describe, it, expect } from 'vitest';

// Teste comparativo mínimo de contrato entre respostas de criação/continuação de rodada
// Não chama backend real: usa objetos mock simulando campos que frontend consome.
// Objetivo: garantir que campos essenciais (round, winner, next, placares) existam em ambos.

describe('Comparação de shape TicTac vs Battleship (contrato de round)', () => {
  function mockTicTacRound(partial: Partial<any> = {}) {
    return {
      round: 3,
      winner: '',
      next: 0,
      scoreX: 2,
      scoreO: 1,
      scoreDraw: 0,
      ...partial
    };
  }
  function mockBattleshipRound(partial: Partial<any> = {}) {
    return {
      round: 3,
      winner: '',
      next: 0,
      scoreA: 2,
      scoreB: 1,
      scoreDraw: 0,
      phase: 'playing',
      currentPlayer: 'a',
      ...partial
    };
  }

  it('Campos de placar estão presentes e numéricos', () => {
    const tic = mockTicTacRound();
    const bat = mockBattleshipRound();
    console.log('DEBUG placares', { tic, bat });
    expect(typeof tic.scoreX).toBe('number');
    expect(typeof tic.scoreO).toBe('number');
    expect(typeof tic.scoreDraw).toBe('number');
    expect(typeof bat.scoreA).toBe('number');
    expect(typeof bat.scoreB).toBe('number');
    expect(typeof bat.scoreDraw).toBe('number');
  });

  it('Campos comuns round/winner/next existem', () => {
    const tic = mockTicTacRound();
    const bat = mockBattleshipRound();
    for (const obj of [tic, bat]) {
      expect(typeof obj.round).toBe('number');
      expect(typeof obj.winner).toBe('string');
      expect(typeof obj.next).toBe('number');
    }
  });

  it('Diferenças específicas não quebram contrato básico', () => {
    const bat = mockBattleshipRound();
    expect(bat.phase).toBeDefined();
    expect(['setup','playing','finished','PhaseSetup','PhasePlaying','PhaseFinished'].some(k => String(bat.phase).toLowerCase().includes(k.replace('Phase','').toLowerCase()))).toBe(true);
    expect(['a','b',''].includes(String(bat.currentPlayer))).toBe(true);
  });
});
