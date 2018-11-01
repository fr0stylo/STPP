#!/usr/bin/env bash

if [ $# -eq 0 ]
  then
    cd ./../../../build/
fi

go get -v ../cmd/time-entries
#go build -v ../cmd/time-entries

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v ../cmd/time-entries