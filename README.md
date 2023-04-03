# Controller validation webhook 
Kubernetes secret validation controller, created with sigs.k8s.io/controller-runtime package.

README: WIP

### Invalid request
kubectl create secret generic test --from-literal=manifest=invalid

### Valid reqiest
kubectl create secret generic test --from-literal=manifest=valid

