#!/usr/bin/env bash

export GOPATH=$HOME/goLibraries #The path of your GOPATH

$GOPATH/bin/gopherjs build realSketch.go -o sketch.js
