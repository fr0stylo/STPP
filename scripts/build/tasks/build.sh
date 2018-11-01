#!/usr/bin/env bash

if [ $# -eq 0 ]
  then
    cd ./../../../build/
fi

go get -v ../cmd/tasks
#go build -v ../cmd/tasks
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v ../cmd/tasks