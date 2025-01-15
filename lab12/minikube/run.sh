#!/bin/bash

minikube delete
minikube start

kubectl apply -f ./pv/main_pv.yaml
kubectl apply -f ./pv/city_pv.yaml

kubectl apply -f ./pv/main_pvc.yaml
kubectl apply -f ./pv/city_pvc.yaml

kubectl apply -f ./deployments/maindbdepl.yaml
kubectl apply -f ./deployments/citydbdepl.yaml 
kubectl apply -f ./deployments/maindepl.yaml
kubectl apply -f ./deployments/citydepl.yaml 

kubectl apply -f ./services/maindb_service.yaml
kubectl apply -f ./services/citydb_service.yaml
kubectl apply -f ./services/main_service.yaml
kubectl apply -f ./services/city_service.yaml


echo "doneee"


