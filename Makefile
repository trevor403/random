SHELL := /bin/bash

.PHONY: all build docker

all: build

docker: docker-build docker-run

driver:
	sudo modprobe cuse

build:
	mkdir -p bin/ build/
	go build -o build/linear.a -buildmode=c-archive github.com/trevor403/random/cmd/library
	gcc -Wall -o bin/srandom_cuse linux_device/character_device.c -Ibuild $(shell pkg-config fuse --cflags --libs) build/linear.a

run: build driver
	sudo ./bin/srandom_cuse

test:
	sudo cat /dev/srandom | rngtest --blockcount=1000 || echo -n

run-basic:
	go run ./cmd/gen-random

clean:
	rm -f build/linear.a build/linear.h bin/srandom_cuse

docker-build:
	docker-compose build

docker-run: driver
	docker-compose run --rm srandom_test
	docker-compose down

docker-device: driver
	docker-compose down
	docker-compose run --rm srandom_cuse