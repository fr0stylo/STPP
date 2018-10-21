#!/usr/bin/env bash
cd ./build

./build.sh

cd ../docker/

docker-compose up --build -d
