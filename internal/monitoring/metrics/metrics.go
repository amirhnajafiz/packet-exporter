package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Metrics sotres exporter prometheus metrics.
type Metrics struct {
	requests   *prometheus.CounterVec
	throughput *prometheus.HistogramVec
}

// New returns a new metrics instance.
func New() *Metrics {
	labels := []string{
		"source",
		"dest",
		"interface",
		"protocol",
	}

	return &Metrics{
		requests: promauto.NewCounterVec(prometheus.CounterOpts{
			Name: "total_packets",
			Help: "counting total packets of an interface",
		}, labels),
		throughput: promauto.NewHistogramVec(prometheus.HistogramOpts{
			Name: "total_throughput",
			Help: "observe the total payload that is being transfer through an interface",
		}, labels),
	}
}

// IncRequest based on its source, dest, protocol, and ifname.
func (m *Metrics) IncRequest(src, dest, ifname string, protocol int) {}

// ObserveThroughput based on its source, dest, protocol, ifname, and payload size.
func (m *Metrics) ObserveThroughput(src, dest, ifname string, protocol int, payload float64) {}
