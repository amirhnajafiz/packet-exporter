package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type response struct {
	Requests     int       `json:"requests"`
	Logs         int       `json:"logs"`
	ResponseTime []float64 `json:"response_time"`
}

func pull(port, interval int) {
	// register metrics
	// make http request to metrics endpoint
	// update metrics
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":2112", nil); err != nil {
		panic(err)
	}
}
