## Carbonaut Deployment

### Deploy Carbonaut for dev

To get the things started localy you need a K8s like [kind](https://kind.sigs.k8s.io/). Create the kind cluster with the dev config `kind create cluster --config ./deployment/dev/kind-conf.yaml`. Ensure that you have the dependencies available via 'helm dependency build ./deployment'.

Next, get Mimir & Grafana started. 
```
kubectl create namespace carbonaut
kubectl create namespace mimir
helm upgrade carbonaut ./deployment --namespace carbonaut -i
```

To access Grafana (http://localhost:3000/) get the password and port-forward the dashbaord
```
kubectl port-forward service/carbonaut-grafana 3000:80 -n carbonaut
kubectl get secret -n carbonaut carbonaut-grafana -o jsonpath="{.data.admin-password}" | base64 --decode
```

Username: admin ;)

### Local testing with Prometheus

In addition to Carbonaut charts you can deploy locally Promethues for generating load and push it to Mimir. The testing directory contains the required config.

```
kubectl create ns prom
helm upgrade -f ./deployment/testing/values.yaml prometheus prometheus-community/kube-prometheus-stack -n prom -i
```