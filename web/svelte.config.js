import preprocess from 'svelte-preprocess'
import { resolve } from 'path'
import adapter from '@sveltejs/adapter-static'

let routeFolder = process.env.ROUTE_FOLDER

if (routeFolder == undefined) {
  routeFolder = 'landing'
}

/** @type {import('@sveltejs/kit').Config} */
const config = {
  // Consult https://github.com/sveltejs/svelte-preprocess
  // for more information about preprocessors
  preprocess: preprocess(),

  kit: {
    // hydrate the <div id="svelte"> element in src/app.html
    adapter: adapter(),
    files: {
      routes: `src/routes/${routeFolder}`,
    },
    prerender: {
      enabled: false,
    },
    vite: {
      resolve: {
        alias: {
          $src: resolve('./src'),
          $stores: resolve('./src/stores'),
          $assets: resolve('./src/assets'),
          $icon: resolve('./node_modules/svelte-bootstrap-icons/lib'),
        },
      },
      test: {
        environment: 'jsdom',
        deps: {
          inline: ['date-fns'],
        },
        globals: true,
        setupFiles: ['src/setupTests.ts'],
      },
    },
  },
}

export default config
