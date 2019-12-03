#!/bin/sh
set -x

cat /dev/srandom | rngtest --blockcount=1000

exit 0