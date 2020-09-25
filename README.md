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
│   ├── cluster-all
│   │   └── allclusters.yaml
│   ├── cluster-specific
│   │   ├── cluster-01
│   │   └── cluster-02
│   └── helmcharts
│       ├── ingress-nginx
│       │   ├── Chart.yaml
│       │   └── values.yaml
│       ├── kube-prometheus-stack
│       │   ├── Chart.yaml
│       │   └── values.yaml
│       └── values
│           └── values.yaml
└── values
    ├── cluster-all
    │   ├── templates
    │   │   └── allclusters.yaml
    │   └── values
    │       └── allclusters.yaml
    ├── cluster-specific
    │   ├── templates
    │   │   ├── cluster.yaml
    │   │   └── team.yaml
    │   ├── values-cluster-01
    │   │   ├── cluster.yaml
    │   │   ├── team-01.yaml
    │   │   └── team-02.yaml
    │   └── values-cluster-02
    │       ├── cluster.yaml
    │       └── team-01.yaml
    └── helmcharts
        └── values.yaml
├── go.mod
├── go.sum
├── k8s-onerepo
├── main.go
├── README.md
```

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/


