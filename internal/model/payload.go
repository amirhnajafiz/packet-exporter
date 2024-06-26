package model

// Payload is a struct which stores a packet data in a
// format that is going to get exported as a prometheus metric.
type Payload struct {
	Src           string
	Dest          string
	Protocol      uint8
	InterfaceName string
	PayloadLen    uint32
}
