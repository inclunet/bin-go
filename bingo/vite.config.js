import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';

export default defineConfig({
	plugins: [sveltekit()],
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
