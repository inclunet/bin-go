import { describe, it, expect } from 'vitest';

// Esses testes são de alto nível e não invocam o backend real.
// Simulam a lógica de criação de rodadas de Battleship replicando o padrão de TicTac.

interface Round {
  Round: number;
  Winner: string;
  Next?: number;
  ScoreA: number;
  ScoreB: number;
  ScoreDraw: number;
}

function newRoundHandlerSim(rounds: Round[], current: number): { round?: Round; error?: string } {
  // Caso base
  if (rounds.length === 0) {
    if (current !== 1) return { error: 'invalid round number' };
    const first: Round = { Round: 1, Winner: '', ScoreA: 0, ScoreB: 0, ScoreDraw: 0 };
    rounds.push(first);
    return { round: first };
  }
  const nextNumber = rounds.length + 1;
  if (current === nextNumber) {
    // Nova cadeia independente
    const nr: Round = { Round: nextNumber, Winner: '', ScoreA: 0, ScoreB: 0, ScoreDraw: 0 };
    rounds.push(nr);
    return { round: nr };
  }
  const sourceIdx = current - 1;
  if (sourceIdx < 0 || sourceIdx >= rounds.length) return { error: 'source round not found' };
  const source = rounds[sourceIdx];
  if (source.Winner === '') return { error: 'source round not finished' };
  if (source.Next && source.Next !== 0) return { error: 'round already continued' };
  const cont: Round = {
    Round: nextNumber,
    Winner: '',
    ScoreA: source.ScoreA,
    ScoreB: source.ScoreB,
    ScoreDraw: source.ScoreDraw,
  };
  rounds.push(cont);
  source.Next = cont.Round;
  return { round: cont };
}

describe('Battleship round creation (alinhado com TicTac)', () => {
  it('cria primeira rodada somente com /1/new', () => {
    const rounds: Round[] = [];
    const fail = newRoundHandlerSim(rounds, 2);
    expect(fail.error).toBe('invalid round number');
    const ok = newRoundHandlerSim(rounds, 1);
    expect(ok.round?.Round).toBe(1);
    expect(rounds.length).toBe(1);
  });

  it('cria nova sequência independente quando current == len+1', () => {
    const rounds: Round[] = [{ Round: 1, Winner: '', ScoreA: 0, ScoreB: 0, ScoreDraw: 0 }];
    const res = newRoundHandlerSim(rounds, 2);
    expect(res.round?.Round).toBe(2);
    expect(rounds.length).toBe(2);
    expect(rounds[1].ScoreA).toBe(0);
  });

  it('retorna erro se tentar continuar rodada não finalizada', () => {
    const rounds: Round[] = [
      { Round: 1, Winner: '', ScoreA: 3, ScoreB: 2, ScoreDraw: 1 },
    ];
    const res = newRoundHandlerSim(rounds, 1);
    expect(res.error).toBe('source round not finished');
  });

  it('continua rodada finalizada herda placar e seta Next', () => {
    const rounds: Round[] = [
      { Round: 1, Winner: 'A', ScoreA: 5, ScoreB: 4, ScoreDraw: 2 },
    ];
    const res = newRoundHandlerSim(rounds, 1);
    expect(res.round?.Round).toBe(2);
    expect(rounds[0].Next).toBe(2);
    expect(res.round?.ScoreA).toBe(5);
    expect(res.round?.ScoreB).toBe(4);
    expect(res.round?.ScoreDraw).toBe(2);
  });

  it('impede segunda continuidade da mesma origem', () => {
    const rounds: Round[] = [
      { Round: 1, Winner: 'B', ScoreA: 1, ScoreB: 2, ScoreDraw: 0 },
    ];
    const first = newRoundHandlerSim(rounds, 1);
    expect(first.round?.Round).toBe(2);
    const second = newRoundHandlerSim(rounds, 1);
    expect(second.error).toBe('round already continued');
  });
});
