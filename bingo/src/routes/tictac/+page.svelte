<script lang="ts">
	import { onMount, onDestroy } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';
	let creating = false;
	let lastRound = 0;
	let loading = true;
	let errorMsg = '';
	type OpenRound = { round: number; hasPlayerX: boolean; hasPlayerO: boolean };
	let openRounds: OpenRound[] = [];
	let scoreX = 0; let scoreO = 0; let scoreDraw = 0;
	let lastWinner = '';
	let openWs: WebSocket | null = null;
	$: scoreSummary = lastRound>0 ? `Placar acumulado até a rodada ${lastRound}: ${scoreX} vitória${scoreX===1?'':'s'} de X, ${scoreO} vitória${scoreO===1?'':'s'} de O${scoreDraw>0?`, ${scoreDraw} empate${scoreDraw===1?'':'s'}`:''}.` : '';

	function sortOpen() { openRounds = [...openRounds].sort((a: OpenRound,b: OpenRound)=> a.round - b.round); }

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
				try { const lastRes = await fetch(`/api/tictac/${lastRound}`); if(lastRes.ok){ const data = await lastRes.json(); scoreX = data.scoreX||0; scoreO = data.scoreO||0; scoreDraw = data.scoreDraw||0; lastWinner = data.winner || ''; } } catch {}
			}
			const openRes = await fetch('/api/tictac/open');
			if(openRes.ok){ openRounds = await openRes.json(); sortOpen(); }
		} catch(e){ errorMsg = 'Falha ao detectar rodadas.'; }
		finally { loading = false; }
	}

	function connectOpenWs(){
		try {
			openWs = new WebSocket(`${location.protocol==='https:'?'wss':'ws'}://${location.host}/ws/tictac/open`);
			openWs.onmessage = (ev: MessageEvent)=>{
				try { const data = JSON.parse(ev.data); if(Array.isArray(data.rounds)){ openRounds = data.rounds as OpenRound[]; sortOpen(); } } catch {}
			};
			openWs.onclose = ()=> { openWs = null; setTimeout(connectOpenWs, 3000); };
		} catch {}
	}

	async function newRound() {
		creating = true; errorMsg='';
		try {
			const next = lastRound + 1 || 1;
			const res = await fetch(`/api/tictac/${next}/new`);
			if(!res.ok){ errorMsg='Erro criando rodada'; creating=false; return; }
			window.location.href = `/tictac/${next}`;
		} catch(e){ errorMsg='Falha de rede'; creating=false; }
	}

	onMount(()=>{ fetchLatest(); connectOpenWs(); });
	onDestroy(()=> { if(openWs){ openWs.close(); openWs=null; } });
</script>

<PageTitle title="Início" game="Jogo da Velha" />

<div class="container py-4 tic-home" style="max-width:62rem;">
	<div class="d-flex flex-column flex-md-row justify-content-between align-items-md-center gap-3 mb-3">
		<h2 class="h1 mb-0">Partidas</h2>
		<div class="actions d-flex gap-3 flex-wrap align-items-center">
			<button class="btn btn-primary btn-lg" on:click={newRound} disabled={creating} aria-busy={creating}>{creating ? 'Criando...' : 'Nova rodada'}</button>
		</div>
	</div>

	<!-- Total de rodadas criadas -->
	{#if !loading && !errorMsg}
		<p class="total-rounds mb-4" role="status" aria-live="polite">Total de {lastRound} {lastRound === 1 ? 'rodada criada' : 'rodadas criadas'} nesta instância.</p>
	{/if}

	{#if loading}
		<p>Carregando...</p>
	{:else if errorMsg}
		<p class="text-danger">{errorMsg}</p>
	{:else}
		{#if openRounds.length === 0}
			<p>Nenhuma rodada aberta aguardando jogadores.</p>
		{:else}
			<h3 class="h5">Rodadas aguardando jogadores</h3>
			<ul class="list-open-rounds mb-5">
				{#each openRounds as r}
					<li><a href={`/tictac/${r.round}`}>Rodada {r.round}</a><span class="status small ms-2">{r.hasPlayerX && r.hasPlayerO ? 'quase começando' : (r.hasPlayerX || r.hasPlayerO ? '1 jogador presente' : 'vazia')}</span></li>
				{/each}
			</ul>
		{/if}
	{/if}

	<div class="info-card p-4 rounded-3 mb-4">
		<h3 class="h4 mb-3">Como funciona?</h3>
		<div class="guide-grid">
			<section>
				<h4 class="h6">Fluxo básico</h4>
				<ol class="small mb-3">
					<li>Clique em <strong>Nova rodada</strong>.</li>
					<li>Na página da rodada escolha X ou O.</li>
					<li>Compartilhe o link (botão de compartilhar ou copiar).</li>
					<li>O segundo jogador entra e escolhe o outro símbolo.</li>
					<li>O tabuleiro inicia automaticamente.</li>
					<li>Ao terminar (vitória ou empate) use <strong>Nova rodada</strong> dentro da partida.</li>
				</ol>
			</section>
			<section>
				<h4 class="h6">Rodadas abertas</h4>
				<p class="small mb-2">Uma rodada aparece aqui enquanto pelo menos um dos símbolos ainda não foi escolhido. Assim você pode:</p>
				<ul class="small mb-3">
					<li>Entrar e escolher um símbolo livre.</li>
					<li>Pegar o link e enviar para outra pessoa.</li>
					<li>Ver rapidamente se alguém já está aguardando.</li>
				</ul>
			</section>
			<section>
				<h4 class="h6">Acessibilidade</h4>
				<ul class="small mb-3">
					<li>Leitores de tela recebem anúncios consolidados (turno, jogada e resultado).</li>
					<li>Navegação do tabuleiro via setas; Enter ou Espaço marca.</li>
					<li>Atalhos: ? abre ajuda; Esc fecha a ajuda.</li>
				</ul>
			</section>
			<section>
				<h4 class="h6">Placar</h4>
				<p class="small mb-0">O placar é acumulado por sequência de rodadas que você continua usando o botão <em>Nova rodada</em>. Ele avança só quando a rodada termina.</p>
			</section>
		</div>
		<details class="mt-3 small more-help">
			<summary>Mais dicas</summary>
			<ul class="mt-2">
				<li>O primeiro jogador que entra numa nova sequência determina quem começa a primeira rodada.</li>
				<li>Vitória: o vencedor inicia a próxima rodada; empate inverte o próximo turno.</li>
				<li>Se recarregar a página, você retoma a mesma rodada pelo número.</li>
			</ul>
		</details>
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
	.total-rounds { font-size:1rem; color:#d2e7f5; }
	.guide-grid { display:grid; gap:1.25rem; grid-template-columns:repeat(auto-fit,minmax(210px,1fr)); }
	.guide-grid h4 { font-weight:600; }
	.guide-grid ol { padding-left:1.2rem; }
	.guide-grid ul { padding-left:1.1rem; }
	.more-help summary { cursor:pointer; }
	.more-help summary::-webkit-details-marker { display:none; }
	.more-help summary { position:relative; padding-left:1.1rem; }
	.more-help summary:before { content:'➕'; position:absolute; left:0; }
	.more-help summary:before { content:'➕'; position:absolute; left:0; }
	.more-help[open] summary:before { content:'➖'; }
</style>