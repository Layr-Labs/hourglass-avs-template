FROM golang:1.23.6-bookworm AS build

WORKDIR /build

# Copy full source
ADD . /build

RUN make build

FROM debian:stable-slim

COPY --from=build /build/bin/performer /usr/local/bin/performer

RUN apt-get update \
  && apt-get install -y --no-install-recommends ca-certificates \
  && rm -rf /var/lib/apt/lists/*

CMD ["/usr/local/bin/performer"]
