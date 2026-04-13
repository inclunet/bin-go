<script lang="ts">
	import { page } from '$app/stores';
	import { get } from 'svelte/store';
	import { onMount } from 'svelte';
	import PageTitle from '$lib/PageTitle.svelte';

	const p = get(page);
	let roundParam = p.params.round;
	let loading = true;
	let errorMsg = '';
	let roundData = null;
	let hasPlayerA = false;
	let hasPlayerB = false;
	let shareMsg = '';
	let copyMsg = '';
	let isMobile = false;
	let shareLabel = 'Compartilhar link da partida';

	async function loadRound() {
		try {
			const res = await fetch(`/api/battleship/${roundParam}`);
			if (res.status === 404) {
				errorMsg = 'Partida não encontrada.';
				loading = false;
				return;
			}
			if (!res.ok) {
				errorMsg = 'Erro ao carregar partida.';
				loading = false;
				return;
			}
			roundData = await res.json();
			hasPlayerA = !!roundData.playerA;
			hasPlayerB = !!roundData.playerB;
		} catch (e) {
			errorMsg = 'Erro de conexão.';
			console.error('Erro ao carregar rodada:', e);
		} finally {
			loading = false;
		}
	}

	function selectPlayer(player: 'a' | 'b') {
		window.location.href = `/battleship/${roundParam}/${player}`;
	}

	async function shareInvite() {
		const url = `${location.origin}/battleship/${roundParam}`;
		try {
			if (navigator.share) {
				await navigator.share({
					title: 'Batalha Naval',
					text: 'Entre na partida de Batalha Naval inclusiva',
					url
				});
				shareMsg = 'Link compartilhado';
			} else {
				await navigator.clipboard.writeText(url);
				shareMsg = 'Link copiado';
			}
		} catch (e) {
			shareMsg = 'Não foi possível compartilhar';
		}
		setTimeout(() => shareMsg = '', 2500);
	}

	function copyLink() {
		const url = `${location.origin}/battleship/${roundParam}`;
		try {
			navigator.clipboard.writeText(url);
			copyMsg = 'Link copiado!';
		} catch {
			copyMsg = 'Falha ao copiar';
		}
		setTimeout(() => copyMsg = '', 2500);
	}

	function detectMobileShare() {
		if (typeof window === 'undefined') return false;
		const ua = navigator.userAgent || '';
		const touch = ('ontouchstart' in window) || (navigator.maxTouchPoints || 0) > 0;
		const coarse = window.matchMedia && window.matchMedia('(pointer:coarse)').matches;
		const uaMatch = /Mobi|Android|iPhone|iPad|iPod/i.test(ua);
		return !!(navigator.share) || (touch && (coarse || uaMatch));
	}

	onMount(() => {
		loadRound();
		isMobile = detectMobileShare();
	});
</script>

<PageTitle title="Selecionar Jogador" game="Batalha Naval" />

