## Deploy K8sCluster Configurations
```shell
$ clusterName=$(kubectl config current-context)

$ kubectl apply -f ./cluster-all
$ kubecrl apply -f ./cluster-specific/$clusterName

$ helm install ore-nginx ./helmcharts/ingress-nginx -f ./helmcharts/ingress-nginx/values-override.yaml
$ helm install ore-prometheus ./helmcharts/kube-prometheus-stack -f ./helmcharts/kube-prometheus-stack/values-override.yaml
```


