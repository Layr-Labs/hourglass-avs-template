# syntax=docker/dockerfile:1.3

FROM golang:1.23.6-bookworm AS build

RUN apt-get update

WORKDIR /build

# Configure git to use HTTPS instead of SSH
RUN git config --global url."https://github.com/".insteadOf "git@github.com:" && \
    git config --global url."https://".insteadOf "ssh://" && \
    git config --global url."https://github.com/".insteadOf "ssh://git@github.com/"

# Setup GitHub token authentication if provided
ARG GITHUB_TOKEN
RUN if [ -n "$GITHUB_TOKEN" ]; then \
        git config --global url."https://${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/"; \
    fi

# Set GOPROXY to use public proxy first, then direct
ENV GOPROXY=https://proxy.golang.org,direct

# Copy full source
ADD . /build

RUN make build

FROM debian:stable-slim

COPY --from=build /build/bin/performer /usr/local/bin/performer

CMD ["/usr/local/bin/performer"]
