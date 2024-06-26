#include <uapi/linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/ip.h>
#include <linux/tcp.h>
#include <linux/udp.h>
#include <linux/bpf.h>
#include <bpf/bpf_helpers.h>

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

struct {
    __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
    __uint(max_entries, 1024);
} events SEC(".maps");

SEC("xdp")
int packet_monitor(struct xdp_md *ctx) {
    // allocate memory parts to read packet info
    void *data_end = (void *)(long)ctx->data_end;
    void *data = (void *)(long)ctx->data;
    struct ethhdr *eth = data;

    // drop if missing
    if ((void *)(eth + 1) > data_end) {
        return XDP_DROP;
    }

    struct iphdr *ip = data + sizeof(*eth);
    if ((void *)(ip + 1) > data_end) {
        return XDP_DROP;
    }

    // read ip data and insert inside a packet_info struct
    struct packet_info pkt = {};
    pkt.src_ip = ip->saddr;
    pkt.dest_ip = ip->daddr;
    pkt.protocol = ip->protocol;
    pkt.ifindex = ctx->ingress_ifindex;
    pkt.payload_len = data_end - data;

    if (ip->protocol == IPPROTO_TCP) {
        struct tcphdr *tcp = (void *)ip + sizeof(*ip);
        if ((void *)(tcp + 1) > data_end) {
            return XDP_DROP;
        }
        pkt.src_port = tcp->source;
        pkt.dest_port = tcp->dest;
    } else if (ip->protocol == IPPROTO_UDP) {
        struct udphdr *udp = (void *)ip + sizeof(*ip);
        if ((void *)(udp + 1) > data_end) {
            return XDP_DROP;
        }
        pkt.src_port = udp->source;
        pkt.dest_port = udp->dest;
    }

    bpf_perf_event_output(ctx, &events, BPF_F_CURRENT_CPU, &pkt, sizeof(pkt));
    return XDP_PASS;
}

char _license[] SEC("license") = "GPL";
