# Description
Manage multiple Kubernetes clusters and teams from a single repository.
* Consistency in configuration across a multi-cloud landscape.
* Separate generic and specific configuration requirements
* Separate configuration and deployments.
* Integrated kubernetes and cloud native services.
* Monitoring and observability of the complete stack.

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

#### Included Kubernetes native services
* Nginx ingress
* Prometheus Operator
* Grafana, Loki and Promtail
* Kubernetes reboot daemon
* OPA Gatekeeper (in development)
* Open Service Mesh (in development)

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/

## Run example
