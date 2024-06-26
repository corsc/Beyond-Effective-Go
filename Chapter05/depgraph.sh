#!/bin/bash

set -o errexit
set -o pipefail

# Note:
# This script should be run in the base directory of the project/service

# Required Tools
# go get github.com/kisielk/godepgraph
# brew install graphviz

# Inputs
#
# This cuts down on typing by allowing you to enter only the sub-directory you wish to graph, instead of the entire
# package
prefix="./"
PKG=${1#$prefix}

# Constants
#
# Save the file in the home directory (so it's easy to find)
DEST_FILE=~/depgraph.png

# Calculate the package in the current directory and assume this is the base or project package
BASE_PKG=$(go list)
BASE_PKG_DELIMITED=$(echo $BASE_PKG | sed 's/\//\\\//g')

EXCLUSIONS="$EXCLUSIONS${GO_BASE_PKG}/vendor"

# Generate the dependency graph
godepgraph -s \
  -p "$EXCLUSIONS" \
  -o "$BASE_PKG" \
  $BASE_PKG/${PKG} |
  sed "s/$BASE_PKG_DELIMITED//g" | dot -Gsplines=true -Tpng -o $DEST_FILE

# Open the result in the default web browser (only works on OS X)
open $DEST_FILE
