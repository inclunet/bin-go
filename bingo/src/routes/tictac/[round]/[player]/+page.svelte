<script>
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';

	const p = get(page);
	let roundParam = p.params.round; // poderá ser atualizado via redirect backend
	const playerParam = (p.params.player || 'x').toLowerCase();

	let board = [['','',''],['','',''],['','','']];
	let row = 0; let col = 0;
	const localPlayer = playerParam === 'o' ? 'O' : 'X';
	let currentTurn = 'X';
	let statusMsg = 'Carregando partida...';
	let winner = '';
	let showHelp = false;
	const colNames = ['a','b','c'];
		const GAME_NAME = 'Jogo da velha inclusivo';
		let ws;
		let redirecting = false;

		/** @param {number} r @param {number} c */
		const cellId = (r,c) => `cell-${r}-${c}`;

	function announcePosition() { /* mantido para futura expansão */ }

	/** @param {string} v */
	function normalizeTurn(v){ return (v||'').toString().trim().toUpperCase(); }
	function isMyTurn(){ return normalizeTurn(currentTurn) === localPlayer; }

	async function loadRound() {
		try {
			const res = await fetch(`/api/tictac/${roundParam}`, { cache: 'no-store' });
			if (res.status === 404) { statusMsg = 'Rodada inexistente.'; return; }
			if (!res.ok) return;
			const data = await res.json();
			if (data?.board) {
				board = data.board;
				currentTurn = normalizeTurn(data.turn) || currentTurn;
				winner = data.winner || '';
				statusMsg = winner ? (winner === 'Empate' ? 'Empate!' : `Jogador ${winner} venceu!`) : `Vez de ${currentTurn}`;
				// debug removido
				handleMaybeRedirect(data);
			}
		} catch(e) { console.error('Erro ao carregar rodada', e); }
	}

	async function ensureRound() {
		try {
			const res = await fetch(`/api/tictac/${roundParam}`);
			if (res.status === 404) {
				const create = await fetch(`/api/tictac/${roundParam}/new`);
				if (create.ok) {
					const data = await create.json();
					if (data?.round && Number(data.round) !== Number(roundParam)) {
						redirecting = true;
						window.location.href = `/tictac/${data.round}/${playerParam}`;
						return;
					}
					await loadRound();
				}
			} else if (res.ok) {
				await loadRound();
			}
		} catch (e) { console.error('Erro garantindo rodada', e); }
	}

	/** @param {number} r @param {number} c */
	async function playMove(r,c) {
		if (winner) return;
		if (board[r][c] !== '') { statusMsg = 'Casa já ocupada.'; return; }
		if (!isMyTurn()) {
			statusMsg = `Ainda não é sua vez. Turno de ${currentTurn}.`;
			return;
		}
		try {
			const res = await fetch(`/api/tictac/${roundParam}/${playerParam}/${r+1}/${c+1}`, { cache: 'no-store' });
			if (!res.ok) return;
			const data = await res.json();
			if (data?.board) {
				board = data.board;
				currentTurn = normalizeTurn(data.turn) || currentTurn;
				winner = data.winner || '';
				if (winner) {
					statusMsg = winner === 'Empate' ? 'Empate!' : `Jogador ${winner} venceu!`;
				} else {
					statusMsg = `Jogada registrada. Vez de ${currentTurn}`;
				}
				// debug removido
				handleMaybeRedirect(data);
			}
		} catch(e){ console.error('Erro jogada', e); }
	}

	/** @param {number} r @param {number} c */
	function disabledCell(r,c) {
		if (winner) return true;
		if (board[r][c] !== '') return true;
		return false; // não desabilita mais fora do turno
	}

		/** @param {KeyboardEvent} e */
		function handleKey(e) {
		if (winner) return;
		const key = e.key;
		if (!['ArrowUp','ArrowDown','ArrowLeft','ArrowRight','Home','End','Enter',' '].includes(key)) return;
		const target = /** @type {HTMLElement} */(e.currentTarget);
		const r = parseInt(target.dataset.r || '0');
		const c = parseInt(target.dataset.c || '0');
		if (key === 'Enter' || key === ' ') {
				row = r; col = c; playMove(r,c); e.preventDefault(); return; }
		let nr = r, nc = c;
		switch (key) {
			case 'ArrowUp': nr = Math.max(0, r-1); break;
			case 'ArrowDown': nr = Math.min(2, r+1); break;
			case 'ArrowLeft': nc = Math.max(0, c-1); break;
			case 'ArrowRight': nc = Math.min(2, c+1); break;
			case 'Home': nr = 0; nc = 0; break;
			case 'End': nr = 2; nc = 2; break;
		}
		if (nr !== r || nc !== c) { row = nr; col = nc; focusActiveCell(); announcePosition(); }
		e.preventDefault();
	}

	function reset() { window.location.reload(); }

	let boardRef = null;
	function focusActiveCell() {
		const el = document.getElementById(cellId(row,col));
		el?.focus();
	}

	onMount(async () => {
		await ensureRound();
		focusActiveCell();
		try {
			ws = new WebSocket(`${location.protocol === 'https:' ? 'wss' : 'ws'}://${location.host}/ws/tictac/${roundParam}/${playerParam}`);
			ws.onmessage = (ev) => {
				try {
					const data = JSON.parse(ev.data);
					if (data.board) {
						board = data.board;
						currentTurn = normalizeTurn(data.turn) || currentTurn;
						winner = data.winner || '';
						statusMsg = winner ? (winner === 'Empate' ? 'Empate!' : `Jogador ${winner} venceu!`) : `Vez de ${currentTurn}`;
						// debug removido
						handleMaybeRedirect(data);
					}
				} catch(e){ console.error('WS parse', e); }
			};
		} catch(e){ console.error('WS error', e); }
	});

