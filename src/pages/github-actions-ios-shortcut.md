---
layout: ../layouts/Post.astro
title: Github Actions iOS Shortcut
date: 2023-01-18
draft: true
keywords:
  - iOS Shortcuts
  - Github Actions
  - Manual Github Action Trigger
  - Github Workflow Dispatch
imgUrl: /image/nordic-sky.jpg
---

To dispatch a Github Actions Workflow,
iPhone Shortcuts app requires a couple of parameters:

- owner,
- workflow id,
- repository,
- branch/ref,
- inputs,
- account.

They can be obtained with [`gh` CLI tool](https://cli.github.com/)[^gh-workflows].

```sh
gh api \
  -H "Accept: application/vnd.github+json" \
  /repos/<owner>/<repo>/actions/workflows
```

The `inputs` have to be in JSON format, as in the following request snippet[^rq].

```sh
curl \
  -X POST \
  -H "Accept: application/vnd.github+json" \
  -H "Authorization: Bearer <YOUR-TOKEN>"\
  -H "X-GitHub-Api-Version: 2022-11-28" \
  <github api url>/repos/<owner>/<repo>/actions/workflows/<workflow id>/dispatches \
  -d '{"ref":"topic-branch","inputs":{"name":"Mona the Octocat","home":"San Francisco, CA"}}'
```

## Example

```yaml
shortcut:
  name: Dispatch Workflow
  owner: github-username
  workflowId: 00000000
  repository: your-repo
  branchOrRef: master
  inputs: |
    {
      "rebuild": true
    }
  account: github-username
```

[^gh-workflows]: <https://docs.github.com/en/rest/actions/workflows#list-repository-workflows>
[^rq]: <https://docs.github.com/en/rest/actions/workflows#create-a-workflow-dispatch-event>

<!-- [[github-workflow-dispatch-iphone-shortcut]] -->
