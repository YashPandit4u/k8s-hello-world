package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	httpRequests = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests.",
		},
		[]string{"method", "path"},
	)

	inFlight = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "http_requests_in_flight",
		Help: "Number of HTTP requests currently in flight.",
	})
)

func main() {
	fmt.Println("Hello World")

	prometheus.MustRegister(httpRequests)
	prometheus.MustRegister(inFlight)

	// Main server path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		inFlight.Inc()
		defer inFlight.Dec()
		httpRequests.WithLabelValues(r.Method, "/").Inc()

		returnString := "Your request was received on: "
		currTime := time.Now()
		w.Write([]byte(returnString))
		w.Write([]byte(currTime.Format(time.RFC3339)))
		fmt.Println("Received request in golang server.")
	})

	// Prometheus metrics path
	http.Handle("/metrics", promhttp.Handler())

	fmt.Println("Listening to 8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error: ", err)
	}
}
