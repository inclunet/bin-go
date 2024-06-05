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
			'/api': {
				target: 'http://localhost:8080',
				changeOrigin: true,
				// rewrite: (path) => path.replace(/^\/api/, '')
			}
		}
	},
});
