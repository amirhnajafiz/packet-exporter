package worker

import (
	"github.com/amirhnajafiz/packet-exporter/internal/metrics"
	"github.com/amirhnajafiz/packet-exporter/internal/model"
)

// New registers worker go-routines.
func New(limit int, channel chan *model.PacketMeta) {
	// create a new metrics instance
	metrics := metrics.New()

	// register workers
	for i := 0; i < limit; i++ {
		go worker{
			metrics: metrics,
			channel: channel,
		}.work()
	}
}
