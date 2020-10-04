## Deploy Configuration
```shell
# set cluster name
clusterName=$(kubectl config current-context)

# install raw manifests
kubectl apply -f ./config/clusters/all-clusters
kubectl apply -f ./config/clusters/$clusterName/spec-cluster.yaml
kubectl apply -f ./config/clusters/$clusterName/team-alpha.yaml
kubectl apply -f ./config/clusters/$clusterName/team-beta.yaml

# install the default nginx ingress controller
helm install co-nginx \
    -f ./config/clusters/$clusterName/nginx-helm.yaml \
    ./config/helmcharts/ingress-nginx --namespace co-ingress

# install nginx ingress controller for internal traffic
helm install co-nginx-internal \
    -f ./config/clusters/$clusterName/nginx-internal-helm.yaml \
    ./config/helmcharts/ingress-nginx --namespace co-ingress-internal

# install prometheus operator
helm install co-prometheus \
    -f ./config/clusters/$clusterName/prometheus-helm.yaml \
    ./config/helmcharts/kube-prometheus-stack --namespace co-monitoring
# get grafana admin password
kubectl get secret --namespace co-monitoring co-prometheus-grafana \
    -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

# install kured
helm install co-kured ./config/helmcharts/kured --namespace co-maintenance

# install open policy agent gate keeper
helm install -f ./config/clusters/$clusterName/gatekeeper-helm.yaml \
    ./config/helmcharts/gatekeeper --generate-name --namespace co-policy

# deploy these manifests after required CRD are created by HELM charts
kubectl apply -f ./config/clusters/$clusterName/post-crd.yaml

```
#### import the following dashboards in grafana
* 9614 

#### Test and expose prometheus stack (do not use in production)
``` shell
# Expose monitoring interfaces
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
```



## Update Configuration
```shell

```

## Remove Configuration
```shell
# remove helm charts
helm uninstall co-nginx --namespace co-ingress
helm uninstall co-nginx-internal --namespace co-ingress-internal
helm uninstall co-prometheus --namespace co-monitoring
helm uninstall co-kured --namespace co-maintenance

# delete all namspaces
kubectl delete ns co-ingress co-ingress-internal co-maintenance co-monitoring co-opa

#clean up gatekeeper crd
kubectl delete crd \
  configs.config.gatekeeper.sh \
  constraintpodstatuses.status.gatekeeper.sh \
  constrainttemplatepodstatuses.status.gatekeeper.sh \
  constrainttemplates.templates.gatekeeper.sh

```
