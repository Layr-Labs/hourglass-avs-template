# syntax=docker/dockerfile:1.3

FROM golang:1.23.6-bookworm AS build

RUN apt-get update

WORKDIR /build

# Setup Git + known_hosts + env
RUN --mount=type=ssh git config --global url."git@github.com:".insteadOf "https://github.com/"
RUN mkdir -p ~/.ssh && ssh-keyscan github.com >> ~/.ssh/known_hosts
ENV GOPRIVATE=github.com/Layr-Labs/*
ENV GIT_SSH_COMMAND="ssh -o StrictHostKeyChecking=accept-new"


# Copy full source
ADD . /build

# Build with host ssh mounted
RUN --mount=type=ssh make build

FROM debian:stable-slim

COPY --from=build /build/bin/performer /usr/local/bin/performer

CMD ["/usr/local/bin/performer"]
