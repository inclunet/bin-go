<script>
	// @ts-nocheck
	export let visibleInfo = '';
</script>

<div class="live-bar">{visibleInfo}</div>

<style>
	.live-bar {
		font-weight: 600;
		font-size: 1.25rem;
		background: #142536;
		padding: 0.55rem 0.95rem;
		border: 1px solid #2c4d6b;
		border-radius: 0.65rem;
		min-height: 2.4rem;
		display: flex;
		align-items: center;
	}

	@media (max-width: 480px) {
		.live-bar {
			font-size: 1.1rem;
			padding: 0.45rem 0.75rem;
		}
	}
</style>