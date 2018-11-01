#!/usr/bin/env bash

if [ $# -eq 0 ]
  then
    cd ../../../build
fi

go get -v ../cmd/config
#go build -v ../cmd/config
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v ../cmd/config