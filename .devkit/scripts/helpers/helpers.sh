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
