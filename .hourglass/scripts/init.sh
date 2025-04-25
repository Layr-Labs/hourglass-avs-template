#!/usr/bin/env bash

# -----------------------------------------------------------------------------
# Init will install os-level dependencies to build your hourglass AVS
#
# Currently it supports linux (ubuntu/debian) and macOS
# -----------------------------------------------------------------------------
if [[ "$OSTYPE" == "linux-gnu"* ]]; then
    # Linux
    # if not debian or ubuntu, exit with error message
    OS_NAME=$(source /etc/os-release && echo $NAME)

    if [[ "$OS_NAME" != *"Ubuntu"* && "$OS_NAME" != *"Debian"* ]]; then
        echo "This script is only supported on Ubuntu or Debian."
        exit 1
    fi

    echo "Installing dependencies for Linux..."
    apt-get update
    apt-get install -y curl jq

elif [[ "$OSTYPE" == "darwin"* ]]; then
    # MacOS
    echo "Installing dependencies for MacOS..."
    brew install jq grpcurl
else
    echo "Unsupported OS: $OSTYPE"
    exit 1
fi
