# Random Device
Random Number Generator implementation

- In any language of your choosing, please build a clone of /dev/random.

## Picking an algorithm
There's a saying in Software Engineering:
> Don't implement your own Crypto.

I think this goes doubly for random sources. Therefore I decided to pick a known, proven algorithm for producing pseudo-random numbers.

I looked into Linear congruential generators as the seem to be a fairly simple and performant algorithm for producing a random sequence. However these generators fail in terms of predictability.

In a security setting, you don't want an actor to be able to derive your random values given the algorithm and seed values. And this is there the Permutation Linear congruential generator comes in. The PCG is able to produce better statistical random than LCG by rotating high bits to increase the period of the algorithm.

All that technical talk basically means that you get safer random numbers for use in cryptography.

## Basic Usage
You can try the generator by running `make run-basic`

Doing so produces output like the following:
```
go run ./cmd/gen-random
==== Get Random numbers (64bits):
[ 0 ]   6710224707072829034
[ 1 ]   10700652705963212999
[ 2 ]   15089469281957138201
[ 3 ]   5728352411276030744
[ 4 ]   13256527356990404865
[ 5 ]   555953303729224554
[ 6 ]   15052000189748938975
[ 7 ]   10686356575835518316
[ 8 ]   13631493465802562088
[ 9 ]   6229364558408494527
==== Get Random number (1-10):
[ 0 ]   6
[ 1 ]   1
[ 2 ]   4
[ 3 ]   4
[ 4 ]   3
[ 5 ]   9
[ 6 ]   9
[ 7 ]   1
[ 8 ]   3
[ 9 ]   2
```

## Taking the prompt too far
Or should I say taking the prompt too literally.
So it was cool to implement a random number generator from some white paper... But that wasn't really good enough...

So I decided to make a functional devfs device. The best way to do this without a kernel module was via FUSE, or more specifically, CUSE. That stands for Character Device in Userspace (which doesn't match up with it's acronym).

Using this linux subsystem I was able to write a userspace program that provides the functionality for a special character device `/dev/srandom`

## Difficulty
Unfortunately my random number generator was written in Golang and I could not find a library that would let me use CUSE in Go.

This is where I decided to export my Golang code as a C compatible library to allow my C code to call my Go functions. 

That functionality is provided by the `github.com/trevor403/random/cmd/library` main package.

## Docker Usage
By running `make docker` you will end up with a `/dev/srandom` device and a
randomness test will be run against the device. This algorithm passes the FIPS 140-2 test.

```
Starting random_srandom_cuse_1 ... done
+ cat /dev/srandom
+ rngtest --blockcount=1000
rngtest 2-unofficial-mt.14
Copyright (c) 2004 by Henrique de Moraes Holschuh
This is free software; see the source for copying conditions.  There is NO warranty; not even for MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

rngtest: starting FIPS tests...
rngtest: bits received from input: 20000032
rngtest: FIPS 140-2 successes: 1000
rngtest: FIPS 140-2 failures: 0
rngtest: FIPS 140-2(2001-10-10) Monobit: 0
rngtest: FIPS 140-2(2001-10-10) Poker: 0
rngtest: FIPS 140-2(2001-10-10) Runs: 0
rngtest: FIPS 140-2(2001-10-10) Long run: 0
rngtest: FIPS 140-2(2001-10-10) Continuous run: 0
rngtest: input channel speed: (min=829.282; avg=10097.134; max=19073.486)Mibits/s
rngtest: FIPS tests speed: (min=36.750; avg=53.405; max=58.508)Mibits/s
rngtest: Program run time: 359426 microseconds
+ exit 0
```

## Host Usage
By running `make docker-device` you will end up with a `/dev/srandom` device which is accessible from the host machine.

```
Building srandom_cuse
...

...
Emulated Character Device Running...
```

You can test it on the host machine by running `make test`

You can even use this device to seed your systems random. Open one terminal and run:
```
sudo rngd --rng-device=/dev/srandom --random-device=/dev/random --foreground
```
And in another terminal you can check your systems entropy:
```
watch -n 1 cat /proc/sys/kernel/random/entropy_avail
```