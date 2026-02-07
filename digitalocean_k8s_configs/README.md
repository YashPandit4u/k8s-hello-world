Install these packages using:

```
kubectl apply -f .
```

After this install the nginx ingress controller:

```
helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
helm install ingress-nginx ingress-nginx/ingress-nginx --create-namespace --namespace ingress-nginx
```
