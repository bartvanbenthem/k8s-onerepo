## Deploy Configuration
```shell
clusterName=$(kubectl config current-context)

kubectl apply -f ./config/clusters/all-clusters
kubectl apply -f ./config/clusters/$clusterName

helm install co-nginx \
    -f ./config/clusters/$clusterName/nginx-helm.yaml \
    ./config/helmcharts/ingress-nginx --namespace co-ingress-internal

helm install co-prometheus \
    -f ./config/clusters/$clusterName/prometheus-helm.yaml \
    ./config/helmcharts/kube-prometheus-stack --namespace co-monitoring

helm install co-gatekeeper ./config/helmcharts/gatekeeper --namespace co-opa

helm install co-kured ./config/helmcharts/kured --namespace co-maintenance

```
