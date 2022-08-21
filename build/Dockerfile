FROM golang:alpine AS base_img

RUN apk update && apk upgrade && apk add --no-cache git

WORKDIR /work
COPY ./ /work

RUN go build ./cmd/apiserver
RUN go build ./cmd/refresher
RUN go build ./cmd/tracker
RUN go build ./cmd/oidc
