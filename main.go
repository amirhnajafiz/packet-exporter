package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amirhnajafiz/packet-exporter/internal/model"
	"github.com/amirhnajafiz/packet-exporter/internal/xdp"
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
	"github.com/cilium/ebpf/rlimit"
	"github.com/vishvananda/netlink"
)

func main() {
	// allow the current process to lock memory for eBPF maps
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("failed to remove memlock limit: %v", err)
	}

	// create a new xdp manager
	mgr, err := xdp.New("bpf/program.o")
	if err != nil {
		log.Fatalf("failed to create new xdp manager: %v\n", err)
	}

	defer mgr.PacketMonitor.Close()
	defer mgr.Events.Close()

	// Attach the program to all network interfaces
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("failed to list network interfaces: %v", err)
	}

	for _, link := range links {
		if err := attachXDP(link.Attrs().Index, mgr.PacketMonitor); err != nil {
			log.Printf("failed to attach XDP to interface %s: %v", link.Attrs().Name, err)
		} else {
			log.Printf("attached XDP to interface %s", link.Attrs().Name)
		}
	}

	// Set up a perf reader to read packet events
	rd, err := perf.NewReader(mgr.Events, os.Getpagesize())
	if err != nil {
		log.Fatalf("failed to create perf reader: %v", err)
	}
	defer rd.Close()

	go func() {
		for {
			record, err := rd.Read()
			if err != nil {
				log.Fatalf("failed to read from perf reader: %v", err)
			}

			if record.LostSamples != 0 {
				log.Printf("lost %d samples", record.LostSamples)
				continue
			}

			reader := bytes.NewReader(record.RawSample)

			var pkt model.PacketMeta
			if err := binary.Read(reader, binary.LittleEndian, &pkt); err != nil {
				log.Printf("failed to decode received data: %v", err)
				continue
			}

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

			_ = &model.Payload{
				Src:           srcIP,
				Dest:          destIP,
				Protocol:      pkt.Protocol,
				InterfaceName: ifaceName,
				PayloadLen:    pkt.PayloadLen,
			}
		}
	}()

	// Listen for program termination signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	log.Println("exiting...")
}

func attachXDP(iface int, prog *ebpf.Program) error {
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   prog,
		Interface: iface,
		Flags:     link.XDPGenericMode,
	})
	if err != nil {
		return err
	}
	defer link.Close()
	return nil
}
