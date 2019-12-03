FROM golang:1.13-buster
RUN apt-get update && apt-get -y install git build-essential libfuse-dev kmod && apt-get clean

WORKDIR /go/src/github.com/trevor403/random
COPY cmd cmd
COPY pkg pkg
COPY linux_device .

RUN go build -o linear.a -buildmode=c-archive github.com/trevor403/random/cmd/library
RUN gcc -Wall -o /bin/srandom_cuse character_device.c $(pkg-config fuse --cflags --libs) linear.a

ENTRYPOINT ["srandom_cuse"]