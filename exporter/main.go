package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type (
	response struct {
		Requests     int       `json:"requests"`
		Logs         int       `json:"logs"`
		ResponseTime []float64 `json:"response_time"`
	}

	metrics struct {
		Requests prometheus.Histogram
		Logs     prometheus.Histogram
		Response prometheus.Histogram
	}
)

func pull(namespace, subsystem string, port, interval int) error {
	// register metrics
	m := metrics{
		Requests: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "total_requests",
			Help:      "total number of service requests",
		}),
		Logs: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "logs_requests",
			Help:      "total number of service logs requests",
		}),
		Response: prometheus.NewHistogram(prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "response_time",
			Help:      "http server response time of each request",
		}),
	}

	// make http request to metrics endpoint
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("127.0.0.1:%d", port), nil)
	if err != nil {
		return err
	}

	for {
		client := &http.Client{}
		rsp, er := client.Do(request)
		if er != nil {
			log.Println(fmt.Errorf("server failed, error=%w", er))

			continue
		}

		responseInstance := new(response)
		if e := json.NewDecoder(rsp.Body).Decode(responseInstance); e != nil {
			log.Println(fmt.Errorf("parse response failed, error=%w", e))

			continue
		}

		m.Requests.Observe(float64(responseInstance.Requests))
		m.Logs.Observe(float64(responseInstance.Logs))

		for _, item := range responseInstance.ResponseTime {
			m.Requests.Observe(item)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func main() {
	http.Handle("/metrics", promhttp.Handler())

	if err := http.ListenAndServe(":2112", nil); err != nil {
		panic(err)
	}
}
