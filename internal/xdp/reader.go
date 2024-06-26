package xdp

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"

	"github.com/amirhnajafiz/packet-exporter/internal/model"
	"github.com/cilium/ebpf/perf"
	"github.com/vishvananda/netlink"
)

// Reader sets up a perf read to read packet events and publish them
// as model.PacketMeta over a channel.
func (x *XDPManager) Reader() (chan *model.PacketMeta, error) {
	// create a channel for return packetmeta objects
	channel := make(chan *model.PacketMeta)

	// set up a perf reader to read packet events
	rd, err := perf.NewReader(x.Events, os.Getpagesize())
	if err != nil {
		return channel, fmt.Errorf("failed to create perf reader: %v", err)
	}

	go func() {
		// close rd after goroutine is closed
		defer rd.Close()

		// manager loop to read packetmeta objects
		for {
			// wait until buffered is full with new data
			record, err := rd.Read()
			if err != nil {
				log.Fatalf("failed to read from perf reader: %v", err)
			}

			if record.LostSamples != 0 {
				log.Printf("lost %d samples", record.LostSamples)
				continue
			}

			// create a new reader
			reader := bytes.NewReader(record.RawSample)

			// bind bytes data to packetmeta
			var pkt model.PacketMeta
			if err := binary.Read(reader, binary.LittleEndian, &pkt); err != nil {
				log.Printf("failed to decode received data: %v", err)
				continue
			}

			// publish over this channel
			channel <- &pkt

			// get link by its index
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
	}()

	return channel, nil
}
