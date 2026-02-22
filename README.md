# k8s-hello-world

Create docker image:
```
docker buildx build --platform linux/amd64,linux/arm64 -t golang-hello-world-server:0.4 .
```

Push the image to docker hub:
```
docker push yashpandit4u/golang-echo-server:0.4
```

To install k6 for load testing:
```
brew install k6
```
Also then run load testing using:
```
k6 run k6-load-test.js
```

Running prometheus locally:
```
docker run -d -p 9090:9090 -v $(pwd)/prometheus.yml:/etc/prometheus/prometheus.yml prom/prometheus
```

Running grafana locally:
```
docker run -d -p 3000:3000 --name=grafana grafana/grafana
```