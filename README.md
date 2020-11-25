# Description
Generate manifests for multiple Kubernetes clusters and teams from a single repository.
* Consistency in configuration across a multi-cloud landscape.
* Separate the configuration release from the deployment process.
* Perfect for GitOps integration when combined with: ArgoCD, Flux or Jenkins.

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

## Download, generate and deploy
Download k8s-onerepo binary and example
```shell
$ git clone https://github.com/bartvanbenthem/k8s-onerepo.git
```

### set environment variables and run example manifest generation
```shell
export K8S_ONEREPO_VALUES='var/values/cluster01'
export K8S_ONEREPO_TEMPLATES='var/templates'
export K8S_ONEREPO_CONFIG='config/cluster'
export K8S_ONEREPO_HELM_TEMPLATES='var/helmcharts'
export K8S_ONEREPO_HELM_CONFIG='config/helmcharts'

$ cd ./k8s-onerepo
$ ./bin/k8s-onerepo
```

### Deploy the generated example configuration to Kubernetes
https://github.com/bartvanbenthem/k8s-onerepo/blob/master/config/README.md

The Example stack contains the following K8s native service configurations:
* Nginx ingress
* Prometheus Operator
* Grafana, Loki and Promtail
