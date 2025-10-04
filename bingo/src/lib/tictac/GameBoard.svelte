<script>
	// @ts-nocheck
	import { createEventDispatcher } from 'svelte';
	
	export let board = [['','',''],['','',''],['','','']];
	export let row = 0;
	export let col = 0;
	export let roundParam = '';

	export let localPlayer = '';

	const dispatch = createEventDispatcher();
	const colNames = ['a','b','c'];

	function cellId(r, c) {
		return `cell-${r}-${c}`;
	}

	function handleCellClick(r, c) {
		dispatch('cellClick', { row: r, col: c });
	}

	function handleCellFocus(r, c) {
		dispatch('cellFocus', { row: r, col: c });
	}

	function handleKeydown(e) {
		dispatch('keydown', { event: e });
	}
</script>

<div class="board-wrapper">
	<div class="col-labels" aria-hidden="true">
		<span>a</span><span>b</span><span>c</span>
	</div>
	<div class="row-labels" aria-hidden="true">
		<span>1</span><span>2</span><span>3</span>
	</div>
	<div class="tic-board" aria-label="Tabuleiro do jogo da velha, rodada {roundParam}. Jogador local {localPlayer}.">
		{#each board as rowVals, r}
			{#each rowVals as cellVal, c}
				<button 
					id={cellId(r,c)} 
					data-r={r} 
					data-c={c} 
					tabindex={row===r && col===c ? 0 : -1}
					aria-pressed={board[r][c] !== ''}
					aria-label={board[r][c] ? `${board[r][c]} ${colNames[c]}${r+1}` : `vazio ${colNames[c]}${r+1}`}
					class="tic-cell {board[r][c] ? `filled-${board[r][c]}`: ''} {row===r && col===c ? 'active' : ''}"
					on:click={() => handleCellClick(r,c)}
					on:focus={() => handleCellFocus(r,c)}
					on:keydown={handleKeydown}
				>
					{cellVal}
				</button>
			{/each}
		{/each}
	</div>
</div>

<style>
	.board-wrapper { 
		position: relative; 
		width: 100%; 
		max-width: 38rem; 
	}
	
	.tic-board { 
		display: grid; 
		grid-template-columns: repeat(3,1fr); 
		width: 100%; 
		aspect-ratio: 1/1; 
		border: 3px solid #3a5f91; 
		border-radius: 0.9rem; 
		background: #101922; 
		position: relative; 
		box-shadow: 0 0 0 3px rgba(43,127,244,0.35); 
	}
	
	.tic-cell { 
		border: 2px solid #3a5f91; 
		display: flex; 
		align-items: center; 
		justify-content: center; 
		font-size: clamp(2.2rem, 8vw, 4.2rem); 
		font-weight: 700; 
		cursor: pointer; 
		user-select: none; 
		background: linear-gradient(#16283a,#1e344b); 
		color: #f5f7fa; 
		transition: background 0.15s, transform 0.08s; 
		box-shadow: inset 0 0 0 1px rgba(43,127,244,0.2); 
	}
	
	.tic-cell:hover { 
		background: #254362; 
	}
	
	.tic-cell:active { 
		transform: scale(0.96); 
	}
	
	.tic-cell.filled-X { 
		color: #2aa9ff; 
	}
	
	.tic-cell.filled-O { 
		color: #ff8b54; 
	}
	
	.tic-cell.active { 
		box-shadow: inset 0 0 0 3px #ffc107, 0 0 0 2px #000; 
		background: #2d5173; 
	}
	
	.col-labels { 
		display: grid; 
		grid-template-columns: repeat(3,1fr); 
		position: absolute; 
		top: -1.6rem; 
		left: 0; 
		width: 100%; 
		font-size: 0.9rem; 
		text-transform: lowercase; 
		color: #c6d6e5; 
		letter-spacing: 0.05em; 
	}
	
	.col-labels span { 
		text-align: center; 
	}
	
	.row-labels { 
		position: absolute; 
		top: 0; 
		left: -1.4rem; 
		height: 100%; 
		display: grid; 
		grid-template-rows: repeat(3,1fr); 
		font-size: 0.9rem; 
		color: #c6d6e5; 
	}
	
	.row-labels span { 
		display: flex; 
		align-items: center; 
		justify-content: center; 
	}
	
	@media (max-width: 480px) { 
		.col-labels { 
			top: -1.3rem; 
		} 
		.row-labels { 
			left: -1.1rem; 
		}
	}
</style>