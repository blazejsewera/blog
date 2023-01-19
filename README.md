# Blog (blog.sewera.dev)

## Writing articles

Minimal frontmatter

```yaml
---
layout: ../layouts/Post.astro
title: Title
date: 1970-01-01
---
```

Complete frontmatter

```yaml
---
layout: ../layouts/Post.astro
title: Title
date: 1970-01-01
author: John Doe # default: Blazej Sewera
license: MIT # default: CC BY-SA 4.0
language: en-US # default: en-US
draft: true # default: false
draftDescription: This is a draft # default: This is a draft article. It may be incomplete.
imgUrl: /image.jpg # default: null
imgDescription: Sample image # default: ''
abstract: | # default: ''
  A short description of an article
keywords: # default []
  - kw1
  - kw2
---
```

## Prerequisites

- Install dependencies and download fonts:
  ```bash
  yarn
  yarn fonts
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
  },
})
```

### Vitest VSCode integration

Remember to run vitest VSCode plugin in "Watch mode".
To do this, open the **Testing** tab,
expand a dropdown for run configurations
(a caret by the double play button),
and select **Run Tests (Watch Mode)**.
