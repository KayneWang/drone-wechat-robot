FROM golang:1.14.1-alpine AS builder
ADD . /build
WORKDIR /build

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod=vendor -ldflags "-s -w" -a -o release/linux/amd64/drone-wechat-robot


FROM plugins/base:multiarch

LABEL maintainer="Kayne Wang <w.zengkai@foxmail.com>"

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/KayneWang/drone-wechat-robot.git"
LABEL org.label-schema.name="Drone Wechat Robot"
LABEL org.label-schema.schema-version="1.0"

COPY --from=builder /build/release/linux/amd64/drone-wechat-robot /bin/
ENTRYPOINT ["/bin/drone-wechat-robot"]
