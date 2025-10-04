<script>
	export let scoreX = 0;
	export let scoreO = 0;
	export let scoreDraw = 0;

	$: scoreboardLine = `Placar: X ${scoreX}, O ${scoreO}${scoreDraw > 0 ? `, empates ${scoreDraw}` : ''}`;
</script>

<div class="scoreboard" role="status" aria-live="polite" aria-atomic="true">
	{scoreboardLine}
</div>

<style>
	.scoreboard {
		font-size: 1.05rem;
		background: #142536;
		padding: 0.4rem 0.85rem;
		border: 1px solid #2c4d6b;
		border-radius: 0.6rem;
		box-shadow: 0 0 0 2px rgba(43, 127, 244, 0.15);
		text-align: center;
		word-wrap: break-word;
	}

	@media (max-width: 480px) {
		.scoreboard {
			font-size: 0.9rem;
			padding: 0.3rem 0.6rem;
			max-width: 100%;
		}
	}
</style>