import adapter from '@sveltejs/adapter-static';
import { sveltePreprocess } from 'svelte-preprocess';

const config = {
  preprocess: sveltePreprocess({
    scss: {
      // opcional: você pode injetar variáveis/mixins globais aqui
      // prependData: `@use 'src/styles/mixins.scss' as *;`
    }
  }),
  kit: {
    adapter: adapter({
      fallback: 'index.html'
    })
  }
};

export default config;
