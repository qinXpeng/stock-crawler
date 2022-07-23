#!/bin/bash

mkdir -p output
mkdir -p output/bin
cp -r conf output/
echo "copy conf,data yes!"
go build -o stock-crawler
echo "build,stockserver,yes"
mv stock-crawler output/bin/
