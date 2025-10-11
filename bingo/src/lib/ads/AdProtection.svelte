<script>
	// @ts-nocheck
	import { onMount } from 'svelte';
	
	export let protection = 'strict'; // 'strict', 'medium', 'light'
	export const areas = ['game-board', 'scoreboard', 'controls']; // áreas a proteger (referência externa)
	
	const protectionRules = {
		strict: {
			// Bloqueio total em áreas de jogo
			selectors: [
				'.tic-board', '.tic-cell', '.scoreboard-container', 
				'.game-controls', '.help-modal', '.status-bar',
				'[role="button"]', '[tabindex]', 'button'
			],
			minDistance: 100, // pixels mínimos de distância
			blockAutoAds: true
		},
		medium: {
			selectors: [
				'.tic-board', '.tic-cell', '.scoreboard-container'
			],
			minDistance: 50,
			blockAutoAds: false
		},
		light: {
			selectors: ['.tic-board'],
			minDistance: 20,
			blockAutoAds: false
		}
	};
	
	function addProtectionCSS() {
		if (typeof document === 'undefined') return;
		
		const rules = protectionRules[protection] || protectionRules.medium;
		const styleId = 'ad-protection-styles';
		
		// Remove estilo anterior se existir
		const existingStyle = document.getElementById(styleId);
		if (existingStyle) {
			existingStyle.remove();
		}
		
		const style = document.createElement('style');
		style.id = styleId;
		
		let css = `
			/* Proteção contra anúncios automáticos */
			.adsbygoogle {
				/* Evitar sobreposição com elementos do jogo */
				z-index: 1 !important;
				position: relative !important;
			}
			
			/* Criar zonas de exclusão */
			${rules.selectors.map(selector => `
				${selector} {
					position: relative;
					z-index: 100;
				}
				
				${selector}::before {
					content: '';
					position: absolute;
					top: -${rules.minDistance}px;
					left: -${rules.minDistance}px;
					right: -${rules.minDistance}px;
					bottom: -${rules.minDistance}px;
					pointer-events: none;
					z-index: 99;
				}
			`).join('')}
			
			/* Específico para o tabuleiro */
			.tic-board {
				isolation: isolate;
				contain: layout style;
			}
			
			.tic-cell {
				isolation: isolate;
				position: relative;
				z-index: 101;
			}
			
			/* Proteção do placar */
			.scoreboard-container {
				isolation: isolate;
				contain: layout;
			}
			
			/* Evitar anúncios em overlays */
			.help-modal,
			[role="dialog"],
			[aria-modal="true"] {
				z-index: 9999 !important;
				isolation: isolate;
			}
		`;
		
		// Bloquear anúncios automáticos se solicitado
		if (rules.blockAutoAds) {
			css += `
				/* Desabilitar anúncios automáticos em áreas protegidas */
				${rules.selectors.map(selector => `
					${selector},
					${selector} * {
						/* AdSense auto ads detection blocker */
						data-ad-client: none !important;
						data-ad-slot: none !important;
					}
				`).join('')}
				
				/* Criar barreira visual */
				.game-area {
					isolation: isolate;
					contain: layout style;
					background: var(--game-bg, transparent);
				}
			`;
		}
		
		style.textContent = css;
		document.head.appendChild(style);
	}
	
	function addMetaTags() {
		if (typeof document === 'undefined') return;
		
		// Meta tags para controlar anúncios automáticos
		const metaTags = [
			{ name: 'google-adsense-account-verification', content: 'disabled' },
			{ name: 'ads', content: 'limited' },
			{ property: 'og:type', content: 'game' }
		];
		
		metaTags.forEach(tag => {
			const existing = document.querySelector(`meta[name="${tag.name}"], meta[property="${tag.property}"]`);
			if (!existing) {
				const meta = document.createElement('meta');
				if (tag.name) meta.name = tag.name;
				if (tag.property) meta.property = tag.property;
				meta.content = tag.content;
				document.head.appendChild(meta);
			}
		});
	}
	
	function configureAdSense() {
		if (typeof window === 'undefined') return;
		
		// Configurar AdSense para respeitar as áreas protegidas
		window.adsbygoogle = window.adsbygoogle || [];
		
		// Adicionar configuração global
		window.adsbygoogle.push({
			google_ad_client: 'ca-pub-XXXXXXX', // Substitua pelo seu ID
			enable_page_level_ads: false, // Desabilitar anúncios automáticos na página
			overlays: {bottom: false}, // Desabilitar overlays
			anchor_ad: {
				is_auto_ads_anchor: false // Desabilitar anúncios âncora automáticos
			}
		});
	}
	
	onMount(() => {
		addProtectionCSS();
		addMetaTags();
		configureAdSense();
		
		// Observer para detectar novos anúncios
		const observer = new MutationObserver((mutations) => {
			mutations.forEach((mutation) => {
				mutation.addedNodes.forEach((node) => {
					if (node.nodeType === 1 && node.classList?.contains('adsbygoogle')) {
						// Verificar se o anúncio está muito próximo de elementos protegidos
						const rules = protectionRules[protection];
						const gameElements = document.querySelectorAll(rules.selectors.join(', '));
						
						gameElements.forEach((element) => {
							const rect1 = element.getBoundingClientRect();
							const rect2 = node.getBoundingClientRect();
							
							const distance = Math.min(
								Math.abs(rect1.right - rect2.left),
								Math.abs(rect1.left - rect2.right),
								Math.abs(rect1.bottom - rect2.top),
								Math.abs(rect1.top - rect2.bottom)
							);
							
							if (distance < rules.minDistance) {
								console.warn('Anúncio muito próximo de elemento do jogo, aplicando proteção');
								node.style.display = 'none';
							}
						});
					}
				});
			});
		});
		
		observer.observe(document.body, {
			childList: true,
			subtree: true
		});
		
		return () => observer.disconnect();
	});
</script>

<!-- Este componente não renderiza nada, apenas protege contra anúncios intrusivos -->