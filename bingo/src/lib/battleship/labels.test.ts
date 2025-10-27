import { describe, it, expect } from 'vitest';
import { getCellLabelUtil, deriveState, resolveShipName } from '$lib/battleship/labels';

function label(row:number,col:number,cell:any, shipNames?:Record<number,string>) {
  return getCellLabelUtil({ row, col, cell, shipNames });
}

describe('labels util - novo formato estado/nome antes da coordenada', () => {
  it('states snapshot (ordem estado->coord)', () => {
    const outputs = {
      empty: label(0,0,{ state:'empty'}),
      shipFromMap: label(0,0,{ state:'ship', shipId:1 }, {1:'Porta-aviões'}),
      shipFromCellName: label(0,0,{ state:'ship', name:'Cruzador Especial'}),
      shipFallbackId: label(0,0,{ shipId:4 }), // deriveState => ship
      hit: label(0,0,{ state:'hit'}),
      miss: label(0,0,{ state:'miss'}),
      sunk: label(0,0,{ state:'sunk'}),
    };
    expect(outputs).toMatchInlineSnapshot(`
      {
        "empty": "Vazio, A1",
        "hit": "Acerto, A1",
        "miss": "Água, A1",
        "shipFallbackId": "submarino, A1",
        "shipFromCellName": "Cruzador Especial, A1",
        "shipFromMap": "Porta-aviões, A1",
        "sunk": "Afundado, A1",
      }
    `);
  });

  it('nenhum label termina com ponto final isolado', () => {
    const labels = [
      label(0,0,{ state:'empty'}),
      label(0,0,{ state:'ship', shipId:1 }, {1:'Porta-aviões'}),
      label(0,0,{ state:'hit'}),
      label(0,0,{ state:'miss'}),
      label(0,0,{ state:'sunk'}),
    ];
    for (const l of labels) {
      expect(l.endsWith('.')).toBe(false);
    }
  });

  it('coordenada extrema J10 vazia', () => {
    const out = label(9,9,{ state:'empty'});
    expect(out).toBe('Vazio, J10');
  });

  it('deriveState implicita via shipId sem state', () => {
    const out = label(0,0,{ shipId:2 });
    expect(out).toContain('encouraçado');
  });
});
