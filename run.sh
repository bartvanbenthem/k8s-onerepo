#!/bin/bash

# set generic environment variables
export K8S_ONEREPO_TEMPLATES='var/templates'
export K8S_ONEREPO_HELM_TEMPLATES='var/helmcharts'
export K8S_ONEREPO_HELM_CONFIG='config/helmcharts'

# set cluster01 specific environment variables
export K8S_ONEREPO_VALUES='var/values/cluster01'
export K8S_ONEREPO_CONFIG='config/cluster01'
# run onerepo gen
./bin/k8s-onerepo

# set cluster02 environment variables
export K8S_ONEREPO_VALUES='var/values/cluster02'
export K8S_ONEREPO_CONFIG='config/cluster02'
# run onerepo gen
./bin/k8s-onerepo

