<script>
// @ts-nocheck  // desabilita checagem TypeScript neste arquivo JS para evitar avisos de 'implicit any'
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';
	import Scoreboard from '$lib/tictac/Scoreboard.svelte';
	import StatusBar from '$lib/tictac/StatusBar.svelte';
	import GameBoard from '$lib/tictac/GameBoard.svelte';
	import HelpModal from '$lib/tictac/HelpModal.svelte';
	import NumberKeyHandler from '$lib/tictac/NumberKeyHandler.svelte';
	import AudioManager from '$lib/tictac/AudioManager.svelte';
	import AdManager from '$lib/ads/AdManager.svelte';
	import AdProtection from '$lib/ads/AdProtection.svelte';
	import { adConfig } from '$lib/ads/adConfig.js';

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
	let hasPlayerX = false; let hasPlayerO = false; // presença dos jogadores
	let opponentAnnounced = false; // evita anúncios duplicados
	// Resumo acessível do placar
	$: scoreSummary = `Placar: ${scoreX} vitória${scoreX===1?'':'s'} de X, ${scoreO} vitória${scoreO===1?'':'s'} de O${scoreDraw>0?`, ${scoreDraw} empate${scoreDraw===1?'':'s'}`:''}.`;
	let lastScoreAnnounced = '';
	let isMobile = false; // heurística ampliada para iOS / dispositivos touch
	$: hasMoves = board.some(r => r.some(c => c !== ''));
	let shareMsg = '';
	let showHelp = false;
	let showHelpModal = false;
	const colNames = ['a','b','c'];
	const GAME_NAME = 'Jogo da velha inclusivo';
	let ws;
	let redirecting = false;
	// removido role application dinâmico para simplificar e evitar avisos a11y

	// Barra única visível (Opção C)
	let visibleInfo = 'Carregando...';
	let lastPiecePlaced = '';
	let lastCoordPlaced = '';
	// temporizador para mensagens temporárias
	let tempTimeout = null;

	// Áudio
	let audioManager;
	let lastOutcome = ''; // evita tocar som duplicado (REST + WS)
	let boardInitialized = false; // evita anunciar carregamento inicial
	let liveAnnounce = ''; // única região aria-live consolidada

	const cellId = (r, c) => `cell-${r}-${c}`;

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
				hasPlayerX = !!data.playerX; hasPlayerO = !!data.playerO;
				scoreX = data.scoreX || 0; scoreO = data.scoreO || 0; scoreDraw = data.scoreDraw || 0;
				visibleInfo = computeVisibleInfo(winner, currentTurn);
				if (boardInitialized) updateAnnouncement(prev, board, winner, currentTurn);
				if(scoreSummary !== lastScoreAnnounced){ lastScoreAnnounced = scoreSummary; }
				boardInitialized = true;
				// debug removido
				handleMaybeRedirect(data);
				maybeAnnounceOpponent();
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
	async function playMove(r, c) {
		if (winner) return;
		// Não bloquear no front: backend valida oponente / turno
		if (board[r][c] !== '') { setTemporaryInfo('Casa já ocupada.'); audioManager?.playError(); return; }
		try {
			const prev = board.map(rr=>[...rr]);
			const res = await fetch(`/api/tictac/${roundParam}/${playerParam}/${r+1}/${c+1}`, { cache: 'no-store' });
			if (!res.ok) {
				if(res.status === 412){ setTemporaryInfo('Aguardando oponente.'); }
				return;
			}
			const data = await res.json();
			if (data?.board) {
				board = data.board;
				currentTurn = normalizeTurn(data.turn) || currentTurn;
				winner = data.winner || '';
				// debug removido
				handleMaybeRedirect(data);
				processOutcomeSound();
				updateAnnouncement(prev, board, winner, currentTurn);
				if(scoreSummary !== lastScoreAnnounced){ lastScoreAnnounced = scoreSummary; }
			}
		} catch(e){ console.error('Erro jogada', e); }
	}

	// Removido disabledCell: não desabilitamos mais visualmente; validação só em playMove

	function reset() { window.location.reload(); }

	let boardRef = null;

	function focusActiveCell() {
		const el = document.getElementById(cellId(row,col));
		el?.focus();
	}

	function detectMobileShare(){
		if (typeof window === 'undefined') return false;
		const ua = navigator.userAgent || '';
		const touch = ('ontouchstart' in window) || (navigator.maxTouchPoints || 0) > 0;
		const coarse = window.matchMedia && window.matchMedia('(pointer:coarse)').matches;
		const uaMatch = /Mobi|Android|iPhone|iPad|iPod/i.test(ua);
		return !!(navigator.share) || (touch && (coarse || uaMatch));
	}

	function bothPlayersPresent(){ return hasPlayerX && hasPlayerO; }
	function iAmX(){ return localPlayer === 'X'; }
	function maybeAnnounceOpponent(){
		if (bothPlayersPresent() && !opponentAnnounced){
			const opp = localPlayer === 'X' ? 'O' : 'X';
			liveAnnounce = `Oponente (${opp}) entrou. ${normalizeTurn(currentTurn)===localPlayer ? 'É a sua vez.' : 'Vez de '+currentTurn+'.'}`;
			visibleInfo = computeVisibleInfo(winner, currentTurn);
			opponentAnnounced = true;
		}
	}

