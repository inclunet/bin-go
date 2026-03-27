<script>
	import { onMount, onDestroy } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';
	let creating = false;
	let lastRound = 0;
	let loading = true;
	let errorMsg = '';
	/** @typedef {{round:number, hasPlayerA:boolean, hasPlayerB:boolean}} OpenRound */
	/** @type {OpenRound[]} */
	let openRounds = [];
	let scoreA = 0; let scoreB = 0; let scoreDraw = 0;
	let lastWinner = '';
	/** @type {WebSocket | null} */
	let openWs = null;
	$: scoreSummary = lastRound>0 ? `Placar acumulado até a rodada ${lastRound}: ${scoreA} vitória${scoreA===1?'':'s'} do Jogador A, ${scoreB} vitória${scoreB===1?'':'s'} do Jogador B${scoreDraw>0?`, ${scoreDraw} empate${scoreDraw===1?'':'s'}`:''}.` : '';

	function sortOpen() { openRounds = [...openRounds].sort((a,b)=> a.round - b.round); }

	async function fetchLatest() {
		loading = true;
		try {
			let r = 1;
			for(; r <= 50; r++) {
				const res = await fetch(`/api/battleship/${r}`);
				if (res.status === 404) break;
				if(!res.ok) break;
			}
			lastRound = r-1;
			if(lastRound > 0){
				try { 
					const lastRes = await fetch(`/api/battleship/${lastRound}`); 
					if(lastRes.ok){ 
						const data = await lastRes.json(); 
						scoreA = data.scoreA||0; 
						scoreB = data.scoreB||0; 
						scoreDraw = data.scoreDraw||0; 
						lastWinner = data.winner || ''; 
					} 
				} catch {}
			}
			const openRes = await fetch('/api/battleship/open');
			if(openRes.ok){ openRounds = await openRes.json(); sortOpen(); }
		} catch(e){ errorMsg = 'Falha ao detectar partidas.'; }
		finally { loading = false; }
	}

	function connectOpenWs(){
		try {
			openWs = new WebSocket(`${location.protocol==='https:'?'wss':'ws'}://${location.host}/ws/battleship/open`);
			openWs.onmessage = (ev)=>{
				try { 
					const data = JSON.parse(ev.data); 
					if(Array.isArray(data.rounds)){ 
						openRounds = data.rounds; 
						sortOpen(); 
					} 
				} catch {}
			};
			openWs.onclose = ()=> { openWs = null; setTimeout(connectOpenWs, 3000); };
		} catch {}
	}

	async function newRound() {
		creating = true; errorMsg='';
		try {
			const next = lastRound + 1 || 1;
			const res = await fetch(`/api/battleship/${next}/new`);
			if(!res.ok){ errorMsg='Erro criando partida'; creating=false; return; }
			window.location.href = `/battleship/${next}`;
		} catch(e){ errorMsg='Falha de rede'; creating=false; }
	}

	onMount(()=>{ fetchLatest(); connectOpenWs(); });
	onDestroy(()=> { if(openWs){ openWs.close(); openWs=null; } });
</script>

<PageTitle title="Início" game="Batalha Naval" />

