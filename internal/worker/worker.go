package worker

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/packet-exporter/internal/metrics"
	"github.com/amirhnajafiz/packet-exporter/internal/model"

	"github.com/vishvananda/netlink"
)

// workers are processes used to handle the input packetmetas.
type worker struct {
	channel chan *model.PacketMeta
	metrics *metrics.Metrics
}

func (w worker) work() {
	// listen for packet events
	for pkt := range w.channel {
		link, err := netlink.LinkByIndex(int(pkt.IfIndex))
		if err != nil {
			log.Printf("failed to get interface name: %v", err)
			continue
		}
		ifaceName := link.Attrs().Name

		srcIP := fmt.Sprintf("%d.%d.%d.%d", pkt.SrcIP&0xff, (pkt.SrcIP>>8)&0xff, (pkt.SrcIP>>16)&0xff, (pkt.SrcIP>>24)&0xff)
		destIP := fmt.Sprintf("%d.%d.%d.%d", pkt.DestIP&0xff, (pkt.DestIP>>8)&0xff, (pkt.DestIP>>16)&0xff, (pkt.DestIP>>24)&0xff)
		log.Printf("Packet: Interface=%s, SrcIP=%s, DestIP=%s, SrcPort=%d, DestPort=%d, Protocol=%d, PayloadLen=%d",
			ifaceName, srcIP, destIP, pkt.SrcPort, pkt.DestPort, pkt.Protocol, pkt.PayloadLen)
	}
}
