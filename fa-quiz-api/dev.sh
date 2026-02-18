#!/usr/bin/env bash

# Locate and navigate to  the directory of this start.sh file.
BASEDIR=$(dirname "$0")
cd $BASEDIR

# Serve the API module
AUTOMIGRATE=1 go run src/main.go serve