// debug removido

	/** @param {any} data */
	function handleMaybeRedirect(data){
		const nxt = Number(data?.next);
		if(!redirecting && nxt && nxt > Number(roundParam)) {
			redirecting = true;
			window.location.href = `/tictac/${nxt}/${playerParam}`;
		}
	}

	async function newRound(){
		if(redirecting) return;
		try {
			const res = await fetch(`/api/tictac/${roundParam}/new`, { method:'GET', cache:'no-store' });
			if(res.ok){
				const data = await res.json();
				handleMaybeRedirect(data);
			}
		} catch(e){ console.error('Erro nova rodada', e); }
	}
</script>

<PageTitle title="Rodada" game="Jogo da Velha" />

<div class="container tic-container py-4 d-flex flex-column align-items-center">
	<p class="turn mb-2">Vez: <strong>{currentTurn}</strong> {currentTurn === localPlayer ? '(Você)' : ''}</p>
	<div class="board-wrapper my-2">
		<button class="btn btn-sm btn-outline-light" style="margin:0 0 .5rem auto; display:block;" on:click={() => showHelp = !showHelp} aria-expanded={showHelp} aria-controls="help-panel">{showHelp ? 'Ocultar ajuda' : 'Ajuda'}</button>
		<div class="col-labels" aria-hidden="true"><span>a</span><span>b</span><span>c</span></div>
		<div class="row-labels" aria-hidden="true"><span>1</span><span>2</span><span>3</span></div>
		<div class="tic-board" aria-label="Tabuleiro do jogo da velha, rodada {roundParam}. Jogador local {localPlayer}." aria-describedby="grid-desc" bind:this={boardRef}>
			{#each board as rowVals, r}
				{#each rowVals as cellVal, c}
					<button id={cellId(r,c)} data-r={r} data-c={c} tabindex={row===r && col===c ? 0 : -1}
						aria-pressed={board[r][c] !== ''}
						aria-label={board[r][c] ? `${board[r][c]} ${colNames[c]}${r+1}` : `vazio ${colNames[c]}${r+1}`}
						class="tic-cell {board[r][c] ? `filled-${board[r][c]}`: ''} {row===r && col===c ? 'active' : ''}"
						on:click={() => { if (!disabledCell(r,c)) { row = r; col = c; playMove(r,c); focusActiveCell(); } }}
						disabled={disabledCell(r,c)}
						on:focus={() => { row = r; col = c; }}
						on:keydown={handleKey}>{cellVal}</button>
				{/each}
			{/each}
		</div>
		<div id="grid-desc" class="sr-only">Colunas a b c; Linhas 1 2 3; Use setas para mover.</div>
		{#if showHelp}
			<div id="help-panel" class="help-box" role="note">
				<p><strong>Como jogar:</strong> Setas movem; Enter ou Espaço marca; coordenadas: a1..c3.</p>
			</div>
		{/if}
	</div>
	<div class="status mt-2" aria-live="polite">{statusMsg}</div>
	{#if winner}
		<p class="mt-2 h5"><strong>Resultado:</strong> {winner}</p>
	{/if}
	<div class="mt-3 d-flex gap-2">
		<button class="btn btn-primary" on:click={newRound} disabled={redirecting}>Nova rodada</button>
	</div>
	<p class="instructions text-center" style="max-width:38rem;">Setas mover | Enter/Espaço marcar | Home/End canto inicial/final | Nova rodada recomeça.</p>
	<div class="sr-only" aria-live="assertive">{winner ? `Fim de jogo. ${statusMsg}` : ''}</div>
<!-- debug removido -->
</div>

<style>
	.tic-container { max-width: 640px; }
	.board-wrapper { position:relative; width:100%; max-width:38rem; }
	.tic-board { display:grid; grid-template-columns:repeat(3,1fr); width:100%; aspect-ratio:1/1; border:3px solid #3a5f91; border-radius:.9rem; background:#101922; position:relative; box-shadow:0 0 0 3px rgba(43,127,244,.35); }
	.tic-cell { border:1px solid #24496f; display:flex; align-items:center; justify-content:center; font-size:clamp(2.2rem, 8vw, 4.2rem); font-weight:700; cursor:pointer; user-select:none; background:linear-gradient(#16283a,#1e344b); color:#f5f7fa; transition:background .15s, transform .08s; }
	.tic-cell:hover { background:#254362; }
	.tic-cell:active { transform:scale(.96); }
	.tic-cell.filled-X { color:#2aa9ff; }
	.tic-cell.filled-O { color:#ff8b54; }
	.tic-cell.active { box-shadow: inset 0 0 0 3px #ffc107, 0 0 0 2px #000; background:#2d5173; }
	.turn { font-weight:600; font-size:1.6rem; }
	.status { min-height:1.6rem; font-size:1.2rem; }
	.col-labels { display:grid; grid-template-columns:repeat(3,1fr); position:absolute; top:-1.6rem; left:0; width:100%; font-size:.9rem; text-transform:lowercase; color:#c6d6e5; letter-spacing:.05em; }
	.col-labels span { text-align:center; }
	.row-labels { position:absolute; top:0; left:-1.4rem; height:100%; display:grid; grid-template-rows:repeat(3,1fr); font-size:.9rem; color:#c6d6e5; }
	.row-labels span { display:flex; align-items:center; justify-content:center; }
	@media (max-width:480px){ .col-labels{ top:-1.3rem;} .row-labels{ left:-1.1rem;} .turn{ font-size:1.4rem;} }
	.help-box { margin-top:.5rem; background:#152635; border:1px solid #2c5278; padding:.75rem .9rem; border-radius:.6rem; font-size:.95rem; line-height:1.35rem; color:#d5e4f1; }
	.help-box strong { color:#fff; }
	.instructions { font-size:.9rem; color:#9fb9cc; }
	.sr-only { position:absolute; width:1px; height:1px; padding:0; margin:-1px; overflow:hidden; clip:rect(0 0 0 0); white-space:nowrap; border:0; }
/* debug styles removidos */
</style>

*** End Patch

