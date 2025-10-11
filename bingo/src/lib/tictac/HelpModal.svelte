<script>
	// @ts-nocheck
	import { createEventDispatcher } from 'svelte';
	
	export let showModal = false;
	export let showQuickHelp = false;

	const dispatch = createEventDispatcher();

	function closeModal() {
		dispatch('close');
	}

	function toggleQuickHelp() {
		dispatch('toggleQuick');
	}
</script>

<button 
	class="btn btn-sm btn-outline-light help-toggle" 
	on:click={toggleQuickHelp} 
	aria-expanded={showQuickHelp} 
	aria-controls="help-panel"
>
	{showQuickHelp ? 'Ocultar ajuda' : 'Ajuda'}
</button>

{#if showQuickHelp}
	<div id="help-panel" class="help-box" role="note">
		<p><strong>Jogo da velha:</strong> Use as setas para navegar entre as casas. Pressione Enter/Espaço para marcar, ou use números do teclado numérico (7-9: linha superior, 4-6: meio, 1-3: inferior). Objetivo: fazer uma linha de 3 símbolos iguais.</p>
	</div>
{/if}

{#if showModal}
	<div id="help-modal" class="help-modal" role="dialog" aria-modal="true" aria-labelledby="help-title">
		<div class="help-dialog p-3">
			<h2 id="help-title" class="h5 mb-3">Como jogar o jogo da velha</h2>
			<ul class="small mb-3">
				<li><strong>Objetivo:</strong> Seja o primeiro a fazer uma linha de 3 símbolos iguais (horizontal, vertical ou diagonal).</li>
				<li><strong>Navegação:</strong> Use as setas do teclado para mover entre as casas do tabuleiro.</li>
				<li><strong>Marcar jogada:</strong> Pressione Enter ou Espaço na casa desejada.</li>
				<li><strong>Atalho numérico:</strong> Use os números do teclado numérico para marcar diretamente - 7,8,9: linha superior; 4,5,6: linha do meio; 1,2,3: linha inferior.</li>
				<li><strong>Navegação rápida:</strong> Home vai para o canto superior esquerdo, End para o inferior direito.</li>
				<li><strong>Sua vez:</strong> Aguarde o anúncio "É a sua vez" para jogar.</li>
				<li><strong>Continuar jogando:</strong> Use o botão "Nova partida" após cada jogo.</li>
				<li><strong>Atalhos:</strong> Pressione ? para esta ajuda; Esc para fechar diálogos.</li>
			</ul>
			<button class="btn btn-sm btn-secondary" on:click={closeModal}>Fechar</button>
		</div>
	</div>
{/if}

<style>
	.help-toggle {
		margin: 0 0 0.5rem auto;
		display: block;
	}

	.help-box {
		margin-top: 0.5rem;
		background: #152635;
		border: 1px solid #2c5278;
		padding: 0.75rem 0.9rem;
		border-radius: 0.6rem;
		font-size: 0.95rem;
		line-height: 1.35rem;
		color: #d5e4f1;
	}

	.help-box strong {
		color: #fff;
	}

	.help-modal {
		position: fixed;
		inset: 0;
		background: rgba(0,0,0,0.65);
		display: flex;
		align-items: center;
		justify-content: center;
		z-index: 1000;
	}

	.help-dialog {
		background: #132433;
		border: 1px solid #2c5278;
		border-radius: 0.8rem;
		width: 100%;
		max-width: 32rem;
		box-shadow: 0 10px 28px -4px rgba(0,0,0,0.6);
	}

	.help-dialog ul {
		padding-left: 1.1rem;
	}

	.help-dialog li {
		margin: 0.3rem 0;
		line-height: 1.25rem;
	}
</style>