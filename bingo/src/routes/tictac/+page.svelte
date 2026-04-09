<script>
	import { onMount, onDestroy } from "svelte";
	import PageTitle from "$lib/PageTitle.svelte";
	import Button from "$lib/Button.svelte";
	let creating = false;
	let lastRound = 0;
	let loading = true;
	let errorMsg = "";
	/** @typedef {{round:number, hasPlayerX:boolean, hasPlayerO:boolean}} OpenRound */
	/** @type {OpenRound[]} */
	let openRounds = [];
	let scoreX = 0;
	let scoreO = 0;
	let scoreDraw = 0;
	let lastWinner = "";
	/** @type {WebSocket | null} */
	let openWs = null;
	$: scoreSummary =
		lastRound > 0
			? `Placar acumulado até a rodada ${lastRound}: ${scoreX} vitória${scoreX === 1 ? "" : "s"} de X, ${scoreO} vitória${scoreO === 1 ? "" : "s"} de O${scoreDraw > 0 ? `, ${scoreDraw} empate${scoreDraw === 1 ? "" : "s"}` : ""}.`
			: "";

	function sortOpen() {
		openRounds = [...openRounds].sort((a, b) => a.round - b.round);
	}

	async function fetchLatest() {
		loading = true;
		try {
			let r = 1;
			for (; r <= 50; r++) {
				const res = await fetch(`/api/tictac/${r}`);
				if (res.status === 404) break;
				if (!res.ok) break;
			}
			lastRound = r - 1;
			if (lastRound > 0) {
				try {
					const lastRes = await fetch(`/api/tictac/${lastRound}`);
					if (lastRes.ok) {
						const data = await lastRes.json();
						scoreX = data.scoreX || 0;
						scoreO = data.scoreO || 0;
						scoreDraw = data.scoreDraw || 0;
						lastWinner = data.winner || "";
					}
				} catch {}
			}
			const openRes = await fetch("/api/tictac/open");
			if (openRes.ok) {
				openRounds = await openRes.json();
				sortOpen();
			}
		} catch (e) {
			errorMsg = "Falha ao detectar rodadas.";
		} finally {
			loading = false;
		}
	}

	function connectOpenWs() {
		try {
			openWs = new WebSocket(
				`${location.protocol === "https:" ? "wss" : "ws"}://${location.host}/ws/tictac/open`,
			);
			openWs.onmessage = (ev) => {
				try {
					const data = JSON.parse(ev.data);
					if (Array.isArray(data.rounds)) {
						openRounds = data.rounds;
						sortOpen();
					}
				} catch {}
			};
			openWs.onclose = () => {
				openWs = null;
				setTimeout(connectOpenWs, 3000);
			};
		} catch {}
	}

	async function newRound() {
		creating = true;
		errorMsg = "";
		try {
			const next = lastRound + 1 || 1;
			const res = await fetch(`/api/tictac/${next}/new`);
			if (!res.ok) {
				errorMsg = "Erro criando rodada";
				creating = false;
				return;
			}
			window.location.href = `/tictac/${next}`;
		} catch (e) {
			errorMsg = "Falha de rede";
			creating = false;
		}
	}

	onMount(() => {
		fetchLatest();
		connectOpenWs();
	});
	onDestroy(() => {
		if (openWs) {
			openWs.close();
			openWs = null;
		}
	});
</script>

<PageTitle title="Início" game="Jogo da Velha" />

