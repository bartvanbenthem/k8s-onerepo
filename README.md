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
│   ├── cluster-specific
│   └── helmcharts
├── utils
│   └── deploy
│   └── filesystem
│   └── manifestgen
└── var
    ├── cluster-all
    │   ├── templates
    │   └── values
    ├── cluster-specific
    │   ├── templates
    │   └── values
    └── helmcharts
    │   ├── templates
    │   └── values
```

## prerequisites
* Install kubectl: https://kubernetes.io/docs/tasks/tools/install-kubectl/
* Install Helm: https://helm.sh/docs/intro/install/


