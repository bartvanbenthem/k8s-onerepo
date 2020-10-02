# Description
*"one repo to rule them all"*

Manage multiple Kubernetes clusters and teams from a single repository.
* Consistency in configuration across a multi-cloud landscape.
* Separate generic and specific configuration requirements
* Separate configuration and deployments.

## project structure
```shell
.
├── config
│   ├── clusters
│   └── helmcharts
├── utils
│   └── filesystem
│   └── manifestgen
└── var
    ├── clusters
    │   ├── templates
    │   └── values
    └── helmcharts
```

### Cloud agnostic services
* Nginx ingress
* Prometheus Operator
* Kured
* OPA Gatekeeper (soon)

### Cloud native services
* External DNS (soon)
* Cert-Manager (soon)

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/

## Run example
