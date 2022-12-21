# Blog (blog.sewera.dev)

## Prerequisites

- Install dependencies:
  ```bash
  yarn
  ```

## Development

- Run tests:
  ```bash
  yarn test
  ```
  or:
  ```bash
  yarn test:watch
  ```
- Run a development server:
  ```bash
  yarn dev
  ```

## Vitest

Vitest is a faster alternative to Jest.

### Troubleshooting

Isolation mode has to be on.
Otherwise, test re-runs in watch mode will fail sometimes.

In `vite.config.mjs`:
```js
export default defineConfig({
  // [...]
  test: {
    isolate: true,
  }
})
```

### Vitest VSCode integration

Remember to run vitest VSCode plugin in "Watch mode".
To do this, open the **Testing** tab,
expand a dropdown for run configurations
(a caret by the double play button),
and select **Run Tests (Watch Mode)**.
