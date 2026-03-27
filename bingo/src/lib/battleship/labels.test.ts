import { describe, it, expect } from 'vitest';
import { getCellLabelUtil } from '$lib/battleship/labels';

function label(row:number,col:number,cell:any, shipNames?:Record<number,string>) {
	return getCellLabelUtil({ row, col, cell, shipNames });
}

describe('labels util básico', () => {
	it('empty cell', () => {
		expect(label(0,0,{ state:'empty'})).toBe('Vazio, A1');
	});
	it('ship id resolves name map', () => {
		expect(label(0,0,{ state:'ship', shipId:1 }, {1:'Porta-aviões'})).toBe('Porta-aviões, A1');
	});
	it('hit', () => {
		expect(label(0,0,{ state:'hit'})).toBe('Acerto, A1');
	});
	it('miss', () => {
		expect(label(0,0,{ state:'miss'})).toBe('Água, A1');
	});
	it('sunk', () => {
		expect(label(0,0,{ state:'sunk'})).toBe('Afundado, A1');
	});
});
