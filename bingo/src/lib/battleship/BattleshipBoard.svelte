<script>
	// @ts-nocheck
  import { createEventDispatcher } from 'svelte';
  import { getCellLabelUtil, deriveState } from '$lib/battleship/labels';

  // Props
  export let board = []; // 10x10
  export let mode = 'playing'; // 'setup' | 'playing' | 'viewing'
  export let isMyTurn = false;
  export let row = 0;
  export let col = 0;
  export let roundParam = '';
  export let setupShip = null; // id/navio selecionado em modo setup
  export let setupOrientation = 'horizontal';
  export let shipNames = {}; // { shipId: nome }
	export let pendingShots = new Set(); // coordenadas 'r-c' aguardando confirmação
	// Referências para evitar warnings de props não utilizadas (usadas externamente para consistência de API)
	void setupShip; void setupOrientation; void roundParam;

  const dispatch = createEventDispatcher();

  // Eixos para exibição visual
  const colLetters = ['A','B','C','D','E','F','G','H','I','J'];
  const rowNumbers = ['1','2','3','4','5','6','7','8','9','10'];

  function cellId(r,c){ return `battleship-cell-${r}-${c}`; }

  function getCellState(r,c){
    return deriveState(board[r]?.[c]);
  }

  function isCellEngaged(r,c){
    const s = getCellState(r,c);
    return s==='ship'||s==='hit'||s==='sunk';
  }

	function getCellLabel(r,c,names = shipNames){
		return getCellLabelUtil({ row:r, col:c, cell: board[r]?.[c], shipNames: names });
	}

	function cellLabelId(r,c){
		return `battleship-cell-label-${r}-${c}`;
	}

  function getCellClass(r,c){
    const s = getCellState(r,c);
    const classes = ['cell-'+s];
    if(row===r && col===c) classes.push('active');
		if(pendingShots && pendingShots.has(r+'-'+c)) classes.push('pending');
    return classes.join(' ');
  }
	function isPending(r,c){
		return pendingShots && pendingShots.has(r+'-'+c);
	}

  function handleCellClick(r,c){
    dispatch('cellClick',{row:r,col:c});
  }
  function handleCellFocus(r,c){
    dispatch('cellFocus',{row:r,col:c});
  }
  function handleKeydown(e){
    dispatch('keydown',{event:e});
  }
</script>

