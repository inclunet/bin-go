<script>
	// @ts-nocheck
  import { createEventDispatcher } from 'svelte';
  import { getCellLabelUtil, deriveState } from '$lib/battleship/labels';

  // Props
  export let board = []; // 10x10
  export let mode = 'playing'; // 'setup' | 'playing' | 'viewing'
  export let viewMode = 'my'; // 'my' (defesa) | 'enemy' (ataque)
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

  // Eixos para exibição visual (fora do grid acessível, como no jogo da velha)
  const colLetters = ['A','B','C','D','E','F','G','H','I','J'];
  const rowNumbers = ['1','2','3','4','5','6','7','8','9','10'];

  function cellId(r,c){ return `battleship-cell-${r}-${c}`; }

  function getCellState(r,c){
    return deriveState(board[r]?.[c]);
  }

	function getCellLabel(r,c,names = shipNames){
		return getCellLabelUtil({ row:r, col:c, cell: board[r]?.[c], shipNames: names });
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
	function getBoardAriaLabel() {
		if (mode === 'setup') return 'Tabuleiro posicionamento';
		if (mode !== 'playing') return 'Tabuleiro visualização';
		if (viewMode === 'my') {
			return isMyTurn ? 'Tabuleiro defesa, seu turno' : 'Tabuleiro defesa, turno do oponente';
		}
		return isMyTurn ? 'Tabuleiro ataque, seu turno' : 'Tabuleiro ataque, turno do oponente';
	}

  function handleKeydown(e){
    dispatch('keydown',{event:e});
  }
</script>

<div class="board-shell">
	<p id="board-instructions" class="sr-only">Tabuleiro interativo. Use NVDA+Espaço para foco. Setas movem; Enter confirma ação.</p>

	<div class="battleship-board-wrapper">
	<!-- Rótulos visuais fora do role=grid (layout em CSS grid, sem position absolute) -->
	<div class="label-corner" aria-hidden="true"></div>
	<div class="col-labels" aria-hidden="true">
		{#each colLetters as letter}
			<span>{letter}</span>
		{/each}
	</div>
	<div class="row-labels" aria-hidden="true">
		{#each rowNumbers as num}
			<span>{num}</span>
		{/each}
	</div>

	<div class="board-aspect">
	<div class="battleship-board"
		class:setup-mode={mode === 'setup'}
		class:playing-mode={mode === 'playing'}
		class:my-turn={isMyTurn}
		role="grid"
		aria-rowcount="10"
		aria-colcount="10"
		aria-label={getBoardAriaLabel()}
		aria-describedby="board-instructions"
	>
		{#each board as boardRow, r}
			{#each boardRow as cell, c (r + '-' + c)}
				<div
					id={cellId(r,c)}
					data-row={r}
					data-col={c}
					tabindex={row===r && col===c ? 0 : -1}
					role="gridcell"
					aria-roledescription="célula"
					aria-rowindex={r+1}
					aria-colindex={c+1}
					aria-label="{getCellLabel(r,c, shipNames)}{isPending(r,c) ? ' aguardando confirmação' : ''}"
					aria-busy={isPending(r,c) ? 'true' : 'false'}
					class="battleship-cell {getCellClass(r,c)}"
					on:click={() => handleCellClick(r,c)}
					on:focus={() => handleCellFocus(r,c)}
					on:keydown={handleKeydown}
				>
					<span class="cell-content" aria-hidden="true">{#if getCellState(r,c) === 'ship'}🚢{:else if getCellState(r,c) === 'hit'}💥{:else if getCellState(r,c) === 'miss'}🌊{:else if getCellState(r,c) === 'sunk'}☠️{/if}</span>
					{#if isPending(r,c)}<span class="pending-indicator" aria-hidden="true"></span>{/if}
				</div>
			{/each}
		{/each}
	</div>
	</div>
	</div>
</div>

<style>
	.board-shell {
		box-sizing: border-box;
		width: 100%;
		max-width: 100%;
		margin-inline: auto;
		position: relative;
	}

	.battleship-board-wrapper {
		--label-size: 1.6rem;
		--board-max: 640px;
		box-sizing: border-box;
		display: grid;
		grid-template-columns: var(--label-size) minmax(0, 1fr);
		grid-template-rows: var(--label-size) auto;
		align-items: stretch;
		width: 100%;
		max-width: min(100%, calc(var(--board-max) + var(--label-size)));
		margin-inline: auto;
		padding: 0;
	}

	.label-corner {
		grid-column: 1;
		grid-row: 1;
	}

	.col-labels {
		grid-column: 2;
		grid-row: 1;
		display: grid;
		grid-template-columns: repeat(10, minmax(0, 1fr));
		align-items: center;
		font-size: 0.85rem;
		font-weight: 600;
		letter-spacing: 0.05em;
		color: #c7d5e0;
		pointer-events: none;
	}
	.col-labels span {
		text-align: center;
	}

	.row-labels {
		grid-column: 1;
		grid-row: 2;
		display: grid;
		grid-template-rows: repeat(10, minmax(0, 1fr));
		justify-items: center;
		align-self: stretch;
		min-height: 0;
		font-size: 0.85rem;
		font-weight: 600;
		letter-spacing: 0.05em;
		color: #c7d5e0;
		pointer-events: none;
	}
	.row-labels span {
		display: flex;
		align-items: center;
		justify-content: center;
	}

	/* Quadrado proporcional à largura (Safari/iOS não encolhe com aspect-ratio em grid auto) */
	.board-aspect {
		grid-column: 2;
		grid-row: 2;
		width: 100%;
		min-width: 0;
		position: relative;
		height: 0;
		padding-top: 100%;
	}

	.battleship-board {
		position: absolute;
		inset: 0;
		display: grid;
		grid-template-columns: repeat(10, minmax(0, 1fr));
		grid-template-rows: repeat(10, minmax(0, 1fr));
		box-sizing: border-box;
		width: 100%;
		height: 100%;
		min-width: 0;
		border: 3px solid #3a5f91;
		border-radius: 1rem;
		background: #101922;
		box-shadow: 0 0 0 3px rgba(43, 127, 244, 0.35);
		overflow: hidden;
		margin: 0;
	}

	/* Portrait mobile: largura do pai (evita overflow de 100vw no Safari/iOS) */
	@media (max-width: 768px) and (orientation: portrait) {
		.board-shell {
			width: 100%;
			max-width: 100%;
			margin-inline: 0;
			padding-left: max(0.15rem, env(safe-area-inset-left, 0px));
			padding-right: max(0.15rem, env(safe-area-inset-right, 0px));
		}
		.battleship-board-wrapper {
			--label-size: 0.95rem;
			--board-max: none;
			width: 100%;
			max-width: none;
			margin-inline: 0;
		}
		.col-labels, .row-labels { font-size: 0.68rem; }
	}

	@media (max-width: 768px) and (orientation: landscape) {
		.battleship-board-wrapper {
			--label-size: 1.2rem;
			width: 100%;
			max-width: 100%;
		}
		.col-labels, .row-labels { font-size: 0.75rem; }
	}

	@media (max-width: 480px) and (orientation: portrait) {
		.battleship-board-wrapper {
			--label-size: 0.85rem;
		}
		.col-labels, .row-labels { font-size: 0.62rem; }
	}

	@media (max-width: 360px) and (orientation: portrait) {
		.battleship-board-wrapper {
			--label-size: 0.78rem;
		}
		.cell-content { font-size: clamp(0.7rem, 1.4vw + 0.35rem, 0.95rem); }
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
		transition: background .2s ease, box-shadow .2s ease;
		font-size: clamp(0.7rem, 1.2vw + 0.35rem, 1.2rem);
		position: relative;
		padding: 0;
		min-height: 0;
		min-width: 0;
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
		font-size: clamp(0.85rem, 1.6vw + 0.4rem, 1.5rem);
		filter: drop-shadow(0 2px 4px rgba(0, 0, 0, 0.5));
		pointer-events: none;
	}
	
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
	
	@keyframes hit-flash {
		0%, 100% { transform: scale(1); }
		50% { transform: scale(1.1); box-shadow: 0 0 20px rgba(255, 0, 0, 0.8); }
	}
	
	@keyframes sink-pulse {
		0%, 100% { opacity: 1; }
		50% { opacity: 0.5; }
	}

	/* Paisagem: limita pelo menor eixo, mas mantém o bloco centralizado */
	@media (max-height: 620px) and (orientation: landscape) {
		.battleship-board-wrapper {
			max-width: min(100%, calc(70vmin + var(--label-size)));
			margin-inline: auto;
		}
	}

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
