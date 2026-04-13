<script>
	export let color = "--primary-button-color";
	export let data_bs_toggle = "";
	export let data_bs_target = "";
	export let data_bs_dismiss = "";
	export let disabled = false;
	export let ariaBusy = false;
	export let aria_label = "";
	export let size = "nm";
</script>

<button
	class="btn"
	data-size={size}
	style={`--btn-color: var(${color}, var(--primary-button-color));`}
	{disabled}
	aria-busy={ariaBusy}
	data-bs-toggle={data_bs_toggle}
	data-bs-target={data_bs_target}
	data-bs-dismiss={data_bs_dismiss}
	aria-label={aria_label}
	on:click
>
	<strong><slot></slot></strong>
</button>

<style lang="scss">
	@mixin hover-darken($black-percent: 10%) {
		background-color: var(--btn-color);
		color: var(--black);

		@media (hover: hover) {
			&:hover {
				background-color: color-mix(
					in srgb,
					var(--btn-color) #{100% - $black-percent},
					black $black-percent
				);
				color: var(--white);
			}
		}
	}

	button {
		margin-top: 0;
		color: var(--black);
		border: none;
		border-radius: 4px;
		cursor: pointer;

		@include hover-darken(15%);
	}

	button[data-size="lg"] {
		padding: 1.4rem;
		font-size: 3rem;
	}

	button[data-size="nm"] {
		padding: 1rem;
		font-size: 2rem;
	}

	button[data-size="sm"] {
		padding: 0.1rem 0.5rem;
		font-size: 1.8rem;
	}

	@media (max-width: 450px) {
		button {
			height: 4.8rem;
			padding: 0.5rem;
			font-size: 1.5rem;
		}
	}

	@media (max-width: 410px) {
		button {
			height: 4.8rem;
			padding: 0.5rem 0.3rem;
			font-size: 1.4rem;
		}
	}
</style>
