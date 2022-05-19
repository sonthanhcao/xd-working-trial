#!/bin/bash
# APISERVER=https://kubernetes.default.svc.cluster.local
# export KUBERNETES_SERVICE_HOST=kubernetes.default.svc.cluster.local
# export KUBERNETES_SERVICE_PORT=443
TIMESTAMP=$(date +%s)
echo $IMAGE_TAG
if [[ -z $IMAGE_TAG ]]; then
  echo "IMAGE_TAG is empty"
  exit 1
fi


envsubst < config.yml > k8s-config.yml && kubectl apply -f k8s-config.yml -n $NAMESPACE
envsubst < deployment.yml > k8s-main.yml && kubectl apply -f k8s-main.yml -n $NAMESPACE

if [[ $? != 0 ]]; then exit 1; fi

kubectl rollout status deployments/$SERVICE_NAME -n $NAMESPACE

if [[ $? != 0 ]]; then
    kubectl logs $(kubectl get pods -n $NAMESPACE --sort-by=.metadata.creationTimestamp | grep "$SERVICE_NAME" | awk '{print $1}' | tac | head -1 ) -n $NAMESPACE --tail=20 && exit 1;
fi
