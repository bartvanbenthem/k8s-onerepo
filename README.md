# k8s-onerepo
one repo to rule them all

```shell
.
├── build-all-manifests.sh
├── clusters
│   ├── all-clusters
│   │   ├── helmcharts
│   │   │   ├── nginx-ingress
│   │   │   └── prometheus-operator
│   │   └── raw
│   │       ├── generated
│   │       │   └── allclusters.yaml
│   │       ├── templates
│   │       │   └── allclusters.yaml
│   │       └── values
│   │           └── allclusters.yaml
│   ├── cluster-01
│   │   ├── cluster-manifest
│   │   └── teams-manifest
│   └── cluster-02
│       ├── cluster-manifest
│       └── teams-manifest
├── go.mod
├── k8s-manifestgen
├── k8s-onerepo
├── main.go
└── README.md
```
