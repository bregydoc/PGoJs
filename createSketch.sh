#!/usr/bin/env bash

export GOPATH=$HOME/goLibraries

$GOPATH/bin/gopherjs build realSketch.go -o sketch.js
