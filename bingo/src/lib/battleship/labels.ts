// Util para geração de labels acessíveis das células do tabuleiro
// Fonte única para manutenção e testes.

export interface CellData {
  state?: string; // 'empty' | 'ship' | 'hit' | 'miss' | 'sunk'
  shipId?: number;
  name?: string; // nome do navio
  shipName?: string; // variante legacy
}

export interface LabelContext {
  row: number; // 0-based
  col: number; // 0-based
  cell: CellData | undefined;
  shipNames?: Record<number, string>;
}

const colLetters = ['A','B','C','D','E','F','G','H','I','J'];
const rowNumbers = ['1','2','3','4','5','6','7','8','9','10'];

function baseShipName(id?: number) {
  if (id == null) return 'Navio';
  const base: Record<number,string> = {
    1: 'porta aviões',
    2: 'encouraçado',
    3: 'cruzador',
    4: 'submarino',
    5: 'destroyer'
  };
  return base[id] || 'navio';
}

export function resolveShipName(cell: CellData | undefined, shipNames?: Record<number,string>) {
  if (!cell) return null;
  if (cell.shipId != null && shipNames && shipNames[cell.shipId]) return shipNames[cell.shipId];
  if (cell.name) return cell.name;
  if (cell.shipName) return cell.shipName;
  if (cell.shipId != null) return baseShipName(cell.shipId);
  return null;
}

export function deriveState(cell: CellData | undefined): string {
  if (!cell) return 'empty';
  let st = cell.state || 'empty';
  if (st === 'empty' && (cell.shipId != null || cell.name || cell.shipName)) {
    st = 'ship';
  }
  return st;
}

export function getCellLabelUtil(ctx: LabelContext): string {
  const { row, col, cell, shipNames } = ctx;
  const coord = `${colLetters[col]}${rowNumbers[row]}`;
  const st = deriveState(cell);
  switch (st) {
    case 'ship': {
      const name = resolveShipName(cell, shipNames);
      // Evitar anunciar palavra genérica "Navio"; se não houver nome resolvível usar só a coordenada
      // Ordem desejada: estado/nome primeiro, coordenada por último
      return name ? `${name}, ${coord}` : `${coord}`;
    }
    case 'hit': return `Acerto, ${coord}`;
    case 'miss': return `Água, ${coord}`;
    case 'sunk': return `Afundado, ${coord}`;
    default: return `Vazio, ${coord}`;
  }
}
