# Description
*"one repo to rule them all"*

Manage multiple Kubernetes clusters and teams from a single repository.
* Consistency in configuration across a multi-cloud landscape.
* Separate generic and specific configuration requirements
* Separate configuration and deployments.
* Integrated kubernetes and cloud native services.

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

#### Cloud agnostic services
* Nginx ingress
* Prometheus Operator
* Grafana, Loki and Promtail
* Kubernetes reboot daemon
* OPA Gatekeeper (in development)
* Istio Service Mesh (in development)

#### Cloud native services
* External DNS (in development)
* Cert-Manager (in development)

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/

## Run example
