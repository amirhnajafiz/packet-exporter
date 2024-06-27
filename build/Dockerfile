# Use an official Golang runtime as a parent image
FROM golang:1.22

# Install necessary packages
RUN apt-get update && apt-get install -y \
    clang \
    llvm \
    iproute2 \
    iputils-ping \
    libbpfcc-dev \
    linux-headers-$(uname -r) \
    make \
    gcc \
    git

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY /bpf /bpf
COPY /internal /internal
COPY main.go main.go
COPY go.mod go.sum ./

# Download Golang deps
RUN go mod download

# Build the eBPF program
RUN clang -O2 -target bpf -c /bpf/program.c -o /bpf/program.o

# Build the Go app
RUN go build -o exporter main.go

# Command to run the executable
CMD ["./exporter"]
