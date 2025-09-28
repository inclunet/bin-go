<script>
	import { onMount, onDestroy } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';
	let creating = false;
	let lastRound = 0;
	let loading = true;
	let errorMsg = '';
	let openRounds = [];
	let scoreX = 0; let scoreO = 0; let scoreDraw = 0;
	let openWs;
	$: scoreSummary = lastRound>0 ? `Placar acumulado até a rodada ${lastRound}: ${scoreX} vitória${scoreX===1?'':'s'} de X, ${scoreO} vitória${scoreO===1?'':'s'} de O${scoreDraw>0?`, ${scoreDraw} empate${scoreDraw===1?'':'s'}`:''}.` : '';

	function sortOpen() { openRounds = [...openRounds].sort((a,b)=> a.round - b.round); }

	async function fetchLatest() {
		loading = true;
		try {
			let r = 1;
			for(; r <= 50; r++) {
				const res = await fetch(`/api/tictac/${r}`);
				if (res.status === 404) break;
				if(!res.ok) break;
			}
			lastRound = r-1;
			if(lastRound > 0){
				try { const lastRes = await fetch(`/api/tictac/${lastRound}`); if(lastRes.ok){ const data = await lastRes.json(); scoreX = data.scoreX||0; scoreO = data.scoreO||0; scoreDraw = data.scoreDraw||0; } } catch {}
			}
			const openRes = await fetch('/api/tictac/open');
			if(openRes.ok){ openRounds = await openRes.json(); sortOpen(); }
		} catch(e){ errorMsg = 'Falha ao detectar rodadas.'; }
		finally { loading = false; }
	}

	function connectOpenWs(){
		try {
			openWs = new WebSocket(`${location.protocol==='https:'?'wss':'ws'}://${location.host}/ws/tictac/open`);
			openWs.onmessage = (ev)=>{
				try { const data = JSON.parse(ev.data); if(Array.isArray(data.rounds)){ openRounds = data.rounds; sortOpen(); } } catch {}
			};
			openWs.onclose = ()=> { openWs = null; setTimeout(connectOpenWs, 3000); };
		} catch {}
	}

	async function newRound() {
		creating = true; errorMsg='';
		try { const next = lastRound + 1 || 1; const res = await fetch(`/api/tictac/${next}/new`); if(!res.ok){ errorMsg='Erro criando rodada'; creating=false; return; } window.location.href = `/tictac/${next}`; }
		catch(e){ errorMsg='Falha de rede'; creating=false; }
	}

	onMount(()=>{ fetchLatest(); connectOpenWs(); });
	onDestroy(()=> { if(openWs){ openWs.close(); openWs=null; } });
</script>

<PageTitle title="Início" game="Jogo da Velha" />

<div class="container py-4 tic-home" style="max-width:62rem;">
	<div class="d-flex flex-column flex-md-row justify-content-between align-items-md-center gap-3 mb-4">
		<h2 class="h1 mb-0">Partidas</h2>
		<div class="actions d-flex gap-3 flex-wrap align-items-center">
			<button class="btn btn-primary btn-lg" on:click={newRound} disabled={creating} aria-busy={creating}>{creating ? 'Criando...' : 'Nova rodada'}</button>
		</div>
	</div>

	{#if loading}
		<p>Carregando...</p>
	{:else if errorMsg}
		<p class="text-danger">{errorMsg}</p>
	{:else}
		{#if openRounds.length === 0}
			<p>Nenhuma rodada aberta aguardando jogadores.</p>
		{:else}
			<h3 class="h5">Rodadas aguardando jogadores</h3>
			<ul class="list-open-rounds mb-4">
				{#each openRounds as r}
					<li><a href={`/tictac/${r.round}`}>Rodada {r.round}</a><span class="status small ms-2">{r.hasPlayerX && r.hasPlayerO ? 'quase começando' : (r.hasPlayerX || r.hasPlayerO ? '1 jogador presente' : 'vazia')}</span></li>
				{/each}
			</ul>
		{/if}
	{/if}

	<div class="info-card p-4 rounded-3 mb-4">
		<h3 class="h4">Como funciona?</h3>
		<ol class="mb-0 small">
			<li>Clique em "Nova rodada".</li>
			<li>Escolha X ou O.</li>
			<li>Compartilhe o link para o outro jogador.</li>
			<li>O tabuleiro atualiza em tempo real.</li>
			<li>Use "Nova rodada" para continuar a mesma sequência.</li>
		</ol>
	</div>
</div>

<style>
	.tic-home { color:#fff; }
	.info-card { background:rgba(255,255,255,.05); border:1px solid rgba(255,255,255,.15); backdrop-filter:blur(3px); }
	.info-card ol { padding-left:1.2rem; }
	.info-card li { margin:.25rem 0; }
	.list-open-rounds { list-style:none; padding-left:0; margin:0; }
	.list-open-rounds li { margin:.4rem 0; }
	.list-open-rounds a { font-weight:600; text-decoration:none; color:#39a8ff; }
	.list-open-rounds a:hover { text-decoration:underline; }
	.status { color:#b4c9d6; }
</style>