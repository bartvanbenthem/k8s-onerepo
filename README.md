# Description
Generate manifests for multiple Kubernetes clusters and teams from a single repository.
* Consistency in configuration across a multi-cloud landscape.
* Separate generic and specific configuration requirements.
* Separate configuration from deployment.

## project structure
```shell
.
├── config
│   ├── cluster
│   └── helmcharts
├── utils
│   └── filesystem
│   └── manifestgen
└── var
    └── helmcharts
    └── templates
    └── values
```

#### Example helm charts
* Nginx ingress
* Prometheus Operator
* Grafana, Loki and Promtail

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/

## Run example
```shell
$ git clone https://github.com/bartvanbenthem/k8s-onerepo.git
$ cd ./k8s-onerepo
$ ./k8s-onerepo
```
