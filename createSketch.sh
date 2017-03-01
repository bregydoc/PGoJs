#!/usr/bin/env bash

export GOPATH=$HOME/goLibraries #The path of your GOPATH

$GOPATH/bin/gopherjs build $GOPATH/src/github.com/bregydoc/PGoJs/Examples/realSketch.go -o ./Examples/sketch.js
