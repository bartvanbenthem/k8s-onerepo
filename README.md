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

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/

#### Example k8s services
* Nginx ingress
* Prometheus Operator
* Grafana, Loki and Promtail

## Run example
```shell
$ git clone https://github.com/bartvanbenthem/k8s-onerepo.git
```

#### set environment variables and run example
```shell
export K8S_ONEREPO_VALUES='var/values/cluster01'
export K8S_ONEREPO_TEMPLATES='var/templates'
export K8S_ONEREPO_CONFIG='config/cluster'
export K8S_ONEREPO_HELM_TEMPLATES='var/helmcharts'
export K8S_ONEREPO_HELM_CONFIG='config/helmcharts'

$ cd ./k8s-onerepo
$ ./bin/k8s-onerepo
```