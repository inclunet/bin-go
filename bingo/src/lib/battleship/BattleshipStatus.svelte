<script>
	// @ts-nocheck
	export let phase = 'setup'; // 'setup', 'playing', 'finished'
	export let currentPlayer = 'a';
	export let localPlayer = 'a';
	export let isMyTurn = false;
	export let message = '';
	export let shipsRemaining = { a: 5, b: 5 };
	export let playerAReady = false;
	export let playerBReady = false;
	export let lastAction = '';

	$: isSetupPhase = phase === 'setup';
	$: isMySetup = isSetupPhase && currentPlayer === localPlayer;
	$: currentPlayerName = currentPlayer === 'a' ? 'Jogador A' : 'Jogador B';
	$: localPlayerName = localPlayer === 'a' ? 'Jogador A' : 'Jogador B';

	function getStatusMessage() {
		// Mensagem externa (ex: erro) tem prioridade, exceto se estiver vazia
		if (message) return message;

		if (phase === 'setup') {
			const localReady = (localPlayer === 'a') ? playerAReady : playerBReady;
			const oppReady = (localPlayer === 'a') ? playerBReady : playerAReady;
			if (!localReady) return 'Posicione seus navios';
			if (localReady && !oppReady) return 'Aguardando oponente';
			return 'Preparando início';
		}

		if (phase === 'playing') {
			// Simplificar e evitar frase longa ou resquício do setup
			if (winner) return `Vitória de ${winner === 'a' ? 'Jogador A' : 'Jogador B'}`;
			return isMyTurn ? 'Seu turno' : `Turno do oponente`;
		}

		if (phase === 'finished') {
			if (winner) return `Fim: ${winner === 'a' ? 'Jogador A venceu' : winner === 'b' ? 'Jogador B venceu' : 'Empate'}`;
			return 'Partida finalizada';
		}

		return '';
	}
</script>

<div class="battleship-status">
	<!-- Região silenciosa visual -->
	<div class="status-main">
		<span class="status-text" aria-hidden="true">{getStatusMessage()}</span>
		{#if lastAction}
			<span class="last-action" aria-hidden="true">• {lastAction}</span>
		{/if}
	</div>
    
	{#if phase === 'playing'}
		<div class="ships-status" aria-hidden="true">
			<div class="ship-count">
				<span class="ship-label">🚢 Jogador A:</span>
				<span class="ship-number" class:low={shipsRemaining.a <= 2}>{shipsRemaining.a} navios</span>
			</div>
			<div class="ship-count">
				<span class="ship-label">⚓ Jogador B:</span>
				<span class="ship-number" class:low={shipsRemaining.b <= 2}>{shipsRemaining.b} navios</span>
			</div>
		</div>
	{/if}

	<!-- Região aria-live dedicada apenas a mudanças de fase/turno (sem repetir resultado de tiro) -->
	<div class="sr-only" aria-live="polite">{getStatusMessage()}</div>
</div>

<style>
	.battleship-status {
		background: linear-gradient(145deg, #1e3a52, #142536);
		padding: 0.8rem 1.2rem;
		border: 2px solid #3a5f91;
		border-radius: 0.75rem;
		box-shadow: 
			0 0 0 2px rgba(43, 127, 244, 0.2),
			0 4px 8px rgba(0, 0, 0, 0.3);
		color: #e8f4fd;
		text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
		display: flex;
		flex-direction: column;
		gap: 0.6rem;
		min-height: 2.8rem;
		justify-content: center;
	}
	
	.status-main {
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;
		flex-wrap: wrap;
		text-align: center;
	}
	
	.status-text {
		font-size: 1.2rem;
		font-weight: 600;
		color: #fff;
	}
	
	.last-action {
		font-size: 1rem;
		color: #a8c5e0;
		font-weight: 500;
		font-style: italic;
	}
	
	.ships-status {
		display: flex;
		justify-content: center;
		gap: 2rem;
		padding-top: 0.4rem;
		border-top: 1px solid rgba(255, 255, 255, 0.1);
	}
	
	.ship-count {
		display: flex;
		align-items: center;
		gap: 0.4rem;
		font-size: 0.95rem;
	}
	
	.ship-label {
		color: #a8c5e0;
		font-weight: 500;
	}
	
	.ship-number {
		font-weight: 700;
		color: #6bc96b;
		background: rgba(107, 201, 107, 0.15);
		padding: 0.2rem 0.5rem;
		border-radius: 0.3rem;
		border: 1px solid rgba(107, 201, 107, 0.3);
	}
	
	.ship-number.low {
		color: #ff6b6b;
		background: rgba(255, 107, 107, 0.15);
		border-color: rgba(255, 107, 107, 0.3);
		animation: warning-pulse 2s ease-in-out infinite;
	}
	
	@keyframes warning-pulse {
		0%, 100% { opacity: 1; }
		50% { opacity: 0.7; }
	}
	
	/* Responsividade */
	@media (max-width: 768px) {
		.battleship-status {
			padding: 0.7rem 1rem;
		}
		
		.status-text {
			font-size: 1.1rem;
		}
		
		.last-action {
			font-size: 0.9rem;
		}
		
		.ships-status {
			gap: 1.5rem;
		}
		
		.ship-count {
			font-size: 0.9rem;
		}
	}
	
	@media (max-width: 480px) {
		.battleship-status {
			padding: 0.6rem 0.8rem;
		}
		
		.status-text {
			font-size: 1rem;
		}
		
		.ships-status {
			flex-direction: column;
			gap: 0.5rem;
			align-items: center;
		}
		
		.ship-count {
			justify-content: center;
		}
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
</style>