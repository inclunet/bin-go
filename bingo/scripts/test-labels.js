#!/usr/bin/env node
import assert from 'node:assert';

const colLetters = ['A','B','C','D','E','F','G','H','I','J'];
const rowNumbers = ['1','2','3','4','5','6','7','8','9','10'];

function getCellLabelUtil({ row, col, cell, shipNames = {} }) {
	const coord = `${colLetters[col]}${rowNumbers[row]}`;
	const st = cell?.state || 'empty';
	switch (st) {
		case 'ship': {
			const name = cell.name || (cell.shipId != null ? shipNames[cell.shipId] : null);
			return name ? `${name}, ${coord}` : `${coord}`;
		}
		case 'hit': return `Acerto, ${coord}`;
		case 'miss': return `Água, ${coord}`;
		case 'sunk': return `Afundado, ${coord}`;
		default: return `Vazio, ${coord}`;
	}
}

const cases = [
	{ row: 0, col: 0, cell: { state: 'ship', name: 'Porta-aviões' }, expected: 'Porta-aviões, A1' },
	{ row: 1, col: 1, cell: { state: 'hit' }, expected: 'Acerto, B2' },
	{ row: 2, col: 2, cell: { state: 'miss' }, expected: 'Água, C3' },
	{ row: 3, col: 3, cell: { state: 'sunk' }, expected: 'Afundado, D4' },
	{ row: 4, col: 4, cell: { state: 'empty' }, expected: 'Vazio, E5' }
];

for (const t of cases) {
	const got = getCellLabelUtil(t);
	assert.strictEqual(got, t.expected, `Esperado '${t.expected}' mas obteve '${got}'`);
}

console.log('OK: labels básicos gerados conforme esperado.');
