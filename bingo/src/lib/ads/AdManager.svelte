<script>
	// @ts-nocheck
	import { onMount } from 'svelte';
	import { adConfig, shouldShowAd, getAdSize, getDeviceType } from './adConfig.js';
	
	export let placement = 'auto'; // 'top', 'bottom', 'sidebar', 'auto'
	export let format = 'auto'; // 'rectangle', 'banner', 'square', 'auto'
	export let responsive = true;
	export let adClient = adConfig.client; // Usar configuração padrão
	export let adSlot = ''; // ID do slot do anúncio
	export let priority = 'low'; // 'high', 'medium', 'low' - controla quando carregar
	
	let adContainer;
	let isVisible = false;
	let adLoaded = false;
	let deviceType = 'desktop';
	
	// Detectar tipo de dispositivo
	function updateDeviceType() {
		deviceType = getDeviceType();
	}
	
	function shouldDisplay() {
		// Verificar se o placement é permitido para este dispositivo
		if (!shouldShowAd(placement, deviceType)) {
			return false;
		}
		
		// Não mostrar em telas muito pequenas
		if (typeof window !== 'undefined' && window.innerWidth < 360) {
			return false;
		}
		
		// Verificar se temos cliente e slot
		return adClient && adSlot;
	}
	
	function getOptimalAdSize() {
		if (!shouldDisplay()) return null;
		
		// Usar configuração automática baseada no dispositivo
		return getAdSize(deviceType, format);
	}
	
	function initializeAd() {
		if (!shouldDisplay()) return;
		
		try {
			// Configurar AdSense responsivo
			const adElement = adContainer.querySelector('.adsbygoogle');
			if (adElement && window.adsbygoogle) {
				window.adsbygoogle.push({});
				adLoaded = true;
			}
		} catch (error) {
			console.warn('Erro ao carregar anúncio:', error);
		}
	}
	
	onMount(() => {
		updateDeviceType();
		
		// Listener para mudanças de tamanho
		const handleResize = () => {
			updateDeviceType();
		};
		window.addEventListener('resize', handleResize);
		
		// Observer para lazy loading
		const observer = new IntersectionObserver((entries) => {
			entries.forEach(entry => {
				if (entry.isIntersecting && !adLoaded) {
					isVisible = true;
					const delay = priority === 'low' ? adConfig.settings.loadDelay : 
								 priority === 'medium' ? 1000 : 0;
					
					setTimeout(initializeAd, delay);
				}
			});
		}, { threshold: 0.1 });
		
		if (adContainer) {
			observer.observe(adContainer);
		}
		
		return () => {
			observer.disconnect();
			window.removeEventListener('resize', handleResize);
		};
	});
	
	$: adSize = getOptimalAdSize();
	$: showAd = shouldDisplay();
	$: isMobile = deviceType === 'mobile';
</script>

{#if showAd && adSize}
	<div 
		bind:this={adContainer}
		class="ad-container ad-{placement}"
		class:ad-mobile={isMobile}
		style="max-width: {adSize.width}px; min-height: {adSize.height}px;"
	>
		<div class="ad-label">Publicidade</div>
		
		{#if isVisible || priority === 'high'}
			<ins class="adsbygoogle"
				style="display:block; width:{adSize.width}px; height:{adSize.height}px;"
				data-ad-client={adClient}
				data-ad-slot={adSlot}
				data-ad-format={responsive ? 'auto' : format}
				data-full-width-responsive={responsive ? 'true' : 'false'}
			></ins>
		{:else}
			<!-- Placeholder para evitar layout shift -->
			<div class="ad-placeholder" style="width:{adSize.width}px; height:{adSize.height}px;">
				<div class="ad-loading">Carregando anúncio...</div>
			</div>
		{/if}
	</div>
{/if}

<style>
	.ad-container {
		margin: 1rem auto;
		text-align: center;
		position: relative;
		background: rgba(255, 255, 255, 0.02);
		border: 1px solid rgba(255, 255, 255, 0.1);
		border-radius: 0.5rem;
		padding: 0.5rem;
		box-sizing: border-box;
	}
	
	.ad-label {
		font-size: 0.7rem;
		color: rgba(255, 255, 255, 0.4);
		text-transform: uppercase;
		letter-spacing: 0.05em;
		margin-bottom: 0.25rem;
		font-family: system-ui, sans-serif;
	}
	
	.ad-placeholder {
		background: linear-gradient(90deg, rgba(255,255,255,0.1) 0%, rgba(255,255,255,0.05) 50%, rgba(255,255,255,0.1) 100%);
		background-size: 200% 100%;
		animation: shimmer 2s infinite;
		border-radius: 0.25rem;
		display: flex;
		align-items: center;
		justify-content: center;
	}
	
	.ad-loading {
		color: rgba(255, 255, 255, 0.3);
		font-size: 0.8rem;
	}
	
	@keyframes shimmer {
		0% { background-position: -200% 0; }
		100% { background-position: 200% 0; }
	}
	
	/* Posicionamento específico */
	.ad-top {
		order: -1;
		margin-bottom: 1.5rem;
	}
	
	.ad-bottom {
		order: 99;
		margin-top: 1.5rem;
	}
	
	.ad-sidebar {
		position: fixed;
		right: 1rem;
		top: 50%;
		transform: translateY(-50%);
		margin: 0;
		z-index: 10;
	}
	
	/* Responsividade */
	@media (max-width: 768px) {
		.ad-container {
			margin: 0.75rem auto;
			max-width: 100% !important;
		}
		
		.ad-sidebar {
			position: static;
			transform: none;
			margin: 1rem auto;
		}
		
		.ad-mobile {
			padding: 0.25rem;
		}
	}
	
	@media (max-width: 480px) {
		.ad-container {
			margin: 0.5rem auto;
			border-radius: 0.25rem;
		}
	}
	
	/* Esconder em telas muito pequenas */
	@media (max-width: 360px) {
		.ad-container {
			display: none;
		}
	}
	
	/* Proteção contra sobreposição */
	.ad-container:not(.ad-sidebar) {
		clear: both;
		isolation: isolate;
	}
	
	/* Garantir que não interfira no tabuleiro */
	:global(.tic-board) + .ad-container,
	:global(.scoreboard-container) + .ad-container {
		margin-top: 2rem;
	}
</style>