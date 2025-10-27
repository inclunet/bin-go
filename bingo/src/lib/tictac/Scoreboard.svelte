<script>
	export let scoreX = 0;
	export let scoreO = 0;
	export let scoreDraw = 0;

	$: scoreboardLine = `Placar: X ${scoreX}, O ${scoreO}${scoreDraw > 0 ? `, empates ${scoreDraw}` : ''}`;
</script>

<div class="scoreboard" role="status" aria-live="polite" aria-atomic="true">
	<span class="score-label">Placar:</span>
	<span class="score-x">X {scoreX}</span>
	<span class="score-divider">•</span>
	<span class="score-o">O {scoreO}</span>
	{#if scoreDraw > 0}
		<span class="score-divider">•</span>
		<span class="score-draw">Empates {scoreDraw}</span>
	{/if}
</div>

<style>
	.scoreboard {
		font-size: 1.1rem;
		font-weight: 600;
		background: linear-gradient(145deg, #1a2d42, #0f1b2a);
		padding: 0.75rem 1rem;
		border: 2px solid #2c4d6b;
		border-radius: 0.8rem;
		box-shadow: 
			0 0 0 2px rgba(43, 127, 244, 0.2),
			0 4px 12px rgba(0, 0, 0, 0.3);
		text-align: center;
		word-wrap: break-word;
		display: flex;
		align-items: center;
		justify-content: center;
		gap: 0.5rem;
		flex-wrap: wrap;
		color: #e8f4fd;
		min-height: 2.5rem;
	}

	.score-label {
		color: #a8c5e0;
		font-weight: 500;
	}

	.score-x {
		color: #2aa9ff;
		font-weight: 800;
		font-size: 1.2em;
		text-shadow: 0 0 8px rgba(42, 169, 255, 0.5);
		background: rgba(42, 169, 255, 0.1);
		padding: 0.2rem 0.4rem;
		border-radius: 0.3rem;
		border: 1px solid rgba(42, 169, 255, 0.3);
	}

	.score-o {
		color: #ff8b54;
		font-weight: 800;
		font-size: 1.2em;
		text-shadow: 0 0 8px rgba(255, 139, 84, 0.5);
		background: rgba(255, 139, 84, 0.1);
		padding: 0.2rem 0.4rem;
		border-radius: 0.3rem;
		border: 1px solid rgba(255, 139, 84, 0.3);
	}

	.score-draw {
		color: #ffc107;
		font-weight: 600;
	}

	.score-divider {
		color: #6c8ba5;
		font-size: 0.8rem;
	}

	@media (max-width: 480px) {
		.scoreboard {
			font-size: 1rem;
			padding: 0.6rem 0.8rem;
			max-width: 100%;
			gap: 0.4rem;
			min-height: 2.2rem;
			border-width: 1px;
		}
		
		.score-label {
			font-size: 0.9rem;
		}
	}

	@media (max-width: 360px) {
		.scoreboard {
			font-size: 0.95rem;
			padding: 0.5rem 0.7rem;
			gap: 0.3rem;
			flex-direction: column;
			min-height: auto;
		}
		
		.score-label {
			font-size: 0.85rem;
			margin-bottom: 0.2rem;
		}
		
		.score-divider {
			display: none;
		}
	}
</style>