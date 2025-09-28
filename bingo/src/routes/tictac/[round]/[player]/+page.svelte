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
	let statusMsg = 'Carregando partida...'; // legado (remover depois) – não mais exibido diretamente
	let winner = '';
	let scoreX = 0; let scoreO = 0; let scoreDraw = 0; // placar acumulado
	// Resumo acessível do placar
	$: scoreSummary = `Placar: ${scoreX} vitória${scoreX===1?'':'s'} de X, ${scoreO} vitória${scoreO===1?'':'s'} de O${scoreDraw>0?`, ${scoreDraw} empate${scoreDraw===1?'':'s'}`:''}.`;
	let showHelp = false;
	let showHelpModal = false;
	const colNames = ['a','b','c'];
		const GAME_NAME = 'Jogo da velha inclusivo';
		let ws;
		let redirecting = false;
		let useAppRole = true; // habilita role="application" (cautela: pode reduzir semântica em alguns leitores)

	// Barra única visível (Opção C)
	let visibleInfo = 'Carregando...';
	let lastPiecePlaced = '';
	let lastCoordPlaced = '';
	let tempTimeout = null;

	// Áudio
	/** @type {HTMLAudioElement} */ let clickAudio;
	/** @type {HTMLAudioElement} */ let errorAudio;
	/** @type {HTMLAudioElement} */ let victoryAudio;
	/** @type {HTMLAudioElement} */ let drawAudio;
	/** @type {HTMLAudioElement} */ let defeatAudio;
	let lastOutcome = ''; // evita tocar som duplicado (REST + WS)
	let role = undefined; // usado para injetar role="application" dinamicamente
	let boardInitialized = false; // evita anunciar carregamento inicial
	let liveAnnounce = ''; // única região aria-live consolidada

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
				const prev = board.map(r=>[...r]);
				board = data.board;
				currentTurn = normalizeTurn(data.turn) || currentTurn;
				winner = data.winner || '';
				scoreX = data.scoreX || 0; scoreO = data.scoreO || 0; scoreDraw = data.scoreDraw || 0;
				visibleInfo = computeVisibleInfo(winner, currentTurn);
				if (boardInitialized) updateAnnouncement(prev, board, winner, currentTurn);
				boardInitialized = true;
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
		if (board[r][c] !== '') { setTemporaryInfo('Casa já ocupada.'); errorAudio?.play(); return; }
		if (!isMyTurn()) { setTemporaryInfo('Ainda não é sua vez.'); errorAudio?.play(); return; }
		try {
			const prev = board.map(rr=>[...rr]);
			const res = await fetch(`/api/tictac/${roundParam}/${playerParam}/${r+1}/${c+1}`, { cache: 'no-store' });
			if (!res.ok) return;
			const data = await res.json();
			if (data?.board) {
				board = data.board;
				currentTurn = normalizeTurn(data.turn) || currentTurn;
				winner = data.winner || '';
				// debug removido
				handleMaybeRedirect(data);
				processOutcomeSound();
				updateAnnouncement(prev, board, winner, currentTurn);
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
		try {
			ws = new WebSocket(`${location.protocol === 'https:' ? 'wss' : 'ws'}://${location.host}/ws/tictac/${roundParam}/${playerParam}`);
			ws.onmessage = (ev) => {
				try {
					const data = JSON.parse(ev.data);
					if (data?.board) {
						const prev = board.map(r=>[...r]);
						board = data.board;
						currentTurn = normalizeTurn(data.turn) || currentTurn;
						winner = data.winner || '';
						scoreX = data.scoreX || 0; scoreO = data.scoreO || 0; scoreDraw = data.scoreDraw || 0;
						visibleInfo = computeVisibleInfo(winner, currentTurn);
						if (boardInitialized) updateAnnouncement(prev, board, winner, currentTurn);
						boardInitialized = true;
						// debug removido
						handleMaybeRedirect(data);
						processOutcomeSound();
					}
				} catch(e){ console.error('WS parse', e); }
			};
		} catch(e){ console.error('WS error', e); }

		// Atalhos de ajuda
		const handler = (e) => {
			if(e.key === '?' || (e.shiftKey && e.key === '/')) { showHelpModal = true; e.preventDefault(); }
			else if(e.key === 'Escape' && showHelpModal){ showHelpModal = false; e.preventDefault(); forceFocusMode(); }
		};
		window.addEventListener('keydown', handler);
		return () => window.removeEventListener('keydown', handler);
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

	function processOutcomeSound(){
		if(!winner) return;
		if(winner === lastOutcome) return; // já tocado
		lastOutcome = winner;
		if(winner === 'Empate') { drawAudio?.play(); return; }
		if(winner === localPlayer) { victoryAudio?.play(); }
		else { defeatAudio?.play(); }
	}

	function handleCellClick(r,c){
		clickAudio?.play();
		if (!disabledCell(r,c)) { row = r; col = c; playMove(r,c); focusActiveCell(); }
	}

	function buildUnifiedAnnouncement(lastPiecePlaced, coord, winnerNow, turnAfter){
		if(winnerNow){
			if(winnerNow === 'Empate') return `Empate. ${scoreSummary}`;
			return `Fim: vitória de ${winnerNow}. ${scoreSummary}`;
		}
		let base = '';
		if(lastPiecePlaced){
			base = `${lastPiecePlaced} em ${coord}. `;
		}
		// turno atual é de turnAfter
		if(normalizeTurn(turnAfter) === localPlayer) base += 'É a sua vez.'; else base += `Vez de ${turnAfter}.`;
		return base.trim();
	}

	function updateAnnouncement(prevBoard, newBoard, winnerNow, turnAfter){
		// localizar a nova peça
		let changed = null; let piece = '';
		for(let r=0;r<3;r++){
			for(let c=0;c<3;c++){
				if(prevBoard[r][c] !== newBoard[r][c]){
					if(prevBoard[r][c] === '' && newBoard[r][c] !== '') { changed = {r,c}; piece = newBoard[r][c]; }
				}
			}
			if(changed) break;
		}
		let coord = changed ? `${colNames[changed.c]}${changed.r+1}` : '';
		if(piece){ lastPiecePlaced = piece; lastCoordPlaced = coord; }
		liveAnnounce = buildUnifiedAnnouncement(piece || '', coord, winnerNow, turnAfter);
		visibleInfo = computeVisibleInfo(winnerNow, turnAfter);
	}

	function computeVisibleInfo(winnerNow, turnAfter){
		if(winnerNow){ if(winnerNow==='Empate') return 'Empate.'; return `Vitória de ${winnerNow}.`; }
		if(lastPiecePlaced && lastCoordPlaced){
			return `${lastPiecePlaced} em ${lastCoordPlaced}. ${normalizeTurn(turnAfter)===localPlayer ? 'Sua vez.' : 'Vez de '+turnAfter+'.'}`;
		}
		return normalizeTurn(turnAfter)===localPlayer ? 'Sua vez.' : `Vez de ${turnAfter}.`;
	}

	function setTemporaryInfo(msg){
		visibleInfo = msg;
		if(tempTimeout) clearTimeout(tempTimeout);
		tempTimeout = setTimeout(()=>{ visibleInfo = computeVisibleInfo(winner, currentTurn); }, 2000);
	}

	function handleBoardContainerKey(e){
		if(['Enter',' '].includes(e.key)) { e.preventDefault(); focusActiveCell(); }
	}
</script>

<PageTitle title="Rodada" game="Jogo da Velha" />

<div class="container tic-container py-4 d-flex flex-column align-items-center">
	<div class="scoreboard mb-3" aria-hidden="true">
		<strong>Placar:</strong> X {scoreX} - {scoreO} O {#if scoreDraw>0}<span class="draws">(Empates {scoreDraw})</span>{/if}
	</div>
	<!-- Removido aria-live separado para placar para reduzir verbosidade -->
	<div class="live-bar mb-2">{visibleInfo}</div>
	<div class="board-wrapper my-2">
		<button class="btn btn-sm btn-outline-light" style="margin:0 0 .5rem auto; display:block;" on:click={() => showHelp = !showHelp} aria-expanded={showHelp} aria-controls="help-panel">{showHelp ? 'Ocultar ajuda' : 'Ajuda'}</button>
		<div class="col-labels" aria-hidden="true"><span>a</span><span>b</span><span>c</span></div>
		<div class="row-labels" aria-hidden="true"><span>1</span><span>2</span><span>3</span></div>
		<div class="tic-board" aria-label="Tabuleiro do jogo da velha, rodada {roundParam}. Jogador local {localPlayer}. Pressione Enter para entrar nas casas." bind:this={boardRef} {role}
			tabindex="0"
			on:keydown={handleBoardContainerKey}
			on:focusin={() => { if(useAppRole) role='application'; }}
			on:focusout={(e) => { if(useAppRole && !e.currentTarget.contains(document.activeElement)) role=undefined; }}>
			{#each board as rowVals, r}
				{#each rowVals as cellVal, c}
					<button id={cellId(r,c)} data-r={r} data-c={c} tabindex={row===r && col===c ? 0 : -1}
						aria-pressed={board[r][c] !== ''}
						aria-label={board[r][c] ? `${board[r][c]} ${colNames[c]}${r+1}` : `vazio ${colNames[c]}${r+1}`}
						class="tic-cell {board[r][c] ? `filled-${board[r][c]}`: ''} {row===r && col===c ? 'active' : ''}"
						on:click={() => handleCellClick(r,c)}
						disabled={disabledCell(r,c)}
						on:focus={() => { row = r; col = c; }}
						on:keydown={handleKey}>{cellVal}</button>
				{/each}
			{/each}
		</div>
		{#if showHelp}
			<div id="help-panel" class="help-box" role="note">
				<p><strong>Como jogar:</strong> Setas movem; Enter ou Espaço marca; coordenadas: a1..c3.</p>
			</div>
		{/if}
	</div>
	<!-- status removido em favor da barra única -->
	{#if winner}
		<p class="mt-2 h5"><strong>Resultado:</strong> {winner}</p>
	{/if}
	<div class="mt-3 d-flex gap-2">
		<button class="btn btn-primary" on:click={newRound} disabled={redirecting || !winner} aria-disabled={!winner} title={!winner ? 'Aguarde terminar (vitória ou empate)' : ''}>Nova rodada</button>
	</div>

	{#if showHelpModal}
		<div id="help-modal" class="help-modal" role="dialog" aria-modal="true" aria-labelledby="help-title">
			<div class="help-dialog p-3">
				<h2 id="help-title" class="h5 mb-3">Ajuda do jogo</h2>
				<ul class="small mb-3">
					<li>Use as setas para mover o foco entre as casas.</li>
					<li>Pressione Enter ou Espaço para marcar sua jogada.</li>
					<li>Use Home para ir ao canto superior esquerdo, End para canto inferior direito.</li>
					<li>O anúncio “É a sua vez.” indica que você pode jogar.</li>
					<li>Use o botão Nova rodada para continuar após o fim.</li>
					<li>Pressione Esc para fechar este diálogo.</li>
					<li>Pressione ? a qualquer momento para abrir esta ajuda.</li>
				</ul>
				<button class="btn btn-sm btn-secondary" on:click={() => { showHelpModal=false; focusActiveCell(); }}>Fechar</button>
			</div>
		</div>
	{/if}
	<div class="sr-only" aria-live="polite">{liveAnnounce}</div>
<!-- debug removido -->
</div>

<!-- Áudio: clique, erro, vitória, empate, derrota -->
<audio bind:this={clickAudio} src="/key.mp3" preload="auto"></audio>
<audio bind:this={errorAudio} src="/wrong.wav" preload="auto"></audio>
<audio bind:this={victoryAudio} src="/correct.wav" preload="auto"></audio>
<audio bind:this={drawAudio} src="/mixkit-cooking-bell-ding-1791.wav" preload="auto"></audio>
<audio bind:this={defeatAudio} src="/wrong.wav" preload="auto"></audio>

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
	/* Barra unificada de estado */
	.live-bar { font-weight:600; font-size:1.25rem; background:#142536; padding:.55rem .95rem; border:1px solid #2c4d6b; border-radius:.65rem; min-height:2.4rem; display:flex; align-items:center; }
	.scoreboard { font-size:1.05rem; background:#142536; padding:.4rem .85rem; border:1px solid #2c4d6b; border-radius:.6rem; box-shadow:0 0 0 2px rgba(43,127,244,.15); }
	.scoreboard .draws { color:#94b8cc; font-weight:500; }
	.col-labels { display:grid; grid-template-columns:repeat(3,1fr); position:absolute; top:-1.6rem; left:0; width:100%; font-size:.9rem; text-transform:lowercase; color:#c6d6e5; letter-spacing:.05em; }
	.col-labels span { text-align:center; }
	.row-labels { position:absolute; top:0; left:-1.4rem; height:100%; display:grid; grid-template-rows:repeat(3,1fr); font-size:.9rem; color:#c6d6e5; }
	.row-labels span { display:flex; align-items:center; justify-content:center; }
	@media (max-width:480px){ .col-labels{ top:-1.3rem;} .row-labels{ left:-1.1rem;} .turn{ font-size:1.4rem;} }
	.help-box { margin-top:.5rem; background:#152635; border:1px solid #2c5278; padding:.75rem .9rem; border-radius:.6rem; font-size:.95rem; line-height:1.35rem; color:#d5e4f1; }
	.help-box strong { color:#fff; }
	.instructions { font-size:.9rem; color:#9fb9cc; }
	.sr-only { position:absolute; width:1px; height:1px; padding:0; margin:-1px; overflow:hidden; clip:rect(0 0 0 0); white-space:nowrap; border:0; }
/* Modal de ajuda */
.help-modal { position:fixed; inset:0; background:rgba(0,0,0,.65); display:flex; align-items:center; justify-content:center; z-index:1000; }
.help-dialog { background:#132433; border:1px solid #2c5278; border-radius:.8rem; width:100%; max-width:32rem; box-shadow:0 10px 28px -4px rgba(0,0,0,.6); }
.help-dialog ul { padding-left:1.1rem; }
.help-dialog li { margin:.3rem 0; line-height:1.25rem; }
/* debug styles removidos */
</style>

<!-- EOF -->

