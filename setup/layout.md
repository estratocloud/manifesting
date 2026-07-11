---
layout: default
title: Layout
permalink: /setup/layout
---

By default Manifesting uses the following filesystem layout:

```
myapp/
├── manifesting.yaml
├── templates/
│   ├── micro-deployment.yaml.gotmpl
│   ├── micro-service.yaml.gotmpl
├── .generated/
│   └── production.yaml
│   └── nonprod.yaml
```

The `templates` directory contains your manifest templates, the `manifesting.yaml` is the configuration that controls what is generated, and the `.generated` directory contains the produced manifests.

We recommend committing the `.generated/` files to source control, and then using CI/CD to regenerate them whenever the templates are changed, that way you can always see exactly what the impact of a particular change is.

You can change the location of the `manifesting.yaml` config file by using `--config` option:
```
manifesting --config=config.yaml
```

You can override the working directory by using `--working-dir`:
```
manifesting --working-dir=/tmp/deploy
```
