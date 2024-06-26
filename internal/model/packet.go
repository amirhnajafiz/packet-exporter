package model

// PacketMeta is used to fetch returned data from bpf/program.o
// into a struct where it will be used in the program's exporter.
type PacketMeta struct {
	SrcIP      uint32
	DestIP     uint32
	SrcPort    uint16
	DestPort   uint16
	Protocol   uint8
	IfIndex    uint32
	PayloadLen uint32
}
