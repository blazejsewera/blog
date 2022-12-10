/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['./src/**/*.{astro,html,js,jsx,md,mdx,svelte,ts,tsx,vue}'],
  theme: {
    fontFamily: {
      geometric: [
        'Euclid, Futura, "Trebuchet MS", Arial, sans-serif',
        { fontFeatureSettings: '"kern"' },
      ],
      sans: [
        '"Suisse Int\'l", -apple-system, BlinkMacSystemFont, Helvetica, Arial, sans-serif',
        { fontFeatureSettings: '"kern"' },
      ],
    },
    extend: {},
  },
  plugins: [],
}
