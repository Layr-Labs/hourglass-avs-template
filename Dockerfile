FROM golang:1.23.6-bookworm AS build

RUN apt-get update

# -----------------------------------------------------------------------------
# Janky hack to get around the hourglass-monorepo module living in a private repo
#
COPY .hourglass/.docker-build-tmp /root/.docker-build-tmp

RUN cp -R /root/.docker-build-tmp/.ssh /root/.ssh
RUN cp /root/.docker-build-tmp/.gitconfig /root/.gitconfig || true

RUN chmod -R 700 /root/.ssh && \
    chmod 644 /root/.ssh/known_hosts
# -----------------------------------------------------------------------------

ADD . /build

WORKDIR /build

RUN make build

RUN rm -rf /root/.ssh

FROM debian:stable-slim

COPY --from=build /build/bin/performer /usr/local/bin/performer

CMD ["/usr/local/bin/performer"]