// QR code removido desta tela (deve existir apenas na seleção de símbolo)

	onMount(() => {
		// executar parte assíncrona sem tornar callback async (evita Promise de cleanup)
		(async () => { await ensureRound(); })();
		// garantir que visibleInfo reflita turno inicial mesmo se ainda vazio
		visibleInfo = computeVisibleInfo(winner, currentTurn);
		// detectar mobile/share (heurística ampliada)
		isMobile = detectMobileShare();
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
						hasPlayerX = !!data.playerX; hasPlayerO = !!data.playerO;
						scoreX = data.scoreX || 0; scoreO = data.scoreO || 0; scoreDraw = data.scoreDraw || 0;
						visibleInfo = computeVisibleInfo(winner, currentTurn);
						if (boardInitialized) updateAnnouncement(prev, board, winner, currentTurn);
						boardInitialized = true;
						// debug removido
						handleMaybeRedirect(data);
						processOutcomeSound();
						maybeAnnounceOpponent();
					}
				} catch(e){ console.error('WS parse', e); }
			};
		} catch(e){ console.error('WS error', e); }

		// Atalhos de ajuda
		const handler = (e) => {
			if(e.key === '?' || (e.shiftKey && e.key === '/')) { showHelpModal = true; e.preventDefault(); }
			else if(e.key === 'Escape' && showHelpModal){ showHelpModal = false; e.preventDefault(); focusActiveCell(); }
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
		audioManager?.playOutcomeSound(winner, localPlayer);
	}

	function handleCellClick(r, c){
		audioManager?.playClick();
		row = r; col = c; // foco lógico
		playMove(r,c); // backend e playMove cuidam das regras (oponente, turno, ocupado)
		focusActiveCell();
	}

	async function shareInvite(){
		const url = `${location.origin}/tictac/${roundParam}`;
		try {
			if(navigator.share){
				await navigator.share({ title: 'Jogo da Velha', text: 'Entre na rodada do jogo da velha inclusivo', url });
				shareMsg = 'Link compartilhado';
			} else {
				await navigator.clipboard.writeText(url);
				shareMsg = 'Link copiado';
			}
		} catch(e){ shareMsg = 'Não foi possível compartilhar'; }
		setTimeout(()=> shareMsg='', 2500);
	}

	function buildUnifiedAnnouncement(lastPiecePlaced, coord, winnerNow, turnAfter){
		if(!bothPlayersPresent()) return 'Aguardando oponente.';
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
		if(!bothPlayersPresent()) return 'Aguardando oponente...';
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

	// Handlers dos componentes
	function handleCellClickFromBoard(event) {
		const { row: r, col: c } = event.detail;
		handleCellClick(r, c);
	}

	function handleCellFocusFromBoard(event) {
		const { row: r, col: c } = event.detail;
		row = r; col = c;
	}

	function handleKeydownFromBoard(event) {
		handleKey(event.detail.event);
	}

	/** @param {KeyboardEvent} e */
	function handleKey(e) {
		if (winner) return;
		const key = e.key;
		if (!['ArrowUp','ArrowDown','ArrowLeft','ArrowRight','Home','End','Enter',' '].includes(key)) return;
		const target = e.currentTarget;
		const r = parseInt(target?.dataset.r || '0');
		const c = parseInt(target?.dataset.c || '0');
		if (key === 'Enter' || key === ' ') {
			row = r; col = c; playMove(r,c); e.preventDefault(); return; 
		}
		
		// Navegação por setas apenas (números são tratados pelo NumberKeyHandler)
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

	function handleNumberKeyFromHandler(event) {
		const { row: r, col: c } = event.detail;
		row = r; col = c;
		playMove(r, c);
		focusActiveCell();
	}

	function handleHelpToggle() {
		showHelpModal = true;
	}

	function handleHelpClose() {
		showHelpModal = false;
		focusActiveCell();
	}

	function handleHelpModalClose() {
		showHelpModal = false;
		focusActiveCell();
	}

	function handleHelpQuickToggle() {
		showHelp = !showHelp;
	}

	// container não mais foco direto; interação via células
</script>

<PageTitle title="Rodada" game="Jogo da Velha" />

<!-- Proteção contra anúncios intrusivos -->
<AdProtection protection="strict" />

<div class="container tic-container py-4 d-flex flex-column align-items-center game-area">
	<!-- Anúncio no topo (opcional, somente em desktop) -->
	<AdManager 
		placement="top" 
		format="banner" 
		adSlot={adConfig.slots.topBanner}
		priority="low"
	/>
	
	<div class="scoreboard-container mb-3">
		<Scoreboard {scoreX} {scoreO} {scoreDraw} />
	</div>
	<!-- Placar em um único elemento para leitura linear -->
	<div class="mb-2">
		<StatusBar {visibleInfo} />
	</div>
	<div class="board-wrapper my-2">
		<HelpModal 
			showModal={showHelpModal} 
			showQuickHelp={showHelp}
			on:close={handleHelpModalClose}
			on:toggleQuick={handleHelpQuickToggle}
		/>
		<GameBoard 
			{board} 
			{row} 
			{col}
			{roundParam}
			localPlayer={localPlayer}
			on:cellClick={handleCellClickFromBoard}
			on:cellFocus={handleCellFocusFromBoard}
			on:keydown={handleKeydownFromBoard}
		/>
	</div>
	<!-- status removido em favor da barra única -->
	{#if winner}
		<p class="mt-2 h5"><strong>Resultado:</strong> {winner}</p>
	{/if}
	
	<!-- Anúncio retângulo após o resultado -->
	<AdManager 
		placement="auto" 
		format="rectangle" 
		adSlot={adConfig.slots.rectangle}
		priority="medium"
	/>
	
	<div class="mt-3 d-flex gap-2 game-controls">
		{#if isMobile && !hasMoves && !winner}
			<button class="btn btn-warning" on:click={shareInvite}>Compartilhar convite</button>
			{#if shareMsg}<span class="small text-success" aria-live="polite">{shareMsg}</span>{/if}
		{:else}
			<button class="btn btn-primary" on:click={newRound} disabled={redirecting || !winner} aria-disabled={!winner} title={!winner ? 'Aguarde terminar (vitória ou empate)' : ''}>Nova rodada</button>
		{/if}
	</div>

	<!-- Anúncio na parte inferior -->
	<AdManager 
		placement="bottom" 
		format="banner" 
		adSlot={adConfig.slots.bottomBanner}
		priority="low"
	/>

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

<!-- Modal de QR removido: QR apenas na tela de seleção -->
	<div class="sr-only" aria-live="polite">{liveAnnounce}</div>
<!-- debug removido -->
</div>

<!-- Componentes de gerenciamento -->
<NumberKeyHandler 
	{showHelpModal} 
	{winner}
	on:numberKey={handleNumberKeyFromHandler}
	on:helpToggle={handleHelpToggle}
	on:helpClose={handleHelpClose}
/>

<!-- Áudio: clique, erro, vitória, empate, derrota -->
<AudioManager bind:this={audioManager} {lastOutcome} />

<style>
	.tic-container { 
		max-width: 720px; /* Aumentado de 640px */
		min-width: 320px;
		width: 100%;
	}
	
	.game-area {
		background: var(--game-bg, transparent);
		isolation: isolate;
		position: relative;
	}
	
	.game-controls {
		isolation: isolate;
		position: relative;
		z-index: 10;
	}
	
	.scoreboard-container {
		width: 100%;
		display: flex;
		justify-content: center;
	}
	
	.btn-warning { 
		background: #c98219; 
		border-color: #c98219; 
	}
	
	.btn-warning:hover { 
		background: #b87415; 
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

	@media (max-width: 480px) {
		.tic-container { 
			padding-left: 1rem; 
			padding-right: 1rem; 
		}
		
		.scoreboard-container {
			width: 100%;
			max-width: none;
		}
	}
	
	@media (max-width: 360px) {
		.tic-container {
			padding-left: 0.75rem;
			padding-right: 0.75rem;
		}
	}
</style>

<!-- EOF -->