<div class="container py-4 battleship-home" style="max-width:62rem;">
	<div class="d-flex flex-column flex-md-row justify-content-between align-items-md-center gap-3 mb-3">
		<h2 class="h1 mb-0">Batalha Naval</h2>
		<div class="actions d-flex gap-3 flex-wrap align-items-center">
			<button class="btn btn-primary btn-lg" on:click={newRound} disabled={creating} aria-busy={creating}>
				{creating ? 'Criando...' : 'Nova partida'}
			</button>
		</div>
	</div>

	<!-- Total de rodadas criadas -->
	{#if !loading && !errorMsg}
		<p class="total-rounds mb-4" role="status" aria-live="polite">
			Total de {lastRound} {lastRound === 1 ? 'partida criada' : 'partidas criadas'} nesta sessão.
		</p>
	{/if}

	{#if loading}
		<p>Carregando...</p>
	{:else if errorMsg}
		<p class="text-danger">{errorMsg}</p>
	{:else}
		{#if openRounds.length === 0}
			<p>Nenhuma partida aguardando jogadores no momento.</p>
		{:else}
			<h3 class="h5">Partidas aguardando jogadores</h3>
			<ul class="list-open-rounds mb-5">
				{#each openRounds as r}
					<li>
						<a href={`/battleship/${r.round}`}>Partida {r.round}</a>
						<span class="status small ms-2">
							{r.hasPlayerA && r.hasPlayerB ? 'pronta para começar' : 
							 (r.hasPlayerA || r.hasPlayerB ? '1 jogador aguardando' : 'aguardando jogadores')}
						</span>
					</li>
				{/each}
			</ul>
		{/if}
	{/if}

	<div class="info-card p-4 rounded-3 mb-4">
		<h3 class="h4 mb-3">Como jogar Batalha Naval?</h3>
		<div class="guide-grid">
			<section>
				<h4 class="h6">Iniciando uma partida</h4>
				<ol class="small mb-3">
					<li>Clique em <strong>Nova partida</strong> para criar uma partida.</li>
					<li>Na página da partida escolha ser Jogador A ou Jogador B.</li>
					<li>Compartilhe o link com outra pessoa usando o botão de compartilhar.</li>
					<li>O segundo jogador entra no link e escolhe o jogador restante.</li>
					<li>Ambos posicionam seus navios no tabuleiro.</li>
					<li>O jogo inicia automaticamente quando ambos estão prontos.</li>
				</ol>
			</section>
			<section>
				<h4 class="h6">Partidas abertas</h4>
				<p class="small mb-2">Uma partida aparece na lista acima enquanto ainda há vagas para jogadores. Você pode:</p>
				<ul class="small mb-3">
					<li>Entrar diretamente e escolher um jogador disponível (A ou B).</li>
					<li>Copiar o link da partida e enviar para um amigo.</li>
					<li>Ver quantos jogadores já estão na partida.</li>
					<li>Aguardar até que os dois jogadores estejam presentes para começar.</li>
				</ul>
			</section>
			<section>
				<h4 class="h6">Controles e acessibilidade</h4>
				<ul class="small mb-3">
					<li><strong>Navegação:</strong> Use as setas do teclado para mover entre as células do tabuleiro.</li>
					<li><strong>Atirar:</strong> Pressione Enter ou Espaço na célula desejada para disparar.</li>
					<li><strong>Posicionar navios:</strong> Durante a preparação, use Enter/Espaço para posicionar e R para rotacionar.</li>
					<li><strong>Atalhos:</strong> Pressione ? para abrir ajuda; Esc para fechar diálogos.</li>
					<li><strong>Leitores de tela:</strong> Recebem anúncios automáticos sobre tiros, acertos e status do jogo.</li>
				</ul>
			</section>
			<section>
				<h4 class="h6">Sistema de placar</h4>
				<p class="small mb-0">O placar é mantido ao longo de uma sequência de partidas. Cada vez que você clica em "Nova partida" após terminar um jogo, o placar continua acumulando vitórias de cada jogador.</p>
			</section>
		</div>
		<details class="mt-3 small more-help">
			<summary>Dicas avançadas</summary>
			<ul class="mt-2">
				<li><strong>Estratégia:</strong> Espalhe seus navios para dificultar a localização pelo oponente.</li>
				<li><strong>Primeiro turno:</strong> O Jogador A sempre começa atirando primeiro.</li>
				<li><strong>Reconexão:</strong> Se recarregar a página, você retoma automaticamente a mesma partida.</li>
				<li><strong>Coordenadas:</strong> O tabuleiro usa coordenadas A-J (horizontal) e 1-10 (vertical).</li>
			</ul>
		</details>
	</div>
</div>

<style>
	.battleship-home { color:#fff; }
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
	.more-help[open] summary:before { content:'➖'; }
</style>