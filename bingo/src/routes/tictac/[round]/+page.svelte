<script>
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';

	const p = get(page);
	let roundParam = p.params.round;
	let loading = true;
	let errorMsg = '';
	let winner = '';
	let turn = '';
	let board = [['','',''],['','',''],['','','']];
	let playerXBusy = false;
	let playerOBusy = false;
	let copyMsg = '';
	let scoreX = 0; let scoreO = 0; let scoreDraw = 0;
	$: scoreSummary = `Placar acumulado: ${scoreX} vitória${scoreX===1?'':'s'} de X, ${scoreO} vitória${scoreO===1?'':'s'} de O${scoreDraw>0?`, ${scoreDraw} empate${scoreDraw===1?'':'s'}`:''}.`;
let isMobile = false; // heurística ampliada
 let shareMsg = '';
let shareLabel = 'Compartilhar link da rodada';

	async function loadRound() {
		try {
			const res = await fetch(`/api/tictac/${roundParam}`);
			if (res.status === 404) { errorMsg = 'Rodada não encontrada.'; return; }
			if(!res.ok) { errorMsg = 'Erro ao obter rodada.'; return; }
			const data = await res.json();
			board = data.board || board;
			winner = data.winner || '';
			turn = data.turn || '';
			playerXBusy = !!data.playerX;
			playerOBusy = !!data.playerO;
			scoreX = data.scoreX || 0; scoreO = data.scoreO || 0; scoreDraw = data.scoreDraw || 0;
		} catch(e){ errorMsg = 'Falha de rede.'; }
		finally { loading = false; }
	}

	function shareLink() {
		const url = `${location.origin}/tictac/${roundParam}`;
		try {
			navigator.clipboard.writeText(url);
			copyMsg = 'Link copiado!';
			setTimeout(()=> copyMsg='', 2500);
		} catch { copyMsg = 'Não foi possível copiar.'; }
	}

	function detectMobileShare(){
		if (typeof window === 'undefined') return false;
		const ua = navigator.userAgent || '';
		const touch = ('ontouchstart' in window) || (navigator.maxTouchPoints || 0) > 0;
		const coarse = window.matchMedia && window.matchMedia('(pointer:coarse)').matches;
		const uaMatch = /Mobi|Android|iPhone|iPad|iPod/i.test(ua);
		return !!(navigator.share) || (touch && (coarse || uaMatch));
	}
	onMount(loadRound);
	onMount(()=> { isMobile = detectMobileShare(); });

 async function shareInvite(){
 	const url = `${location.origin}/tictac/${roundParam}`;
 	try {
 		if(navigator.share){ await navigator.share({ title:'Jogo da Velha', text:'Entre na rodada do jogo da velha inclusivo', url}); shareMsg='Link compartilhado'; }
 		else { await navigator.clipboard.writeText(url); shareMsg='Link copiado'; }
 	} catch { shareMsg='Não foi possível compartilhar'; }
 	setTimeout(()=> shareMsg='', 2500);
 }
</script>

<PageTitle title="Escolher jogador" game="Jogo da velha inclusivo" />

<div class="container py-4" style="max-width:52rem;">
	<h2 class="h1 mb-3">Rodada {roundParam}</h2>
	{#if !loading && !errorMsg}
		<div class="scoreboard mb-3" aria-hidden="true"><strong>Placar:</strong> X {scoreX} - {scoreO} O {#if scoreDraw>0}<span class="draws">(Empates {scoreDraw})</span>{/if}</div>
		<div class="sr-only" aria-live="polite">{scoreSummary}</div>
	{/if}
	{#if loading}
		<p>Carregando...</p>
	{:else if errorMsg}
		<p class="text-danger">{errorMsg}</p>
	{:else}
		<p class="lead">Escolha seu marcador para esta rodada:</p>
		<div class="d-flex gap-3 flex-wrap mb-3" role="group" aria-label="Escolher marcador">
			<button class="btn btn-primary btn-lg" disabled={playerXBusy} aria-disabled={playerXBusy} aria-label={playerXBusy ? 'Jogador X já escolhido' : 'Jogar como X'} on:click={()=> !playerXBusy && (location.href=`/tictac/${roundParam}/x`)}>X {#if playerXBusy}<span class="badge bg-dark ms-2">ocupado</span>{/if}</button>
			<button class="btn btn-warning btn-lg" disabled={playerOBusy} aria-disabled={playerOBusy} aria-label={playerOBusy ? 'Jogador O já escolhido' : 'Jogar como O'} on:click={()=> !playerOBusy && (location.href=`/tictac/${roundParam}/o`)}>O {#if playerOBusy}<span class="badge bg-dark ms-2">ocupado</span>{/if}</button>
			{#if isMobile && (!playerXBusy || !playerOBusy)}
				{@html (()=>{ shareLabel = (playerXBusy || playerOBusy) ? 'Compartilhar link para o segundo jogador' : 'Compartilhar link da rodada'; return '' })()}
				<button class="btn btn-outline-info btn-lg" aria-label={shareLabel} on:click={shareInvite}>Compartilhar</button>
			{/if}
		</div>
		{#if shareMsg}<div class="small text-success" aria-live="polite">{shareMsg}</div>{/if}
		<div class="mb-4 d-flex align-items-center gap-2 flex-wrap">
			<button class="btn btn-outline-info" on:click={shareLink}>Copiar link da rodada</button>
			{#if copyMsg}<span aria-live="polite" class="small text-success">{copyMsg}</span>{/if}
		</div>
		<div class="preview" aria-hidden="true">
			<div class="mini-board">
				{#each board as rowVals}
					<div class="mini-row">{#each rowVals as c}<span>{c || '-'}</span>{/each}</div>
				{/each}
			</div>
			<p class="status small mt-2">{winner ? `Resultado: ${winner}` : (turn ? `Vez atual: ${turn}` : (playerXBusy || playerOBusy ? 'Aguardando outro jogador' : 'Aguardando seleção de marcadores'))}</p>
		</div>
	{/if}
	<div class="mt-5">
		<a href="/tictac" class="btn btn-outline-secondary">Voltar ao início</a>
	</div>
</div>

<style>
	.mini-board { display:inline-flex; flex-direction:column; border:2px solid #444; border-radius:.4rem; overflow:hidden; }
	.mini-row { display:flex; }
	.mini-row span { width:2.4rem; height:2.4rem; display:inline-flex; align-items:center; justify-content:center; font-weight:600; background:#1f1f1f; color:#eee; border-right:1px solid #333; border-bottom:1px solid #333; font-size:1.2rem; }
	.mini-row span:last-child { border-right:none; }
	.mini-row:last-child span { border-bottom:none; }
	.badge { font-size:.9rem; }
	.scoreboard { font-size:1rem; background:#142536; padding:.35rem .75rem; border:1px solid #2c4d6b; border-radius:.55rem; display:inline-block; }
	.scoreboard .draws { color:#94b8cc; font-weight:500; }
	.btn-outline-info { --bs-btn-color:#20b3ff; --bs-btn-border-color:#20b3ff; }
	.sr-only { position:absolute; width:1px; height:1px; padding:0; margin:-1px; overflow:hidden; clip:rect(0 0 0 0); white-space:nowrap; border:0; }
</style>
