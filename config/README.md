## Deploy Configuration
```shell
# set cluster name
clusterName=$(kubectl config current-context)

# install raw manifests
kubectl apply -f ./config/clusters/all-clusters
kubectl apply -f ./config/clusters/$clusterName

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
    ./config/helmcharts/gatekeeper --generate-name --namespace co-opa

```

#### Local-lab environments (microk8s, k3s, etc.)
Edit the nginx ingress service manifests and add the nodes ip to the Loadbalancer external IP.
```yaml
type: LoadBalancer
externalIPs:
- <<node-ip-adress>>
```

## Update Configuration
```shell

```

## Remove Configuration
```shell

helm uninstall co-nginx --namespace co-ingress
helm uninstall co-nginx-internal --namespace co-ingress-internal
helm uninstall co-prometheus --namespace co-monitoring
helm uninstall co-kured --namespace co-maintenace

# delete all namspaces
kubectl delete ns co-ingress co-ingress-internal co-maintenance co-monitoring co-opa

#clean up gatekeeper crd
kubectl delete crd \
  configs.config.gatekeeper.sh \
  constraintpodstatuses.status.gatekeeper.sh \
  constrainttemplatepodstatuses.status.gatekeeper.sh \
  constrainttemplates.templates.gatekeeper.sh

```
