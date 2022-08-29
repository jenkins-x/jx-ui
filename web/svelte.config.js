import preprocess from 'svelte-preprocess'
import { resolve } from 'path'
import adapter from '@sveltejs/adapter-static'
import { configDefaults } from 'vitest/config'

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
    adapter: adapter({
      pages: 'build',
      assets: 'build',
      fallback: 'index.html',
    }),
    files: {
      routes: `src/routes/${routeFolder}`,
    },
    prerender: {
      default: false,
    },
  },
}

export default config
