#!/usr/bin/env bash
function cleanUp {
    docker-compose down
}

trap cleanUp EXIT

cd ./build

./build.sh

cd ../docker/

docker-compose up --build
