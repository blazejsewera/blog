import { defineConfig } from 'astro/config'
import solidJs from '@astrojs/solid-js'
import tailwind from '@astrojs/tailwind'

// https://astro.build/config
export default defineConfig({
  integrations: [solidJs(), tailwind()],
  markdown: {
    drafts: true,
    syntaxHighlight: 'shiki',
    shikiConfig: {
      theme: 'github-light',
    },
  },
})
