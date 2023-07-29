#!/usr/bin/env bash

k8s_path=$(realpath "$(dirname "$0")"/../../.deploy/k8s)
pids=()

function error() {
  echo "$*" 2>&1 && exit 1

}

function usage() {
  echo -e \
"- start minikube and deploy services::
$ src/shared/minikube.sh start
- update services::
$ src/shared/minikube.sh update
- open a tunnel to minikube cluster and forward ports::
$ src/shared/minikube.sh open"
}

function start() {
  assumption-bin
  echo "> spinning minikube"
  minikube start
  deploy
}

function stop() {
  assumption-bin
  delete
  minikube stop
}

function deploy() {
  assumption-bin
  echo "> updating k8s"
  kubectl apply -f "$k8s_path"
}

function tear() {
  assumption-bin
  echo "> deleting k8s"
  kubectl delete -f "$k8s_path"
}

function open() {
  assumption-bin
  assumption-svc
  echo "> opening tunnel"
  trap "sigterm" TERM INT
  minikube tunnel &
  pids+=($!)
  sleep 2
  kubectl port-forward deployment/ml 5000:localhost:5000 &
  pids+=($!)
  kubectl port-forward deployment/s3 9000:localhost:9000 &
  pids+=($!)
  wait "${pids[0]}"
}

function sigterm() {
  echo "> terminating tunnel"
  kill -TERM "${pids[0]}"
}

function assumption-bin() {
  echo "> checking binaries"
  minikube version >/dev/null 2>/dev/null || error 'minikube' missing
  kubectl version --client=true >/dev/null 2>/dev/null || error 'kubectl' missing
}

function assumption-svc() {
  echo "> checking services(s)"
  kubectl get service ml s3 >/dev/null || error either 'ml' or 's3' service not running in the cluster
}


case "$1" in
  start) start;;
  deploy|update) deploy;;
  tear|delete) tear;;
  open) open;;
  stop) stop;;
  *) usage && exit 1;;
esac