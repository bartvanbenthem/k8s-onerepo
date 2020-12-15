## Deploy Configuration
```shell
# set cluster to deploy
cluster='cluster01'

# install raw manifests
kubectl apply -f ./config/$cluster/generic-cluster.yaml
kubectl apply -f ./config/$cluster/team-alpha.yaml
kubectl apply -f ./config/$cluster/team-beta.yaml

# install the default nginx ingress controller
helm install co-nginx \
  -f ./config/$cluster/nginx-helm.yaml \
  ./config/helmcharts/ingress-nginx --namespace co-ingress

# install prometheus operator
helm install co-prometheus \
  -f ./config/$cluster/prometheus-helm.yaml \
  ./config/helmcharts/kube-prometheus-stack --namespace co-monitoring
# get grafana admin password
kubectl get secret --namespace co-monitoring co-prometheus-grafana \
  -o jsonpath="{.data.admin-password}" | base64 --decode ; echo
    
# install loki
helm install co-loki ./config/helmcharts/loki --namespace co-monitoring

# install promtail
helm install co-promtail \
  -f ./config/$cluster/promtail-helm.yaml \
  ./config/helmcharts/promtail --namespace co-monitoring
#--set "loki.serviceName=co-loki"

# deploy these manifests after required CRD are created by HELM charts
kubectl apply -f ./config/$cluster/monitoring-postcrd.yaml

```

#### Test and expose (do not use in production)
``` shell
# Expose Prometheus Operator monitoring interfaces
cat <<EOF | kubectl apply -f -
apiVersion: networking.k8s.io/v1beta1
kind: Ingress
metadata:
  name: co-monitoring
  namespace: co-monitoring
  annotations:
    kubernetes.io/ingress.class: "nginx"
spec:
  spec:
  rules:
  - host: prometheus
    http:
      paths:
        - path: /
          backend:
            serviceName: co-prometheus-kube-prometh-prometheus
            servicePort: 9090
  - host: grafana
    http:
      paths:
        - path: /
          backend:
            serviceName: co-prometheus-grafana
            servicePort: 80
  - host: alertmanager
    http:
      paths:
        - path: /
          backend:
            serviceName: co-prometheus-kube-prometh-alertmanager 
            servicePort: 9093
EOF

# test promtail metrics working
kubectl --namespace co-monitoring port-forward daemonset/co-promtail 3101 
curl http://127.0.0.1:3101/metrics

# verify loki working
kubectl --namespace co-monitoring port-forward service/co-loki 3100
curl http://127.0.0.1:3100/api/prom/label

```
#### Grafana config
* Import nginx ingress dashboard: 9614 
* Add the Loki datasource


## Remove Configuration
```shell
# remove k8s components
kubectl delete-f ./config/$cluster/generic-cluster.yaml
kubectl delete -f ./config/$cluster/team-alpha.yaml
kubectl delete -f ./config/$cluster/team-beta.yaml

# remove helm charts
helm uninstall co-nginx --namespace co-ingress
helm uninstall co-nginx-internal --namespace co-ingress-internal
helm uninstall co-prometheus --namespace co-monitoring
helm uninstall co-promtail --namespace co-monitoring
helm uninstall co-loki --namespace co-monitoring

# delete all namespaces
kubectl delete ns co-ingress co-ingress-internal co-maintenance co-monitoring

```
