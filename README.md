# Blog (blog.sewera.dev)

[![Blog CD](https://github.com/blazejsewera/blog/actions/workflows/cd.yml/badge.svg)](https://github.com/blazejsewera/blog/actions)

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
subtitle: Subtitle
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

- Install [Go](https://go.dev)

## Development

- Run tests:
  ```sh
  yarn test
  ```
  or:
  ```sh
  yarn test:watch
  ```
- Run tests with the snapshot updates:
  ```sh
  yarn test:update
  ```
- Run a development server:
  ```sh
  yarn dev
  ```

## Deploy and rollback

The blog is deployed on every push to `master`.
All the applicable deployment and rollback workflows,
both automatic and manual are in the [actions tab].

[actions tab]: https://github.com/blazejsewera/blog/actions

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

## License

TLDR: All files in `src/pages` directory are under CC-BY-SA-4.0.
All other files are under GPL-3.0.
There can be exceptions,
which have to be clearly specified in the licensed file.

A more detailed explanation can be found in the [LICENSE](./LICENSE) file.

Photos are sourced from [Unsplash](https://unsplash.com)
under the [Unsplash license](https://unsplash.com/license),
unless stated otherwise.
