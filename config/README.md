## Deploy K8sCluster Configurations
```shell
$ clusterName=$(kubectl config current-context)

$ kubectl apply -f ./cluster-all
$ kubecrl apply -f ./cluster-specific/$clusterName

$ helm install cro-nginx -f ./helmcharts/ingress-nginx/values-override.yaml ./helmcharts/ingress-nginx
$ helm install cro-prometheus -f ./helmcharts/kube-prometheus-stack/values-override.yaml ./helmcharts/kube-prometheus-stack
```


