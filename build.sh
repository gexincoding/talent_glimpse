#!/bin/zsh

RUN_NAME="talent.glimpse"

rm -rf output

mkdir -p output output/conf

find conf -type f | xargs -I{} cp {} ./output/conf/

go build -o ./output/${RUN_NAME} main.go
