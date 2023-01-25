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

## Deployment

Github Actions should make a compressed archive
named `blog-sewera-dev-dist-archive`.
The only gimmick is that it is tarred and zipped separately,
so it's not `.tar.gz`, but a `.tar` inside a `.zip` file.

Normal `.tar.gz` can be also created using `yarn tar:gz` script.

To download an archive built in Github Actions:

1. Get download URL of the latest dist:
   ```bash
   gh api \
     -H 'Accept: application/vnd.github+json' \
     '/repos/blazejsewera/blog/actions/artifacts?per_page=1&name=blog-sewera-dev-dist-archive' \
   | jq -r '.artifacts[0].archive_download_url'
   ```
   or
   ```bash
   curl \
     -H 'Accept: application/vnd.github+json' \
     -H 'Authorization: Bearer <YOUR-TOKEN>'\
     -H 'X-GitHub-Api-Version: 2022-11-28' \
     'https://api.github.com/repos/blazejsewera/blog/actions/artifacts?per_page=1&name=blog-sewera-dev-dist-archive' \
   | jq -r '.artifacts[0].archive_download_url'
   ```
2. Download it:
   ```bash
   gh api \
     -H 'Accept: application/vnd.github+json' \
     <URL_from_previous_point> > blog-sewera-dev-dist-archive.zip
   ```
   or
   ```bash
   curl \
    -L \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer <YOUR-TOKEN>"\
    -H "X-GitHub-Api-Version: 2022-11-28" \
    -o 'blog-sewera-dev-dist-archive.zip' \
    '<URL_from_previous_point>'
   ```
3. Extract it twice (once may be enough, depending on the `tar` implementation):
   ```bash
   tar -xzvf blog-sewera-dev-dist-archive.zip && tar -xvf dist.tar
   ```
4. Remove archives:
   ```bash
   rm -f blog-sewera-dev-dist-archive.zip dist.tar
   ```

A fine-grained personal access token will be necessary.
It should have read-only access to metadata and Github Actions
(to download artifacts).
It has to be set for `blazejsewera` organization,
and `blog` repo.
To create it, refer to the [documentation].

[documentation]: https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token#creating-a-fine-grained-personal-access-token

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
