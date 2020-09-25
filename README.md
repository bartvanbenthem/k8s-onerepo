# k8s-onerepo
one repo to rule them all

```shell
.
├── clusters
│   ├── cluster-all
│   │   ├── helmcharts
│   │   │   ├── nginx-ingress
│   │   │   └── prometheus-operator
│   │   ├── manifest
│   │   │   └── allclusters.yaml
│   │   ├── templates
│   │   │   └── allclusters.yaml
│   │   └── values
│   │       └── allclusters.yaml
│   └── cluster-specific
│       ├── cluster-01
│       │   ├── manifest
│       │   └── values
│       │       ├── cluster.yaml
│       │       ├── team-01.yaml
│       │       └── team-02.yaml
│       ├── cluster-02
│       │   ├── manifest
│       │   └── values
│       │       ├── cluster.yaml
│       │       └── team-01.yaml
│       └── templates
│           ├── cluster.yaml
│           └── team.yaml
├── go.mod
├── go.sum
├── k8s-manifestgen
├── k8s-onerepo
├── main.go
└── README.md
```
