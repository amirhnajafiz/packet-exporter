# Traffic Monitor

Traffic monitor is an eBPF based cloud-native application that monitors input/output network traffic of a system (specifically, Kubernetes underlayer nodes). This application exports network charactristics of a node to be used by network agents (troubleshooting, performance analysis, traffic analysis, security, monitoring, etc.).

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
