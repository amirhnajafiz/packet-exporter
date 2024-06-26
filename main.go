package main

import (
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amirhnajafiz/packet-exporter/internal/model"
	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
	"github.com/cilium/ebpf/perf"
	"github.com/cilium/ebpf/rlimit"
	"github.com/vishvananda/netlink"
)

func main() {
	// Allow the current process to lock memory for eBPF maps
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatalf("failed to remove memlock limit: %v", err)
	}

	// Load pre-compiled programs into the kernel
	objs := struct {
		PacketMonitor *ebpf.Program `ebpf:"packet_monitor"`
		Events        *ebpf.Map     `ebpf:"events"`
	}{}
	spec, err := ebpf.LoadCollectionSpec("packet_filter.o")
	if err != nil {
		log.Fatalf("failed to load BPF program: %v", err)
	}
	if err := spec.LoadAndAssign(&objs, nil); err != nil {
		log.Fatalf("failed to load and assign BPF objects: %v", err)
	}
	defer objs.PacketMonitor.Close()
	defer objs.Events.Close()

	// Attach the program to all network interfaces
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("failed to list network interfaces: %v", err)
	}

	for _, link := range links {
		if err := attachXDP(link.Attrs().Index, objs.PacketMonitor); err != nil {
			log.Printf("failed to attach XDP to interface %s: %v", link.Attrs().Name, err)
		} else {
			log.Printf("attached XDP to interface %s", link.Attrs().Name)
		}
	}

	// Set up a perf reader to read packet events
	rd, err := perf.NewReader(objs.Events, os.Getpagesize())
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

			var pkt model.PacketMeta
			if err := binary.Read(record.RawSample, binary.LittleEndian, &pkt); err != nil {
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
