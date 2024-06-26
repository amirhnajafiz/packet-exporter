package xdp

import (
	"fmt"

	"github.com/cilium/ebpf/link"
)

// Attach adds the XDP to the given interface.
func (x *XDPManager) Attach(ifname string, iface int) error {
	link, err := link.AttachXDP(link.XDPOptions{
		Program:   x.PacketMonitor,
		Interface: iface,
		Flags:     link.XDPGenericMode,
	})
	if err != nil {
		return fmt.Errorf("failed to attach XDP to interface %s (index: %d): %v", ifname, iface, err)
	}

	defer link.Close()

	return nil
}
