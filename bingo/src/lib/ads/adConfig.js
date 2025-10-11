// Configuração de anúncios para o jogo da velha
export const adConfig = {
	// Substitua pelos seus IDs reais do AdSense
	client: 'ca-pub-XXXXXXX', // Seu ID de cliente AdSense
	
	slots: {
		topBanner: 'XXXXXXXXX',     // Banner superior
		rectangle: 'XXXXXXXXX',     // Retângulo lateral/meio
		bottomBanner: 'XXXXXXXXX',  // Banner inferior
		sidebar: 'XXXXXXXXX'        // Barra lateral (apenas desktop)
	},
	
	// Configurações de comportamento
	settings: {
		// Proteção rigorosa durante o jogo
		protectionLevel: 'strict', // 'strict', 'medium', 'light'
		
		// Atrasar carregamento de anúncios para não afetar performance
		loadDelay: 2000, // ms
		
		// Tamanhos máximos permitidos
		maxSizes: {
			mobile: { width: 320, height: 100 },
			tablet: { width: 728, height: 150 },
			desktop: { width: 728, height: 200 }
		},
		
		// Posições permitidas por tipo de dispositivo
		allowedPlacements: {
			mobile: ['top', 'bottom'],
			tablet: ['top', 'bottom', 'rectangle'],
			desktop: ['top', 'bottom', 'rectangle', 'sidebar']
		},
		
		// Distâncias mínimas dos elementos do jogo (pixels)
		minDistances: {
			fromBoard: 100,
			fromScoreboard: 50,
			fromControls: 80
		}
	},
	
	// Áreas que nunca devem ter anúncios
	blockedAreas: [
		'.tic-board',
		'.tic-cell', 
		'.scoreboard-container',
		'.game-controls',
		'.help-modal',
		'[role="dialog"]',
		'[aria-modal="true"]'
	],
	
	// Meta tags para controle de anúncios automáticos
	metaTags: {
		'google-adsense-account-verification': 'disabled',
		'ads': 'limited',
		'robots': 'index, follow, noarchive',
		'og:type': 'game'
	}
};

// Função para verificar se um anúncio deve ser exibido
export function shouldShowAd(placement, deviceType) {
	const allowed = adConfig.settings.allowedPlacements[deviceType] || [];
	return allowed.includes(placement);
}

// Função para obter configuração de tamanho
export function getAdSize(deviceType, format) {
	const maxSize = adConfig.settings.maxSizes[deviceType];
	
	const formatSizes = {
		banner: { width: Math.min(728, maxSize.width), height: Math.min(90, maxSize.height) },
		rectangle: { width: Math.min(300, maxSize.width), height: Math.min(250, maxSize.height) },
		square: { width: Math.min(250, maxSize.width), height: Math.min(250, maxSize.height) }
	};
	
	return formatSizes[format] || formatSizes.rectangle;
}

// Função para detectar tipo de dispositivo
export function getDeviceType() {
	if (typeof window === 'undefined') return 'desktop';
	
	const width = window.innerWidth;
	if (width <= 480) return 'mobile';
	if (width <= 768) return 'tablet';
	return 'desktop';
}