FROM golang:1.23.6-bookworm

RUN apt-get update

ADD . /build

WORKDIR /build

RUN make build

COMMAND ["/build/bin/performer"]
