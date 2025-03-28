<script>
	import Tooltip from "./Tooltip.svelte";
	import BrailleDot from "./BrailleDot.svelte";
	import BrailleWord from "./BrailleWord.svelte";
	import { onMount, onDestroy } from "svelte";

	export let brailleWord = "";

	let brailleCell = 0;
	let enableSpaceTip = false;
	let spacebar;
	let isClient = typeof window !== "undefined";
	let buttonLimpar;
	let buttonBackSpace;

	const handleBrailleKey = (event = { detail: { key: 0 } }) => {
		if (event.detail.key == 0) {
			brailleCell = 0;
		}

		if (event.detail.key > 0 && event.detail.key < 9) {
			brailleCell ^= 1 << (event.detail.key - 1);
		}
	};

	const handleBrailleTypping = (event = { detail: "" }) => {
		if (brailleCell > 0 || event.detail == "space") {
			brailleWord += String.fromCharCode(0x2800 + brailleCell);
		}
		brailleCell = 0;
	};

	const handleSpaceKey = () => {
		handleBrailleTypping({ detail: "space" });
	};

	const handleBackspaceKey = () => {
		brailleWord = brailleWord.slice(0, -1);
	};

	const handleClearKey = () => {
		brailleWord = "";
	};

	const handleEnableSpaceTip = (
		/** @type {String} */ nameSound,
		/** @type {String} */ preFix
	) => {
		if (brailleCell > 0) {
			spacebar.focus();
			enableSpaceTip = true;
			handleInsertSound(nameSound, preFix);
		}
	};

	const handleDisableSpaceTip = () => {
		enableSpaceTip = false;
	};

	const handleInsertSound = (
		/** @type {String} */ nameSound,
		/** @type {String} */ type
	) => {
		const audio = new Audio(`/${nameSound}.${type}`);
		audio.play();
	};

	const handleHotKeys = (event) => {
		if (event.ctrlKey && event.altKey && event.key.toLowerCase() === "e") {
			event.preventDefault();
			handleSpaceKey();
			handleInsertSound("markLetter", "mp3");
		}

		if (event.ctrlKey && event.altKey && event.key.toLowerCase() === "l") {
			event.preventDefault();
			handleInsertSound("cleanLetter", "mp3");
			buttonLimpar.click();
		}

		if (event.ctrlKey && event.altKey && event.key.toLowerCase() === "b") {
			event.preventDefault();
			handleInsertSound("backLetter", "mp3");
			buttonBackSpace.click();
		}
	};

	onMount(() => {
		if (isClient) {
			document.addEventListener("keydown", handleHotKeys);
		}
	});

	onDestroy(() => {
		if (isClient) {
			document.removeEventListener("keydown", handleHotKeys);
		}
	});
</script>

<div role="region" aria-label="Resposta" class="container">
	<div>
		<BrailleWord
			on:brailleKey={handleBrailleKey}
			on:brailleEnter={handleBrailleTypping}
			on:backspacePress={handleBackspaceKey}
			on:backspacePress={() => handleInsertSound("backLetter", "mp3")}
			on:submitChallenge
			on:enableSpaceTip={() => handleEnableSpaceTip("tooltipShow", "mp3")}
			bind:brailleWord
			bind:brailleCell
			brailleKeyboard={true}
		/>
	</div>
	<div role="region" aria-label="Teclado Braille" class="container_keyboard">
		<div class="container_numbers">
			<div class="brailleDot_numbers brailleDot_left-numbers">
				<div class="um-quatro">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={1}
					/>
				</div>
				<div class="dois-cinco">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={2}
					/>
				</div>
				<div class="tres-seis">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={3}
					/>
				</div>
				<div class="sete-oito">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={7}
					/>
				</div>
			</div>

			<div class="brailleDot_numbers brailleDot_right-numbers">
				<div class="um-quatro">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={4}
					/>
				</div>
				<div class="dois-cinco">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={5}
					/>
				</div>
				<div class="tres-seis">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={6}
					/>
				</div>
				<div class="sete-oito">
					<BrailleDot
						on:brailleKey={handleBrailleKey}
						bind:brailleCell
						brailleDot={8}
					/>
				</div>
			</div>
		</div>
		<div class="container_buttons">
			<button
				on:click={handleClearKey}
				on:click={() => handleInsertSound("cleanLetter", "mp3")}
				bind:this={buttonLimpar}
				class="btn"
				id="limpar">Limpar</button
			>

			<div class="container_button-tooltip">
				<button
					bind:this={spacebar}
					on:click={handleSpaceKey}
					on:click={handleDisableSpaceTip}
					on:click={() => handleInsertSound("markLetter", "mp3")}
					class="btn"
					id="espaco">Espaço</button
				>
				<Tooltip {enableSpaceTip} marginTop="9rem">
					<!-- <p class="texto_tip">
                    </p> -->
					<p class="texto_tip">
						Ei, não se esqueça! <br />
						Pressione espaço para confirmar a letra que você quer enviar.
					</p>
				</Tooltip>
			</div>

			<button
				bind:this={buttonBackSpace}
				on:click={handleBackspaceKey}
				on:click={() => handleInsertSound("backLetter", "mp3")}
				class="btn"
				id="backspace">Backspace</button
			>
		</div>
	</div>
