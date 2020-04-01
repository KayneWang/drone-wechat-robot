#!/bin/sh

GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -o release/linux/amd64/drone-wechat-robot
docker build --rm -t kaynewang/drone-wechat-robot .