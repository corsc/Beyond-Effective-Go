#!/bin/bash

if [ -z "$1" ]; then
    echo "Please supply a destination directory"
	exit -1
fi

# Ensure the inputted director is in a predictable format
DIR=${1%...}
PKG_DIR=${DIR%/}/

DIRS=$(find $PKG_DIR -type f -name '*.go' -not -path "*/vendor/*")

# Clean and simplify the code
gofmt -w -s -l $DIRS
