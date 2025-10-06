<script>
	// @ts-nocheck
	export let visibleInfo = '';
</script>

<div class="live-bar">{visibleInfo}</div>

<style>
	.live-bar {
		font-weight: 600;
		font-size: 1.3rem;
		/* Gradiente mais visível */
		background: linear-gradient(145deg, #1e3a52, #142536);
		padding: 0.7rem 1.1rem;
		/* Borda mais destacada */
		border: 2px solid #3a5f91;
		border-radius: 0.75rem;
		min-height: 2.8rem;
		display: flex;
		align-items: center;
		justify-content: center;
		text-align: center;
		/* Sombra para destaque */
		box-shadow: 
			0 0 0 2px rgba(43, 127, 244, 0.2),
			0 4px 8px rgba(0, 0, 0, 0.3);
		/* Cor do texto mais brilhante */
		color: #e8f4fd;
		text-shadow: 0 1px 2px rgba(0, 0, 0, 0.5);
	}

	@media (max-width: 480px) {
		.live-bar {
			font-size: 1.15rem;
			padding: 0.6rem 0.9rem;
			min-height: 2.5rem;
		}
	}
	
	@media (max-width: 360px) {
		.live-bar {
			font-size: 1.05rem;
			padding: 0.5rem 0.7rem;
		}
	}
</style>