#!/bin/bash

## Build line
docker run --rm -it \
    -v "$PWD"/build:/go/src/github.com/Yara-Rules/yago/build \
    -v "$PWD":/go/src/github.com/Yara-Rules/yago \
    -w /go/src/github.com/Yara-Rules/yago \
    go-builder \
    bash godep.sh

