---
layout: default
title: Config
permalink: /setup/config
---

To use Manifesting you must provide a config file to control how your manifests are generated, at a minimum it must declare at least one [environment](../config/environments) and one [resource](../config/resources).

```yaml
environments:
  - name: "default"
resources:
  - name: "web"
    template: "web"
```
