#!/bin/bash

# set environment variables
export K8S_ONEREPO_VALUES='var/values/cluster01'
export K8S_ONEREPO_TEMPLATES='var/templates'
export K8S_ONEREPO_CONFIG='config/cluster'
export K8S_ONEREPO_HELM_TEMPLATES='var/helmcharts'
export K8S_ONEREPO_HELM_CONFIG='config/helmcharts'

# run binary
./bin/k8s-onerepo

