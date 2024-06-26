package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/amirhnajafiz/packet-exporter/internal/worker"
	"github.com/amirhnajafiz/packet-exporter/internal/xdp"
	"github.com/cilium/ebpf/rlimit"
	"github.com/vishvananda/netlink"
)

func main() {
	// listen for program termination signals
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

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

	// attach the program to all network interfaces
	links, err := netlink.LinkList()
	if err != nil {
		log.Fatalf("failed to list network interfaces: %v", err)
	}

	// NOTE: change this section to select your give interfaces
	for _, link := range links {
		if err := mgr.Attach(link.Attrs().Name, link.Attrs().Index); err != nil {
			log.Printf("failed to attach XDP to interface: %v", err)
		} else {
			log.Printf("attached XDP to interface %s", link.Attrs().Name)
		}
	}

	// run manager reader method
	channel, err := mgr.Reader()
	if err != nil {
		log.Fatalf("failed to start manager reader: %v\n", err)
	}

	// create a new pool to process packetmetas into prometheus metrics
	worker.New(5, channel)

	// wait for termination signal
	<-sig
	log.Println("exiting...")
}
