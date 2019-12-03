FROM golang:1.13-buster
RUN apt-get update && apt-get -y --no-install-recommends install git build-essential libfuse-dev kmod && rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/github.com/trevor403/random
COPY cmd cmd
COPY pkg pkg
COPY linux_device .

RUN go build -o linear.a -buildmode=c-archive github.com/trevor403/random/cmd/library
RUN gcc -Wall -o /bin/srandom_cuse character_device.c $(pkg-config fuse --cflags --libs) linear.a

ENTRYPOINT ["srandom_cuse"]