# k8s-hello-world

Create docker image:
```
docker buildx build --platform linux/amd64,linux/arm64 -t golang-hello-world-server:0.3 .
```

Push the image to docker hub:
```
docker push yashpandit4u/golang-echo-server:0.3
```

To install k6 for load testing:
```
brew install k6
```
Also then run load testing using:
```
k6 run k6-load-test.js
```