# Description
*"one repo to rule them all"*

## project structure
```shell
.
.
├── clusters
│   ├── cluster-all
│   │   ├── manifest
│   │   │   └── allclusters.yaml
│   │   ├── templates
│   │   │   └── allclusters.yaml
│   │   └── values
│   │       └── allclusters.yaml
│   ├── cluster-specific
│   │   ├── cluster-01
│   │   │   ├── manifest
│   │   │   └── values
│   │   │       ├── cluster.yaml
│   │   │       ├── team-01.yaml
│   │   │       └── team-02.yaml
│   │   ├── cluster-02
│   │   │   ├── manifest
│   │   │   └── values
│   │   │       ├── cluster.yaml
│   │   │       └── team-01.yaml
│   │   └── templates
│   │       ├── cluster.yaml
│   │       └── team.yaml
│   └── helmcharts
│       ├── ingress-nginx
│       │   ├── Chart.yaml
│       │   └── values.yaml
│       └── kube-prometheus-stack
│           ├── Chart.yaml
│           └── values.yaml
├── go.mod
├── go.sum
├── k8s-onerepo
├── main.go
└── README.md
```

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/


