#!/bin/bash

minikube delete
minikube start
#minikube addons enable ingress

kubectl apply -f ./lab12/minikube/pv/main_pv.yaml
kubectl apply -f ./lab12/minikube/pv/city_pv.yaml

kubectl apply -f ./lab12/minikube/pv/main_pvc.yaml
kubectl apply -f ./lab12/minikube/pv/city_pvc.yaml

kubectl apply -f ./lab12/minikube/deployments/maindepl.yaml
kubectl apply -f ./lab12/minikube/deployments/citydepl.yaml 

kubectl apply -f ./lab12/minikube/services/main_service.yaml
kubectl apply -f ./lab12/minikube/services/city_service.yaml

#kubectl apply -f ./ingress/ingress.yaml

echo "doneee"