<div class="container py-4 d-flex flex-column align-items-center" style="max-width: 640px;">
	<h2 class="text-center mb-4">Partida {roundParam}</h2>

	{#if loading}
		<div class="text-center">
			<p>Carregando partida...</p>
		</div>
	{:else if errorMsg}
		<div class="text-center">
			<p class="text-danger">{errorMsg}</p>
			<a href="/battleship" class="btn btn-secondary mt-3">Voltar</a>
		</div>
	{:else}
		<div class="selection-container">
			<h3 class="text-center mb-4">Escolha seu jogador</h3>
			
			<div class="players-grid">
				<!-- Jogador A -->
				<div class="player-card {hasPlayerA ? 'occupied' : 'available'}">
					<div class="player-icon">🚢</div>
					<h4>Jogador A</h4>
					<p class="status">
						{hasPlayerA ? 'Ocupado' : 'Disponível'}
					</p>
					{#if !hasPlayerA}
						<button 
							class="btn btn-primary btn-lg"
							on:click={() => selectPlayer('a')}
							aria-label="Selecionar Jogador A"
						>
							Escolher
						</button>
					{:else}
						<button class="btn btn-secondary btn-lg" disabled>
							Ocupado
						</button>
					{/if}
				</div>

				<!-- Jogador B -->
				<div class="player-card {hasPlayerB ? 'occupied' : 'available'}">
					<div class="player-icon">⚓</div>
					<h4>Jogador B</h4>
					<p class="status">
						{hasPlayerB ? 'Ocupado' : 'Disponível'}
					</p>
					{#if !hasPlayerB}
						<button 
							class="btn btn-primary btn-lg"
							on:click={() => selectPlayer('b')}
							aria-label="Selecionar Jogador B"
						>
							Escolher
						</button>
					{:else}
						<button class="btn btn-secondary btn-lg" disabled>
							Ocupado
						</button>
					{/if}
				</div>
			</div>

			<div class="share-section mt-4">
				<div class="mb-3 d-flex flex-wrap gap-3 align-items-start qr-share-wrapper">
					<div class="d-flex flex-column gap-2 share-buttons" aria-label="Ações de convite">
						<button class="btn btn-outline-info" on:click={copyLink} aria-label="Copiar link da partida">Copiar link</button>
						{#if isMobile}
							<button class="btn btn-warning" on:click={shareInvite} aria-label={shareLabel}>Compartilhar</button>
						{/if}
						{#if copyMsg}<span class="small text-success" aria-live="polite">{copyMsg}</span>{/if}
						{#if shareMsg}<span class="small text-success" aria-live="polite">{shareMsg}</span>{/if}
					</div>
					<div class="qr-inline" aria-label="QR Code para convidar outro jogador">
						<h4 class="h6 text-center mb-2">Entrar na partida</h4>
						<img src={`/api/battleship/${roundParam}/qr`} alt="QR Code da partida {roundParam}" />
						<p class="small text-muted mt-2 mb-0 text-center">Escaneie para abrir /battleship/{roundParam}</p>
					</div>
				</div>
				<div class="text-center mt-2">
					<a href="/battleship" class="btn btn-outline-light">Voltar ao início</a>
				</div>
			</div>

			{#if hasPlayerA && hasPlayerB}
				<div class="alert alert-info mt-4">
					<strong>Partida completa!</strong> Ambos os jogadores estão presentes. 
					Entre como espectador ou aguarde uma nova partida.
				</div>
			{/if}
		</div>
	{/if}
</div>

<style>
	.selection-container {
		width: 100%;
		max-width: 560px;
	}

	.players-grid {
		display: grid;
		grid-template-columns: 1fr 1fr;
		gap: 2rem;
		margin-bottom: 2rem;
	}

	.player-card {
		background: linear-gradient(145deg, #1e3a52, #142536);
		border: 2px solid #3a5f91;
		border-radius: 1rem;
		padding: 2rem 1rem;
		text-align: center;
		transition: all 0.3s ease;
		position: relative;
	}

	.player-card.available {
		border-color: #28a745;
		box-shadow: 0 0 0 2px rgba(40, 167, 69, 0.2);
	}

	.player-card.available:hover {
		transform: translateY(-2px);
		box-shadow: 0 4px 12px rgba(40, 167, 69, 0.3);
	}

	.player-card.occupied {
		border-color: #6c757d;
		background: linear-gradient(145deg, #2a2a2a, #1a1a1a);
		opacity: 0.7;
	}

	.player-icon {
		font-size: 3rem;
		margin-bottom: 1rem;
		opacity: 0.8;
	}

	.player-card h4 {
		color: #e8f4fd;
		margin-bottom: 0.5rem;
		font-weight: 600;
	}

	.status {
		font-size: 0.9rem;
		margin-bottom: 1.5rem;
		color: #b4c9d6;
	}

	.player-card.available .status {
		color: #28a745;
		font-weight: 600;
	}

	.player-card.occupied .status {
		color: #6c757d;
	}

	.alert {
		background: rgba(23, 162, 184, 0.1);
		border: 1px solid rgba(23, 162, 184, 0.3);
		border-radius: 0.5rem;
		padding: 1rem;
		color: #bee5eb;
	}

	.qr-share-wrapper { align-items:stretch; }
	.qr-inline { background:#152635; border:1px solid #2c4d6b; padding:.9rem .9rem 1rem; border-radius:.75rem; width:170px; display:flex; flex-direction:column; }
	.qr-inline img { width:100%; height:auto; background:#fff; border-radius:.4rem; padding:.3rem; box-shadow:0 0 0 3px rgba(255,255,255,.05); }
	.share-buttons button { min-width:180px; }

	@media (max-width: 600px) {
		.qr-share-wrapper { flex-direction:column; align-items:flex-start; }
		.qr-inline { width:190px; }
	}

	@media (max-width: 480px) {
		.players-grid {
			grid-template-columns: 1fr;
			gap: 1.5rem;
		}
		
		.player-card {
			padding: 1.5rem 1rem;
		}
		
		.player-icon {
			font-size: 2.5rem;
		}
	}
</style>