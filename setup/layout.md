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

We recommend committing the `.generated/` files to source control, and then using CI/CD to regenerate them whenever the templates are changed, that way you can always see exactly what the impact of a particular change is. We provide a GitHub Action to handle this for Pull Requests.

You can change the location of the `manifesting.yaml` config file by using `--config` option:
```
manifesting --config=config.yaml
```

You can change the location of the `templates/` directory (and the filename suffix) by using `--templates-dir` and `--templates-suffix`:
```
manifesting --templates-dir=k8s --templates-suffix=.yaml
```

You can change the location of the generated manifests by using `--generated-dir`:
```
manifesting --generated-dir=manifests
```

Finally you can override the working directory by using `--working-dir`:
```
manifesting --working-dir=/tmp/deploy
```

Using all of the above options together would result in the following filesystem layout:

```
/tmp/deploy/
├── config.yaml
├── k8s/
│   ├── micro-deployment.yaml
│   ├── micro-service.yaml
├── manifests/
│   └── production.yaml
│   └── nonprod.yaml
```
