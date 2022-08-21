#!/usr/bin/env bash
set -e

docker build -t temp:latest . -f "${PWD}"/build/Dockerfile

docker build -t arpitjindal1997/apiserver:latest . -f "${PWD}"/build/apiserver/Dockerfile
docker build -t arpitjindal1997/refresher:latest . -f "${PWD}"/build/refresher/Dockerfile
docker build -t arpitjindal1997/tracker:latest . -f "${PWD}"/build/tracker/Dockerfile
docker build -t arpitjindal1997/oidc:latest . -f "${PWD}"/build/oidc/Dockerfile

docker push arpitjindal1997/apiserver:latest
docker push arpitjindal1997/refresher:latest
docker push arpitjindal1997/tracker:latest
docker push arpitjindal1997/oidc:latest
