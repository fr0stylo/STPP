#!/usr/bin/env bash

cd ../../build

echo "Config service build"
./../scripts/build/config/build.sh 1
echo "Projects service build"
./../scripts/build/projects/build.sh 1
echo "Tasks service build"
./../scripts/build/tasks/build.sh 1
echo "Time entries service build"
./../scripts/build/time-entries/build.sh 1
echo "Api gateway service build"
./../scripts/build/api-gateway/build.sh 1
