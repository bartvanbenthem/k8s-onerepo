## Deploy K8sCluster Configurations
```shell
$ clusterName=$(kubectl config current-context)

$ kubectl apply -f ./clusters/all-clusters/
$ kubecrl apply -f ./clusters/$clusterName

$ helm install co-nginx \
    -f .config/clusters/$clusterName/nginx-helm.yaml \
    ./config/helmcharts/ingress-nginx
$ helm install co-prometheus \
    -f .config/clusters/$clusterName/prometheus-helm.yaml \
    ./config/helmcharts/kube-prometheus-stack

```