<div class="container">
	<h2>Jogo da Velha</h2>

	<!-- aplicar componente Button -->
	<Button on:click={newRound} disabled={creating} ariaBusy={creating}>
		{creating ? "Criando..." : "Nova partida"}
	</Button>

	<!-- Total de rodadas criadas -->
	{#if !loading && !errorMsg}
		<p role="status" aria-live="polite">
			Total de {lastRound}
			{lastRound === 1 ? "partida criada" : "partidas criadas"} nesta sessão.
		</p>
	{/if}

	{#if loading}
		<p>Carregando...</p>
	{:else if errorMsg}
		<p class="text-danger">{errorMsg}</p>
	{:else if openRounds.length === 0}
		<p>Nenhuma partida aguardando jogadores no momento.</p>
	{:else}
		<h3>Partidas aguardando jogadores</h3>
		<ul class="list-open-rounds">
			{#each openRounds as r}
				<li class="li-size">
					<!-- aplicar componente Button -->
					<a class="btn a-size" href={`/tictac/${r.round}`}>
						<!-- ajustar o tamanho do botão para aceitar tamanhos personalizados do bootstrap -->
						<Button color="--tertiary-button-color" size="sm">
							Partida {r.round}
						</Button>
					</a><span
						>{r.hasPlayerX && r.hasPlayerO
							? "pronta para começar"
							: r.hasPlayerX || r.hasPlayerO
								? "1 jogador aguardando"
								: "aguardando jogadores"}</span
					>
				</li>
			{/each}
		</ul>
	{/if}

	<h3>Como jogar o jogo da velha?</h3>
	<section>
		<h4>Iniciando uma partida</h4>
		<ol>
			<li class="li-size">
				Clique em <strong>Nova partida</strong> para criar uma partida.
			</li>
			<li class="li-size">
				Na página da partida escolha seu símbolo: X ou O.
			</li>
			<li class="li-size">
				Compartilhe o link com outra pessoa usando o botão de
				compartilhar.
			</li>
			<li class="li-size">
				O segundo jogador entra no link e escolhe o símbolo restante.
			</li>
			<li class="li-size">
				O jogo da velha inicia automaticamente quando ambos jogadores
				estão presentes.
			</li>
			<li class="li-size">
				Ao terminar (vitória ou empate), use <strong
					>Nova partida</strong
				> para continuar jogando.
			</li>
		</ol>
	</section>
	<section>
		<h4>Partidas abertas</h4>
		<p>
			Uma partida aparece na lista acima enquanto ainda há vagas para
			jogadores. Você pode:
		</p>
		<ul>
			<li class="li-size">
				Entrar diretamente e escolher um símbolo disponível (X ou O).
			</li>
			<li class="li-size">
				Copiar o link da partida e enviar para um amigo.
			</li>
			<li class="li-size">Ver quantos jogadores já estão na partida.</li>
			<li class="li-size">
				Aguardar até que os dois jogadores estejam presentes para
				começar.
			</li>
		</ul>
	</section>
	<section>
		<h4>Controles e acessibilidade</h4>
		<ul>
			<li class="li-size">
				<strong>Navegação:</strong> Use as setas do teclado para mover entre
				as casas do tabuleiro.
			</li>
			<li class="li-size">
				<strong>Marcar jogada:</strong> Pressione Enter ou Espaço, ou use
				os números do teclado numérico (7-9: linha superior, 4-6: meio, 1-3:
				inferior).
			</li>
			<li class="li-size">
				<strong>Atalhos:</strong> Pressione ? para abrir ajuda; Esc para
				fechar diálogos.
			</li>
			<li class="li-size">
				<strong>Leitores de tela:</strong> Recebem anúncios automáticos sobre
				turnos, jogadas e resultados.
			</li>
		</ul>
	</section>
	<section>
		<h4>Sistema de placar</h4>
		<p>
			O placar é mantido ao longo de uma sequência de partidas. Cada vez
			que você clica em "Nova partida" após terminar um jogo, o placar
			continua acumulando. O placar mostra quantas vitórias cada símbolo
			(X e O) conquistou, além dos empates.
		</p>
	</section>
	<!-- <details class="more-help"> -->
	<details>
		<summary class="summary">Dicas avançadas</summary>
		<ul class="mt-2">
			<li class="li-size">
				<strong>Primeiro turno:</strong> O primeiro jogador que entra numa
				nova sequência determina quem começa jogando com X.
			</li>
			<li class="li-size">
				<strong>Próximas partidas:</strong> O vencedor sempre inicia a próxima
				partida; em caso de empate, alterna quem começa.
			</li>
			<li class="li-size">
				<strong>Reconexão:</strong> Se recarregar a página, você retoma automaticamente
				a mesma partida pelo número da URL.
			</li>
			<li class="li-size">
				<strong>Teclado numérico:</strong> A disposição dos números segue
				o layout físico do teclado numérico para facilitar o uso.
			</li>
		</ul>
	</details>
</div>

<style>
	/* .tic-home { color:#fff; }
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
	.guide-grid ul { padding-left:1.1rem; }*/
	/* .more-help summary {
		cursor: pointer;
	}
	.more-help summary::-webkit-details-marker {
		display: none;
	}
	.more-help summary {
		position: relative;
		padding-left: 1.1rem;
	}
	.more-help summary:before {
		content: "➕";
		position: absolute;
		left: 0;
	}
	.more-help summary:before {
		content: "➕";
		position: absolute;
		left: 0;
	}
	.more-help[open] summary:before {
		content: "➖";
	} */
	.summary {
		font-size: 2.2rem;
		margin-top: 3rem;
		/* color: #1d4ed8; */
		color: var(--octonary-color);
	}
</style>
