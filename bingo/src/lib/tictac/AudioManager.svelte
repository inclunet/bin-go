<script>
	// @ts-nocheck
	import { onMount } from 'svelte';
	
	export let lastOutcome = '';

	let clickAudio = null;
	let errorAudio = null;
	let victoryAudio = null;
	let drawAudio = null;
	let defeatAudio = null;

	export function playClick() {
		clickAudio?.play();
	}

	export function playError() {
		errorAudio?.play();
	}

	export function playOutcomeSound(winner, localPlayer) {
		if (!winner || winner === lastOutcome) return;
		
		if (winner === 'Empate') {
			drawAudio?.play();
		} else if (winner === localPlayer) {
			victoryAudio?.play();
		} else {
			defeatAudio?.play();
		}
	}
</script>

<!-- Áudio: clique, erro, vitória, empate, derrota -->
<audio bind:this={clickAudio} src="/key.mp3" preload="auto"></audio>
<audio bind:this={errorAudio} src="/wrong.wav" preload="auto"></audio>
<audio bind:this={victoryAudio} src="/correct.wav" preload="auto"></audio>
<audio bind:this={drawAudio} src="/mixkit-cooking-bell-ding-1791.wav" preload="auto"></audio>
<audio bind:this={defeatAudio} src="/wrong.wav" preload="auto"></audio>