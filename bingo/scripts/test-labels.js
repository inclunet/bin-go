#!/usr/bin/env node
const assert = require('assert');

// Replica mínima da lógica atual de label (mantida simples e isolada)
const colLetters = ['A','B','C','D','E','F','G','H','I','J'];
const rowNumbers = ['1','2','3','4','5','6','7','8','9','10'];

function label({r,c,cell}) {
  const coord = `${colLetters[c]}${rowNumbers[r]}`;
  const state = cell.state || 'empty';
  switch (state) {
    case 'ship': return `${coord}. ${cell.name || 'Navio'}`;
    case 'hit': return `${coord}. Acerto`;
    case 'miss': return `${coord}. Água`;
    case 'sunk': return `${coord}. Afundado`;
    default: return `${coord}. Vazio`;
  }
}

function run() {
  const cases = [
    {r:0,c:0,cell:{state:'ship', name:'Porta-aviões'}, expected:'A1. Porta-aviões'},
    {r:1,c:1,cell:{state:'hit'}, expected:'B2. Acerto'},
    {r:2,c:2,cell:{state:'miss'}, expected:'C3. Água'},
    {r:3,c:3,cell:{state:'sunk'}, expected:'D4. Afundado'},
    {r:4,c:4,cell:{state:'empty'}, expected:'E5. Vazio'}
  ];
  for (const t of cases) {
    const got = label(t);
    assert.strictEqual(got, t.expected, `Esperado '${t.expected}' mas obteve '${got}'`);
  }
  console.log('OK: labels básicos gerados conforme esperado.');
}

run();
