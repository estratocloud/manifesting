---
layout: default
title: Resources
permalink: /config/resources
---

Resources are the main building block of Manifesting, they define what kubernetes resources will be part of your generated manifests file.

The simplest version of them is having a common template that is used for lots of deployments, eg in a micro services environment you might have the following:

<div id="example-name"></div>
<script>
setupFileNavigator("example-name", {
    "manifesting.yaml": `environments:
  - name: "default"
resources:
  - name: "basket"
    template: "micro-deployment"
  - name: "products"
    template: "micro-deployment"
  - name: "checkout"
    template: "micro-deployment"`,
    "templates": {
        "micro-deployment.yaml.gotmpl": `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-{{ .RESOURCE_NAME }}
  labels:
    app: myapp-{{ .RESOURCE_NAME }}
spec:
  replicas: 3
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
          image: docker.io/myapp-microservice-{{ .RESOURCE_NAME }}`
    },
    ".generated": {
        "default.yaml": `apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-basket
  labels:
    app: myapp-basket
spec:
  replicas: 3
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
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-products
  labels:
    app: myapp-products
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp-products
  template:
    metadata:
      name: myapp-products
      labels:
        app: myapp-products
    spec:
      containers:
        - name: myapp
          image: docker.io/myapp-microservice-products
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-checkout
  labels:
    app: myapp-checkout
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp-checkout
  template:
    metadata:
      name: myapp-checkout
      labels:
        app: myapp-checkout
    spec:
      containers:
        - name: myapp
          image: docker.io/myapp-microservice-checkout`,
    }
})
</script>

Any references to names from the environments envvars will have the value included, and anything override will be left alone
