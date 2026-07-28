---
layout: default
title: Template Syntax
permalink: /templating/syntax
---

Manifesting uses [Go Templates] as the templating engine, so any syntax supported there is supported here.
We also provide the sprig functions to provide you with everything you could possibly need to build your manifests.
And there are a few Manifesting specific things available:

```yaml
environments:
  - name: "default"
resources:
  - name: "web"
    template: "web"
```
