package worker

import (
	"github.com/amirhnajafiz/packet-exporter/internal/model"
	"github.com/amirhnajafiz/packet-exporter/internal/monitoring/logging"
	"github.com/amirhnajafiz/packet-exporter/internal/monitoring/metrics"
)

// New registers worker go-routines.
func New(limit int, channel chan *model.PacketMeta) {
	// create a new metrics instance
	metrics := metrics.New()

	// create a new logger instance
	logr := logging.New()

	// register workers
	for i := 0; i < limit; i++ {
		go worker{
			metrics: metrics,
			channel: channel,
			logr:    logr,
		}.work()
	}
}
