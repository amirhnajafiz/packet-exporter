package worker

import (
	"github.com/amirhnajafiz/packet-exporter/internal/metrics"
	"github.com/amirhnajafiz/packet-exporter/internal/model"
)

// workers are processes used to handle the input packetmetas.
type worker struct {
	channel chan *model.PacketMeta
	metrics *metrics.Metrics
}

func (w worker) work() {

}
