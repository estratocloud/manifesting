---
layout: default
title: Variables
permalink: /config/variables
---

Variables allow you to generate different resources from a common template, for example...

<div id="example-basic"></div>
<script>
setupFileNavigator("example-basic", {
    "manifesting.yaml": `environments:
  - name: "default"
resources:
  - name: "basket"
    template: "micro-deployment"
    vars:
      Replicas: 3
      Image: basket-service
  - name: "products"
    template: "micro-deployment"
    vars:
      Replicas: 10
      Image: productsv2
  - name: "checkout"
    template: "micro-deployment"
    vars:
      Replicas: 5
      Image: ordercomplete`,
    "templates": {
        "micro-deployment.yaml.gotmpl": `---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-{{ .RESOURCE_NAME }}
  labels:
    app: myapp-{{ .RESOURCE_NAME }}
spec:
  replicas: {{ .Replicas }}
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
          image: docker.io/myapp-{{ .Image }}`
    },
    ".generated": {
        "production.yaml": `apiVersion: apps/v1
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
          image: docker.io/myapp-basket-service
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-products
  labels:
    app: myapp-products
spec:
  replicas: 10
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
          image: docker.io/myapp-productsv2
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp-checkout
  labels:
    app: myapp-checkout
spec:
  replicas: 5
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
          image: docker.io/myapp-ordercomplete`,
    }
})
</script>
