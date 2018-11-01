#!/usr/bin/env bash

if [ $# -eq 0 ]
  then
    cd ./../../../build/
fi

go get -v ../cmd/projects
#go build -v ../cmd/projects
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -v ../cmd/projects