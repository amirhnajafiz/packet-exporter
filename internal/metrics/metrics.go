package metrics

// Metrics sotres exporter prometheus metrics.
type Metrics struct{}

// IncRequest based on its source, dest, protocol, and ifname.
func (m *Metrics) IncRequest(src, dest, protocol, ifname string) {}

// ObserveThroughput based on its source, dest, protocol, ifname, and payload size.
func (m *Metrics) ObserveThroughput(src, dest, protocol, ifname string, payload float64) {}
