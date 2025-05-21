# syntax=docker/dockerfile:1.3

FROM golang:1.23.6-bookworm AS build

RUN apt-get update

WORKDIR /build

# Copy full source
ADD . /build

RUN make build

FROM debian:stable-slim

COPY --from=build /build/bin/performer /usr/local/bin/performer

CMD ["/usr/local/bin/performer"]
