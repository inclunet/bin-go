<script>
	import { onMount } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';
	let creating = false;
	let lastRound = 0;
	let loading = true;
	let errorMsg = '';

	async function fetchLatest() {
		loading = true;
		try {
			// tentativa de descobrir última rodada: pedir rounds sequenciais até falhar (simplês) – assumindo poucas.
			// Melhor seria um endpoint dedicado; por enquanto brute force até 50.
			let r = 1;
			for(; r <= 50; r++) {
				const res = await fetch(`/api/tictac/${r}`);
				if (res.status === 404) break;
				if(!res.ok) break;
			}
			lastRound = r-1;
		} catch(e){ errorMsg = 'Falha ao detectar rodadas.'; }
		finally { loading = false; }
	}

	async function newRound() {
		creating = true;
		errorMsg='';
		try {
			const next = lastRound + 1 || 1;
			const res = await fetch(`/api/tictac/${next}/new`);
			if(!res.ok) { errorMsg='Erro criando rodada'; creating=false; return; }
			window.location.href = `/tictac/${next}`;
		} catch(e){ errorMsg='Falha de rede'; creating=false; }
	}

	onMount(fetchLatest);
</script>

<PageTitle title="Início" game="Jogo da Velha" />

<div class="container py-4 tic-home" style="max-width:62rem;">
	<div class="d-flex flex-column flex-md-row justify-content-between align-items-md-center gap-3 mb-4">
		<h2 class="h1 mb-0">Partidas</h2>
		<div class="actions d-flex gap-3 flex-wrap align-items-center">
			<button class="btn btn-primary btn-lg" on:click={newRound} disabled={creating} aria-busy={creating}>{creating ? 'Criando...' : 'Nova rodada'}</button>
			{#if !loading && lastRound > 0}
				<a class="btn btn-outline-light" href={`/tictac/${lastRound}`}>Última ({lastRound})</a>
			{/if}
		</div>
	</div>

	{#if loading}
		<p>Carregando estatísticas...</p>
	{:else if errorMsg}
		<p class="text-danger">{errorMsg}</p>
	{:else if lastRound === 0}
		<p>Nenhuma rodada criada ainda. Clique em "Nova rodada" para iniciar.</p>
	{:else}
		<p class="mb-4">Total de {lastRound} {lastRound === 1 ? 'rodada' : 'rodadas'} criadas.</p>
	{/if}

	<div class="info-card p-4 rounded-3 mb-4">
		<h3 class="h4">Como funciona?</h3>
		<ol class="mb-0 small">
			<li>Clique em "Nova rodada".</li>
			<li>Na página da rodada escolha X ou O.</li>
			<li>Compartilhe o link da rodada.</li>
			<li>O tabuleiro atualiza em tempo real.</li>
			<li>Use "Nova rodada" para continuar jogando.</li>
		</ol>
	</div>
</div>

<style>
	.tic-home { color:#fff; }
	.info-card { background:rgba(255,255,255,.05); border:1px solid rgba(255,255,255,.15); backdrop-filter:blur(3px); }
	.info-card ol { padding-left:1.2rem; }
	.info-card li { margin:.25rem 0; }
</style>