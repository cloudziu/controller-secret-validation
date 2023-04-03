# Kubernetes secret validation webhook build with controller-runtime

README: WIP

### Invalid request
kubectl create secret generic test --from-literal=manifest=invalid

### Valid reqiest
kubectl create secret generic test --from-literal=manifest=valid

