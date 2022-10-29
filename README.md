# Carbonaut

![carbonaut-banner](./assets/carbonaut-banner.png)

[![GitHub go.mod Go version of a Go module](https://img.shields.io/github/go-mod/go-version/carbonaut-cloud/carbonaut.svg)](https://github.com/carbonaut-cloud/carbonaut)
[![Go Report Card](https://goreportcard.com/badge/carbonaut-cloud/carbonaut)](https://goreportcard.com/report/carbonaut-cloud/carbonaut)
[![Coverage Status](https://coveralls.io/repos/github/carbonaut-cloud/carbonaut/badge.svg?branch=main)](https://coveralls.io/github/carbonaut-cloud/carbonaut?branch=main)
[![Slack](https://img.shields.io/badge/Slack-%23general-blueviolet)](https://carbonautgroup.slack.com/archives/C03B9P2T3AB)
[![CircleCI](https://circleci.com/gh/carbonaut-cloud/carbonaut/tree/main.svg?style=svg)](https://circleci.com/gh/carbonaut-cloud/carbonaut/tree/main)

**🚧 PROJECT UNDER DEVELOPMENT AND NOT READY TO USE 🚧**

Carbonaut is a open source tool to measure your carbon emissions, analyze your resource consumptions and support you in optimizing your green house gas footprint.

Carbonaut targets any ICT infrastructure, also in the first phases of development public cloud provider and IaaS provider are the main target. The system will also integrate with Kubernetes and other tools on the market which are able to manage and predict resource utilization.

Our target is to provide precises insights which are not based on estimations (where possible).


## Deploy Carbonaut for dev

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