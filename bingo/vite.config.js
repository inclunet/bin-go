import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';
// @ts-ignore - campo 'test' suportado pelo Vitest, mas pode não estar no tipo dependendo da versão

export default defineConfig({
	plugins: [sveltekit()],
	// @ts-ignore
	test: {
		environment: 'jsdom'
	},
	server: {
		proxy: {
			'/ws': {
				target: "ws://localhost:8080",
				changeOrigin: true,
				ws: true
			},
			'/qr': {
				target: 'http://localhost:8080',
				changeOrigin: true,
				// rewrite: (path) => path.replace(/^\/api/, '')
			},
			'/api': {
				target: 'http://localhost:8080',
				changeOrigin: true,
				// rewrite: (path) => path.replace(/^\/api/, '')
			}
		}
	},
});
