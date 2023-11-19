---
layout: ../layouts/Post.astro
title: Github Actions iOS Shortcut
date: 2023-01-18
keywords:
  - iOS Shortcuts
  - Github Actions
  - Manual Github Action Trigger
  - Github Workflow Dispatch
imgUrl: /github-actions-ios-shortcut/bus-shortcut.jpg
imgDescription: A bus going through a shortcut
abstract: |
  Manual Github Workflow dispatching can be a convenient way
  to trigger repeatable tasks that are not tied to our usual
  Continuous Delivery Pipeline. For example, we might need to
  quickly trigger a rollback of our system, when we get
  an on-call incident. We might use our computer and trigger
  the workflow in the Github web app.
  However, quickly dispatching a workflow with iOS Shortcuts
  may be more handy. This article describes how to set it up.
---

To dispatch a Github Actions Workflow,
iPhone Shortcuts app requires a couple of parameters:

- owner,
- workflow id,
- repository,
- branch/ref,
- inputs, and
- account.

They can be easily obtained with the Github CLI tool[^gh-cli].
The following snippet is taken from the Github API docs[^gh-workflows].

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
  -d '{"ref":"topic-branch","inputs":{"name":"Example name","home":"San Francisco, CA"}}'
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
      "name": "Example name",
      "home": "San Francisco, CA"
    }
  account: github-username
```

[^gh-cli]: [Github CLI tool (`gh`)](https://cli.github.com/) (accessed Feb. 05, 2023)
[^gh-workflows]: [List repository workflows — Github REST API docs](https://docs.github.com/en/rest/actions/workflows#list-repository-workflows) (accessed Feb. 05, 2023)
[^rq]: [Create a workflow dispatch event — Github REST API docs](https://docs.github.com/en/rest/actions/workflows#create-a-workflow-dispatch-event) (accessed Feb. 05, 2023)

<!-- [[github-workflow-dispatch-iphone-shortcut]] -->
