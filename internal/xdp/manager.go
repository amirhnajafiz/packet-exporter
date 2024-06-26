package xdp

import (
	"fmt"

	"github.com/cilium/ebpf"
)

// XDPManager handles eBPF program and events.
type XDPManager struct {
	PacketMonitor *ebpf.Program `ebpf:"packet_monitor"`
	Events        *ebpf.Map     `ebpf:"events"`
}

func New(program string) (*XDPManager, error) {
	// load pre-compiled programs into the kernel
	objs := &XDPManager{}

	spec, err := ebpf.LoadCollectionSpec(program)
	if err != nil {
		return nil, fmt.Errorf("failed to load BPF program: %v", err)
	}

	if err := spec.LoadAndAssign(&objs, nil); err != nil {
		return nil, fmt.Errorf("failed to load and assign BPF objects: %v", err)
	}

	return objs, nil
}
