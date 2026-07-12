---
layout: default
title: Environments
permalink: /config/environments
---

Environments are a core part of Manifesting, as they're a common part of how kubernetes is managed (eg different configuration for Production than for Non-Prod).
If you don't require any different configuration, then you can simply define a default environment and ignore them:

```yaml
environments:
  - name: "default"
```

For everybody else, let's see what environments can do...

<div id="example-per-environment"></div>
<script>
setupFileNavigator("example-per-environment", {
    "manifesting.yaml": `environments:
  - name: "production"
  - name: "nonprod"
resources:
  - name: "web"
    template: "web"
    vars:
      Replicas: { production: 5, nonprod: 1 }`,
    "templates": {
        "web.yaml.gotmpl": `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-web
  labels:
    app: myapp-web
spec:
  replicas: {{ perEnvironment .Replicas }}
  selector:
    matchLabels:
      app: myapp-web
  template:
    metadata:
      name: myapp-web
      labels:
        app: myapp-web
    spec:
      containers:
        - name: web
          image: caddy:2.6-alpine`,
    },
    ".generated": {
        "nonprod.yaml": `apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-web
  labels:
    app: myapp-web
spec:
  replicas: 1
  selector:
    matchLabels:
      app: myapp-web
  template:
    metadata:
      name: myapp-web
      labels:
        app: myapp-web
    spec:
      containers:
        - name: web
          image: caddy:2.6-alpine`,
        "production.yaml": `apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-web
  labels:
    app: myapp-web
spec:
  replicas: 5
  selector:
    matchLabels:
      app: myapp-web
  template:
    metadata:
      name: myapp-web
      labels:
        app: myapp-web
    spec:
      containers:
        - name: web
          image: caddy:2.6-alpine`,
    }
})
</script>

Another feature of environments is avoiding duplication of envvars, if you've got several manifests that share a common set of envvars, you can define them once, and then just "include" the ones you need:

<div id="example-env-from"></div>
<script>
setupFileNavigator("example-env-from", {
    "manifesting.yaml": ``,
    "templates": {
        "web.yaml.gotmpl": ``,
    },
    ".generated": {
        "nonprod.yaml": ``,
        "production.yaml": ``,
    }
})
</script>


```yaml
environments:
  - name: "production"
    envFrom: "templates/envvars.yaml"
resources:
  - name: "basket"
    template: "micro-deployment"
```

```yaml
- name: ACTIVE_FEATURE_FLAGS
  value: "2457"
- name: RUN_IN_DEBUG_MODE
  value: "false"
```

```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-{{ .RESOURCE_NAME }}
  labels:
    app: myapp-{{ .RESOURCE_NAME }}
spec:
  replicas: 10
  selector:
    matchLabels:
      app: myapp-{{ .RESOURCE_NAME }}
  template:
    metadata:
      name: myapp-{{ .RESOURCE_NAME }}
      labels:
        app: myapp-{{ .RESOURCE_NAME }}
    spec:
      containers:
        - name: myapp
          image: docker.io/myapp-microservice-{{ .RESOURCE_NAME }}
          env:
            - name: ACTIVE_FEATURE_FLAGS
            - name: ENVIRONMENT
              value: "{{ .ENVIRONMENT }}"
            - name: RUN_IN_DEBUG_MODE
```

Any references to names from the environments envvars will have the value included, and anything override will be left alone:
```yaml
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-basket
  labels:
    app: myapp-basket
spec:
  replicas: 10
  selector:
    matchLabels:
      app: myapp-basket
  template:
    metadata:
      name: myapp-basket
      labels:
        app: myapp-basket
    spec:
      containers:
        - name: myapp
          image: docker.io/myapp-microservice-basket
          env:
            - name: ACTIVE_FEATURE_FLAGS
              value: "2457"
            - name: ENVIRONMENT
              value: "production"
            - name: RUN_IN_DEBUG_MODE
              value: "false"
```
