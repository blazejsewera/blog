/// <reference types="vitest" />
/// <reference types="vite/client" />

import { defineConfig } from 'vite'
import solid from 'vite-plugin-solid'

export default defineConfig({
  plugins: [solid()],
  test: {
    environment: 'jsdom',
    transformMode: {
      web: [/.[jt]sx?/],
    },
    isolate: true,
    setupFiles: ['./vitest.setup.mjs'],
    deps: {
      optimizer: {
        web: {
          enabled: true,
        },
      },
    },
  },
})
