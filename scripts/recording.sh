#!/usr/bin/env bash

set -e

rm -rf ./resources/{*.cast,*.gif}

go clean -modcache
go clean -cache

asciinema rec \
  --title '[DEMO]: gocosi' \
  --command ./scripts/demo.sh \
  --cols 120 \
  --rows 20 \
  ./resources/cosi-demo.cast

agg ./resources/cosi-demo.cast ./resources/demo.gif
