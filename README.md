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
│       ├── manifest
│       ├── templates
│       │   ├── clusters.yaml
│       │   └── team.yaml
│       └── values
│           ├── cluster-01
│           │   ├── cluster.yaml
│           │   ├── team-01.yaml
│           │   └── team-02.yaml
│           └── cluster-02
│               ├── cluster.yaml
│               ├── team-01.yaml
│               └── team-02.yaml
├── go.mod
├── go.sum
├── k8s-manifestgen
├── k8s-onerepo
├── main.go
└── README.md
```
