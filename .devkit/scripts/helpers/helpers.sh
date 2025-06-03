#!/usr/bin/env bash

# Function to log to stderr
log() {
    echo "$@" >&2
}

function ensureJq() {
    if ! command -v jq &> /dev/null; then
        log "Error: jq not found. Please run 'avs create' first."
        exit 1
    fi
}


function ensureYq() {
    if ! command -v yq &> /dev/null; then
        log "Error: yq not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureMake() {
    if ! command -v make &> /dev/null; then
        log "Error: make not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureDocker() {
    if ! command -v docker &> /dev/null; then
        log "Error: docker not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureRealpath() {
    if ! command -v realpath &> /dev/null; then
        log "Error: realpath not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureForge() {
    if ! command -v forge &> /dev/null; then
        log "Error: forge not found. Please run 'avs create' first."
        exit 1
    fi
}

function ensureGomplate() {
    if ! command -v gomplate &> /dev/null; then
        log "Error: gomplate not found. Please run 'avs create' first."
        exit 1
    fi
}

# Pass in RPC_URL ($1)
function ensureDockerHost() {
    # Detect OS and default DOCKERS_HOST when not provided
    if [[ "$(uname)" == "Linux" ]]; then
        # Lookup the host using iproute
        DOCKERS_HOST=${DOCKERS_HOST:-$(ip addr show docker0 | awk '/inet /{print $2}' | cut -d/ -f1)}
    else
        DOCKERS_HOST=${DOCKERS_HOST:-host.docker.internal}
    fi

    # Replace localhost/127.0.0.1 in RPC_URL with docker equivalent for environment
    DOCKER_RPC_URL=$(
    echo "$1" |
    sed -E \
        -e "s#(https?://)(localhost|127\.0\.0\.1)(:[0-9]+)?#\1${DOCKERS_HOST}\3#g"
    )

    # Return properly formed RPC url
    echo $DOCKER_RPC_URL
}