<div class="battleship-board-wrapper">
	<!-- Labels das colunas (A-J) -->
	<div class="col-labels" aria-hidden="true">
		{#each colLetters as letter}
			<span>{letter}</span>
		{/each}
	</div>
	
	<!-- Labels das linhas (1-10) -->
	<div class="row-labels" aria-hidden="true">
		{#each rowNumbers as number}
			<span>{number}</span>
		{/each}
	</div>
	
	<!-- Tabuleiro principal (loops keyeados para recriar botão quando estado do navio muda) -->
	<p id="board-instructions" class="sr-only">Tabuleiro interativo. Use NVDA+Espaço para entrar em modo foco. Depois use as setas para mover, Enter para atirar ou posicionar. Em modo navegação (browse) apenas o conteúdo visual é lido; para ouvir coordenadas e estado das células ative o modo foco.</p>
	<div
		class="battleship-board"
		class:setup-mode={mode === 'setup'}
		class:playing-mode={mode === 'playing'}
		class:my-turn={isMyTurn}
		role="grid"
		aria-rowcount="10"
		aria-colcount="10"
		aria-label="Tabuleiro {mode === 'setup' ? 'posicionamento' : (mode === 'playing' ? (isMyTurn ? 'ataque, seu turno' : 'defesa, turno do oponente') : 'visualização')}"
		aria-describedby="board-instructions"
	>
		{#each board as boardRow, r}
			<div role="row" aria-rowindex={r + 1} class="grid-row">
				{#each boardRow as cell, c (r + '-' + c) }
					<div
						id={cellId(r, c)}
						data-row={r}
						data-col={c}
						tabindex={row === r && col === c ? 0 : -1}
						role="gridcell"
						aria-roledescription="célula"
						aria-rowindex={r + 1}
						aria-colindex={c + 1}
						class="battleship-cell {getCellClass(r, c)}"
						aria-busy={isPending(r,c) ? 'true' : 'false'}
						on:click={() => handleCellClick(r, c)}
						on:focus={() => handleCellFocus(r, c)}
						on:keydown={handleKeydown}
					>
						<!-- Texto acessível único dentro da célula para modo browse (sem aria-label para evitar perda de leitura). -->
						<span class="sr-only-accessible">{getCellLabel(r, c, shipNames)}{isPending(r,c) ? ' aguardando confirmação' : ''}</span>
						<span class="cell-coord" aria-hidden="true">{colLetters[c]}{rowNumbers[r]}</span>
						<span class="cell-content" aria-hidden="true">
							{#if getCellState(r, c) === 'ship'}
								🚢
							{:else if getCellState(r, c) === 'hit'}
								💥
							{:else if getCellState(r, c) === 'miss'}
								🌊
							{:else if getCellState(r, c) === 'sunk'}
								☠️
							{/if}
						</span>
						{#if isPending(r,c)}
							<span class="pending-indicator" aria-hidden="true"></span>
						{/if}
					</div>
				{/each}
			</div>
		{/each}
	</div>

</div>

<style>
	.battleship-board-wrapper {
		position: relative;
		padding: 2.5rem 2.5rem 0 2.5rem;
		/* espaço para labels externos */
		width: fit-content;
		margin: 0 auto;
	}

	.col-labels, .row-labels {
		font-size: 0.85rem;
		font-weight: 600;
		letter-spacing: 0.05em;
		color: #c7d5e0;
		font-family: system-ui, sans-serif;
	}

	.col-labels {
		position: absolute;
		top: 0;
		left: 2.5rem;
		display: grid;
		grid-template-columns: repeat(10, 1fr);
		width: calc(100% - 2.5rem);
		transform: translateY(-55%);
		pointer-events: none;
	}
	.col-labels span { text-align: center; }

	.row-labels {
		position: absolute;
		left: 0;
		top: 2.5rem;
		display: grid;
		grid-template-rows: repeat(10, 1fr);
		height: calc(100% - 2.5rem);
		transform: translateX(-55%);
		pointer-events: none;
	}
	.row-labels span { display: flex; align-items: center; justify-content: center; }

	/* Responsividade do tabuleiro: tamanho quadrado fluido baseado na largura disponível */
	.battleship-board {
		display: grid;
		grid-template-columns: repeat(10, 1fr);
		grid-template-rows: repeat(10, 1fr);
		/* Usa clamp para limitar tamanho máximo mantendo legibilidade em desktop e adaptação em mobile */
		--board-size: clamp(280px, 90vw, 600px);
		width: var(--board-size);
		height: var(--board-size);
		border: 3px solid #3a5f91;
		border-radius: 1rem;
		background: #101922;
		position: relative;
		box-shadow: 0 0 0 3px rgba(43, 127, 244, 0.35);
		overflow: hidden;
	}
	
	.battleship-cell {
		border: 1px solid #2c4d6b;
		display: flex;
		align-items: center;
		justify-content: center;
		background: linear-gradient(145deg, #16283a, #1e344b);
		color: #f5f7fa;
		cursor: pointer;
		user-select: none;
		transition: all 0.2s ease;
		font-size: 1.2rem;
		position: relative;
		padding: 0;
		min-height: 45px;
	}
	
	.battleship-cell:hover:not(:disabled) {
		background: linear-gradient(145deg, #254362, #2d4b6a);
		border-color: #4a7ba7;
	}
	
	.battleship-cell:disabled {
		cursor: not-allowed;
		opacity: 0.6;
	}
	
	.battleship-cell.active {
		box-shadow: inset 0 0 0 3px #ffc107;
		background: linear-gradient(145deg, #2d5173, #354a6b);
		z-index: 2;
	}
	
	/* Estados das células */
	.battleship-cell.cell-empty {
		background: linear-gradient(145deg, #16283a, #1e344b);
	}
	
	.battleship-cell.cell-ship {
		background: linear-gradient(145deg, #1a4a1a, #0f2f0f);
		border-color: #4a7c59;
	}
	
	.battleship-cell.cell-hit {
		background: linear-gradient(145deg, #4a1a1a, #2f0f0f);
		border-color: #cc3333;
		animation: hit-flash 0.5s ease-in-out;
	}
	
	.battleship-cell.cell-miss {
		background: linear-gradient(145deg, #1a1a4a, #0f0f2f);
		border-color: #4169e1;
	}
	
	.battleship-cell.cell-sunk {
		background: linear-gradient(145deg, #2f1a1a, #1a0f0f);
		border-color: #8b0000;
		animation: sink-pulse 1s ease-in-out;
	}

	/* Estado pendente */
	.battleship-cell.pending {
		position: relative;
		outline: 1px dashed #ffc107;
	}
	.pending-indicator {
		position: absolute;
		width: 60%;
		height: 60%;
		border-radius: 50%;
		border: 3px solid rgba(255,193,7,0.4);
		border-top-color: #ffc107;
		animation: spin 0.8s linear infinite;
		pointer-events: none;
	}
	@keyframes spin {
		0% { transform: rotate(0deg); }
		100% { transform: rotate(360deg); }
	}
	
	.cell-content {
		font-size: 1.5rem;
		filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.5));
	}

	/* Coordenada visual pequena para apoiar navegação em modo browse de leitores */
	.cell-coord {
		position: absolute;
		top: 2px;
		left: 4px;
		font-size: 0.55rem;
		font-weight: 600;
		letter-spacing: .5px;
		color: rgba(255,255,255,0.55);
		text-shadow: 0 0 2px rgba(0,0,0,0.8);
		pointer-events: none;
		user-select: none;
	}
	
	/* Modos do tabuleiro */
	.battleship-board.setup-mode {
		border-color: #28a745;
		box-shadow: 0 0 0 3px rgba(40, 167, 69, 0.35);
	}
	
	.battleship-board.playing-mode.my-turn {
		border-color: #ffc107;
		box-shadow: 0 0 0 3px rgba(255, 193, 7, 0.35);
	}
	
	.battleship-board.playing-mode:not(.my-turn) {
		border-color: #6c757d;
		box-shadow: 0 0 0 3px rgba(108, 117, 125, 0.35);
	}
	
	/* Animações */
	@keyframes hit-flash {
		0%, 100% { transform: scale(1); }
		50% { transform: scale(1.1); box-shadow: 0 0 20px rgba(255, 0, 0, 0.8); }
	}
	
	@keyframes sink-pulse {
		0%, 100% { opacity: 1; }
		50% { opacity: 0.5; }
	}
	
	/* Responsividade */
	@media (max-width: 768px) {
		.battleship-board-wrapper { padding: 2rem 1.5rem 0 1.5rem; }
		.battleship-board { --board-size: clamp(260px, 92vw, 480px); }
		.battleship-cell { min-height: 32px; }
		.cell-content { font-size: 1.1rem; }
		.col-labels, .row-labels { font-size: 0.8rem; }
		.col-labels { top: -1.3rem; }
		.row-labels { left: -1.3rem; }
	}
	
	@media (max-width: 480px) {
		.battleship-board-wrapper { padding: 1.6rem 1rem 0 1rem; }
		.battleship-board { --board-size: clamp(240px, 94vw, 420px); }
		.battleship-cell { min-height: 24px; }
		.cell-content { font-size: 0.9rem; }
		.col-labels, .row-labels { font-size: 0.7rem; }
		.cell-coord { font-size: 0.5rem; }
	}

	@media (max-width: 360px) {
		.battleship-board { --board-size: clamp(220px, 96vw, 360px); }
		.battleship-cell { min-height: 22px; }
		.cell-content { font-size: 0.75rem; }
	}

	/* Reutiliza padrão de classe escondida para labels explícitos das células */
	.sr-only {
		position: absolute;
		width: 1px;
		height: 1px;
		padding: 0;
		margin: -1px;
		overflow: hidden;
		clip: rect(0 0 0 0);
		white-space: nowrap;
		border: 0;
	}
	</style>