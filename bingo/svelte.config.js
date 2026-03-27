import adapter from '@sveltejs/adapter-static';
import { vitePreprocess } from '@sveltejs/vite-plugin-svelte';

// Adiciona preprocess para habilitar suporte a <script lang="ts"> e outros recursos (postcss, etc.)
// Sem isso, o compilador Svelte encontra tokens TypeScript e gera ParseError.
export default {
    preprocess: vitePreprocess(),
    kit: {
        adapter: adapter({
            fallback: 'index.html' // pode variar conforme host
        })
    },
};