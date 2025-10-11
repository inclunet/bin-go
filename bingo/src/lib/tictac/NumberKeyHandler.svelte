<script>
	// @ts-nocheck
	import { createEventDispatcher, onMount } from 'svelte';
	
	export let showHelpModal = false;
	export let winner = '';

	const dispatch = createEventDispatcher();

	function handleNumberKey(key) {
		if (['1','2','3','4','5','6','7','8','9'].includes(key)) {
			const num = parseInt(key);
			// Mapear números para corresponder ao layout do teclado numérico
			let row, col;
			if (num >= 7 && num <= 9) { // linha superior (7,8,9)
				row = 0; col = num - 7;
			} else if (num >= 4 && num <= 6) { // linha meio (4,5,6)
				row = 1; col = num - 4;
			} else { // linha inferior (1,2,3)
				row = 2; col = num - 1;
			}
			dispatch('numberKey', { row, col, key });
			return true;
		}
		return false;
	}

	function handleGlobalKeydown(e) {
		if (e.key === '?' || (e.shiftKey && e.key === '/')) {
			dispatch('helpToggle');
			e.preventDefault();
		}
		else if (e.key === 'Escape' && showHelpModal) {
			dispatch('helpClose');
			e.preventDefault();
		}
		// Navegação global por números 1-9 quando não há modal aberto
		else if (!showHelpModal && !winner && handleNumberKey(e.key)) {
			e.preventDefault();
		}
	}

	onMount(() => {
		window.addEventListener('keydown', handleGlobalKeydown);
		return () => window.removeEventListener('keydown', handleGlobalKeydown);
	});
</script>

<!-- Este componente não renderiza nada, apenas gerencia eventos de teclado -->