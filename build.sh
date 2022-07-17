#!/bin/bash

mkdir -p output
mkdir -p output/bin
cp -r conf output/
cp -r data output/
echo "copy conf,data yes!"
go build -o stockserver
echo "build,stockserver,yes"
mv stockserver output/bin/
