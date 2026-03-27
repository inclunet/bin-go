<script>
	// @ts-nocheck
	export let scoreA = 0;
	export let scoreB = 0;
	export let scoreDraw = 0;
	export let currentPlayer = 'a';
	export let phase = 'setup'; // 'setup', 'playing', 'finished'
	export let winner = '';

	$: scoreboardLine = `Placar: Jogador A ${scoreA}, Jogador B ${scoreB}${scoreDraw > 0 ? `, empates ${scoreDraw}` : ''}`;
	$: currentPlayerName = currentPlayer === 'a' ? 'Jogador A' : 'Jogador B';
</script>

<div class="battleship-scoreboard" role="status" aria-live="polite" aria-atomic="true">
	<div class="score-display">
		<span class="score-label">Placar:</span>
		<div class="scores">
			<span class="score-player score-a" class:active={currentPlayer === 'a'}>
				<span class="player-icon">🚢</span>
				<span class="player-label">Jogador A</span>
				<span class="score-value">{scoreA}</span>
			</span>
			<span class="score-divider">•</span>
			<span class="score-player score-b" class:active={currentPlayer === 'b'}>
				<span class="player-icon">⚓</span>
				<span class="player-label">Jogador B</span>
				<span class="score-value">{scoreB}</span>
			</span>
			{#if scoreDraw > 0}
				<span class="score-divider">•</span>
				<span class="score-draw">
					<span class="draw-label">Empates</span>
					<span class="score-value">{scoreDraw}</span>
				</span>
			{/if}
		</div>
	</div>
	
	{#if phase !== 'finished'}
		<div class="phase-indicator">
			{#if phase === 'setup'}
				<span class="phase-text setup">📋 Posicionando navios</span>
			{:else if phase === 'playing'}
				<span class="phase-text playing">
					🎯 Vez de {currentPlayerName}
				</span>
			{/if}
		</div>
	{:else if winner}
		<div class="winner-announcement">
			<span class="winner-text">
				🏆 {winner === 'draw' ? 'Empate!' : `${winner === 'a' ? 'Jogador A' : 'Jogador B'} venceu!`}
			</span>
		</div>
	{/if}
</div>

<style>
	.battleship-scoreboard {
		font-size: 1.1rem;
		font-weight: 600;
		background: linear-gradient(145deg, #1a2d42, #0f1b2a);
		padding: 1rem 1.25rem;
		border: 2px solid #2c4d6b;
		border-radius: 0.8rem;
		box-shadow: 
			0 0 0 2px rgba(43, 127, 244, 0.2),
			0 4px 12px rgba(0, 0, 0, 0.3);
		color: #e8f4fd;
		min-height: 3rem;
		display: flex;
		flex-direction: column;
		gap: 0.75rem;
	}
	
	.score-display {
		display: flex;
		flex-direction: column;
		gap: 0.5rem;
		align-items: center;
	}
	
	.score-label {
		color: #a8c5e0;
		font-weight: 500;
		font-size: 0.95rem;
	}
	
	.scores {
		display: flex;
		align-items: center;
		gap: 0.75rem;
		flex-wrap: wrap;
		justify-content: center;
	}
	
	.score-player {
		display: flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.4rem 0.8rem;
		border-radius: 0.5rem;
		transition: all 0.3s ease;
		background: rgba(255, 255, 255, 0.05);
		border: 1px solid transparent;
	}
	
	.score-player.active {
		background: rgba(255, 193, 7, 0.15);
		border-color: rgba(255, 193, 7, 0.3);
		box-shadow: 0 0 8px rgba(255, 193, 7, 0.2);
	}
	
	.score-a {
		color: #4fa8da;
	}
	
	.score-a.active {
		color: #6bc0ff;
	}
	
	.score-b {
		color: #ff8b54;
	}
	
	.score-b.active {
		color: #ffab7a;
	}
	
	.player-icon {
		font-size: 1.2rem;
	}
	
	.player-label {
		font-weight: 600;
		font-size: 0.95rem;
	}
	
	.score-value {
		font-weight: 800;
		font-size: 1.1rem;
		color: #fff;
		background: rgba(255, 255, 255, 0.1);
		padding: 0.2rem 0.5rem;
		border-radius: 0.3rem;
		min-width: 1.5rem;
		text-align: center;
	}
	
	.score-draw {
		display: flex;
		align-items: center;
		gap: 0.4rem;
		color: #ffc107;
		font-weight: 600;
	}
	
	.draw-label {
		font-size: 0.9rem;
	}
	
	.score-divider {
		color: #6c8ba5;
		font-size: 1rem;
		opacity: 0.7;
	}
	
	.phase-indicator {
		text-align: center;
	}
	
	.phase-text {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.4rem 0.8rem;
		border-radius: 0.5rem;
		font-size: 0.95rem;
		font-weight: 600;
	}
	
	.phase-text.setup {
		background: rgba(40, 167, 69, 0.15);
		color: #6bc96b;
		border: 1px solid rgba(40, 167, 69, 0.3);
	}
	
	.phase-text.playing {
		background: rgba(255, 193, 7, 0.15);
		color: #ffd93d;
		border: 1px solid rgba(255, 193, 7, 0.3);
	}
	
	.winner-announcement {
		text-align: center;
	}
	
	.winner-text {
		display: inline-flex;
		align-items: center;
		gap: 0.5rem;
		padding: 0.5rem 1rem;
		background: linear-gradient(145deg, rgba(40, 167, 69, 0.2), rgba(28, 144, 56, 0.15));
		color: #6bc96b;
		border: 2px solid rgba(40, 167, 69, 0.4);
		border-radius: 0.6rem;
		font-size: 1.1rem;
		font-weight: 700;
		box-shadow: 0 0 12px rgba(40, 167, 69, 0.3);
		animation: victory-glow 2s ease-in-out infinite alternate;
	}
	
	@keyframes victory-glow {
		0% { box-shadow: 0 0 12px rgba(40, 167, 69, 0.3); }
		100% { box-shadow: 0 0 20px rgba(40, 167, 69, 0.5); }
	}
	
	/* Responsividade */
	@media (max-width: 768px) {
		.battleship-scoreboard {
			font-size: 1rem;
			padding: 0.8rem 1rem;
		}
		
		.scores {
			gap: 0.5rem;
		}
		
		.score-player {
			padding: 0.3rem 0.6rem;
			gap: 0.4rem;
		}
		
		.player-label {
			font-size: 0.85rem;
		}
		
		.score-value {
			font-size: 1rem;
			padding: 0.1rem 0.4rem;
		}
	}
	
	@media (max-width: 480px) {
		.battleship-scoreboard {
			padding: 0.6rem 0.8rem;
		}
		
		.scores {
			flex-direction: column;
			gap: 0.4rem;
		}
		
		.score-divider {
			display: none;
		}
		
		.score-player {
			justify-content: center;
			min-width: 120px;
		}
		
		.phase-text, .winner-text {
			font-size: 0.9rem;
		}
	}
</style>