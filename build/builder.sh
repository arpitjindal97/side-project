#!/usr/bin/env bash
set -e

docker build -t temp:latest . -f "${PWD}"/build/Dockerfile

docker build -t arpitjindal1997/apiserver:latest . -f "${PWD}"/build/apiserver/Dockerfile
docker build -t arpitjindal1997/refresher:latest . -f "${PWD}"/build/refresher/Dockerfile
docker push arpitjindal1997/apiserver:latest
docker push arpitjindal1997/refresher:latest