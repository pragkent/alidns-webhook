#!/usr/bin/env bash

curl https://storage.googleapis.com/kubebuilder-tools/kubebuilder-tools-1.14.1-linux-amd64.tar.gz -o kubebuilder-tools.tar.gz
tar -zvxf kubebuilder-tools.tar.gz

rm kubebuilder-tools.tar.gz
