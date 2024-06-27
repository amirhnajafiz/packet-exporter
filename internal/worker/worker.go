package worker

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/packet-exporter/internal/model"
	"github.com/amirhnajafiz/packet-exporter/internal/monitoring/metrics"

	"github.com/vishvananda/netlink"
)

// workers are processes used to handle the input packetmetas.
type worker struct {
	channel chan *model.PacketMeta
	metrics *metrics.Metrics
}

func (w worker) work() {
	// listen for packetmeta events over the give channel by XDP reader
	for pkt := range w.channel {
		// find interface
		link, err := netlink.LinkByIndex(int(pkt.IfIndex))
		if err != nil {
			log.Printf("failed to get interface name: %v", err)
			continue
		}

		// export data
		ifaceName := link.Attrs().Name
		srcIP := fmt.Sprintf("%d.%d.%d.%d", pkt.SrcIP&0xff, (pkt.SrcIP>>8)&0xff, (pkt.SrcIP>>16)&0xff, (pkt.SrcIP>>24)&0xff)
		destIP := fmt.Sprintf("%d.%d.%d.%d", pkt.DestIP&0xff, (pkt.DestIP>>8)&0xff, (pkt.DestIP>>16)&0xff, (pkt.DestIP>>24)&0xff)
		src := fmt.Sprintf("%s:%d", srcIP, pkt.SrcPort)
		dest := fmt.Sprintf("%s:%d", destIP, pkt.DestPort)

		log.Printf(
			"Packet: Interface=%s, SrcIP=%s, DestIP=%s, SrcPort=%d, DestPort=%d, Protocol=%d, PayloadLen=%d",
			ifaceName, srcIP, destIP, pkt.SrcPort, pkt.DestPort, pkt.Protocol, pkt.PayloadLen,
		)

		// export metrics
		w.metrics.IncRequest(src, dest, ifaceName, int(pkt.Protocol))
		w.metrics.ObserveThroughput(src, dest, ifaceName, int(pkt.Protocol), float64(pkt.PayloadLen))
	}
}
