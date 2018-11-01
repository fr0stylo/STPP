#!/usr/bin/env bash

if [ $# -eq 0 ]
  then
    cd ./../../../api-gateway
  else
    cd ../api-gateway
fi

yarn install
yarn build
