package metrics

import (
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// NewServer runs a prometheus metrics exporter
// on the given port.
func NewServer(port int) {
	go func() {
		http.Handle("/metrics", promhttp.Handler())

		if err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil); err != nil {
			log.Fatal(err)
		}
	}()
}
