FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go .

RUN CGO_ENABLED=0 GO111MODULE=on go build -o /main

FROM scratch

WORKDIR /app

COPY --from=builder main .

CMD ./main