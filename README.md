# Packet Exporter

This is an exporter implemented by __Golang__ and __eBFP__ to capture packets within your system's interfaces, and export their data as `prometheus` metrics.

Packet exporter contains two main programs, the first program is an eBPF application that is set to the system kernel in order to export packets data. The second program is a Golang application that uses eBPF libraries to fetch the exported packets and convert their data as prometheus metrics.

## BPF

Our eBPF program sets a function as xdp to the kernel which gets packet data from memory and returns a struct as follow:

```c
// packet_info is a struct that I defined in order to extract a packet data
// from memory. this struct will be returned to our eBPF program in Golang.
struct packet_info {
    __u32 src_ip;
    __u32 dest_ip;
    __u16 src_port;
    __u16 dest_port;
    __u8 protocol;
    __u32 ifindex;
    __u32 payload_len;
};
```

As you can see, it returns source and dest ip and port, packet protocol and interface with the payload's size. For each network event, it captures the packets and converts them to this struct.

## Exporter

The exporter has a `XDPManager` which attachs the following method to all system interfaces.

```go
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
```

As you can see, it reads packets data and publishs them over a channel to be exproted as prometheus metrics. Finally, the system workers read from this channel and publish metrics on prometheus http server.

## Envs

You can set the number of workers as `PE_WORKERS`, and the metrics port as `PE_PORT`. After that you can fetch the exporter metrics at `localhost:PE_PORT/metrics`.

## Metrics

Exporter returns two metrics, `total_packets` per each interface, and `total_throughput` per each interface and protocol.

## Run

You can run the export on your system using `docker-compose up -d`. If you are running the exporter on your local system, make sure to install `Golang`, `Clang`, `gcc`, and/or `libbpfcc-dev`, `libbpf`.