</div>

<style>
	.container {
		display: flex;
		flex-direction: column;
	}
	.container_keyboard {
		width: 100%;
		display: flex;
		flex-direction: column;
		align-items: center;
		justify-content: center;
		margin-top: 1em;
		margin-bottom: 2em;
	}

	.brailleDot_right-numbers {
		display: flex;
	}
	.brailleDot_left-numbers {
		display: flex;
		flex-direction: row-reverse;
	}
	.btn {
		color: var(--black);
		padding: 1rem 5rem;
		border: 1px solid var(--black);
		border-radius: 10rem;
		font-weight: 500;
	}
	.sete-oito {
		display: block;
		margin-top: 4rem;
	}
	.tres-seis {
		display: block;
		margin-top: 2rem;
	}
	.dois-cinco {
		display: block;
		margin-top: 0.5rem;
	}
	.um-quatro {
		display: block;
	}
	#limpar {
		background-color: var(--primary-button-limpar-color);
	}

	#limpar:active {
		background-color: var(--secondary-button-limpar-color);
		color: var(--white);
	}
	#espaco {
		background-color: var(--primary-button-espaco-color);
		padding-left: 15rem;
		padding-right: 15rem;
	}

	#espaco:active {
		background-color: var(--secondary-button-espaco-color);
		color: var(--white);
	}
	.texto_tip {
		margin: 0;
		padding: 0;
		font-size: 1.4rem;
		line-height: 3rem;
	}
	#backspace {
		background-color: var(--primary-button-backspace-color);
	}

	#backspace:active {
		background-color: var(--secondary-button-backspace-color);
		color: var(--white);
	}

	button {
		font-size: 1.8rem;
	}

	.container_numbers {
		display: flex;
		gap: 5rem;
	}
	.container_buttons {
		width: 100%;
		margin-top: 2rem;
		display: flex;
		justify-content: space-around;
	}
	.container_button-tooltip {
		display: flex;
		flex-direction: column;
		align-items: center;
		font-size: 1.2rem;
	}

	@media (hover: hover) {
		#limpar:hover {
			background-color: var(--secondary-button-limpar-color);
			color: var(--white);
		}
		#espaco:hover {
			background-color: var(--secondary-button-espaco-color);
			color: var(--white);
		}
		#backspace:hover {
			background-color: var(--secondary-button-backspace-color);
			color: var(--white);
		}
	}
	@media (max-width: 767px) {
		.container_buttons {
			justify-content: space-between;
		}
		.btn {
			width: 15rem;
			padding: 1rem 0;
		}
		#espaco {
			padding-left: 5rem;
			padding-right: 5rem;
		}
	}

	@media (max-width: 547px) {
		.container_numbers {
			gap: 2rem;
		}
	}

	@media (max-width: 518px) {
		button {
			margin: 0;
		}
		.container_numbers {
			gap: 0.5rem;
		}
	}

	@media (max-width: 490px) {
		.container_numbers {
			background-color: var(--back-keyboard-color);
			border-radius: 8rem;
			padding: 2rem;
		}
		.brailleDot_numbers {
			flex-direction: column;
		}

		.sete-oito {
			margin-top: 1rem;
		}
		.tres-seis {
			margin-top: 1rem;
		}
		.dois-cinco {
			margin-top: 1rem;
		}
		.btn {
			width: 12rem;
			font-size: 2rem;
			padding: 0 1rem;
			font-weight: 500;
		}
		#espaco {
			padding: 1rem 0;
		}

		#backspace {
			padding: 0 1rem !important;
		}

		.container_buttons {
			width: 90vw;
		}
	}
	@media (max-width: 424px) {
		.btn {
			padding: 0 0.5rem;
		}
		#limpar {
			width: 9rem;
		}
		#espaco {
			width: 10rem;
		}
	}
</style>
