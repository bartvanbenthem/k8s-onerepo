# k8s-onerepo
one repo to rule them all

```shell
.
├── clusters
│   ├── all-clusters
│   │   ├── helmcharts
│   │   │   ├── nginx-ingress
│   │   │   └── prometheus-operator
│   │   └── raw
│   │       ├── manifest
│   │       │   └── allclusters.yaml
│   │       ├── templates
│   │       │   └── allclusters.yaml
│   │       └── values
│   │           └── allclusters.yaml
│   ├── cluster-01
│   │   ├── cluster-manifest
│   │   │   ├── manifest
│   │   │   ├── templates
│   │   │   │   └── clusters.yaml
│   │   │   └── values
│   │   │       └── clusters.yaml
│   │   └── teams-manifest
│   │       ├── manifest
│   │       ├── templates
│   │       │   └── team.yaml
│   │       └── values
│   │           ├── team-01.yaml
│   │           └── team-02.yaml
│   └── cluster-02
│       ├── cluster-manifest
│       │   ├── manifest
│       │   ├── templates
│       │   │   └── clusters.yaml
│       │   └── values
│       │       └── clusters.yaml
│       └── teams-manifest
│           ├── manifest
│           ├── templates
│           │   └── team.yaml
│           └── values
│               ├── team-01.yaml
│               └── team-02.yaml
├── go.mod
├── k8s-manifestgen
├── k8s-onerepo
├── main.go
└── README.md
```
