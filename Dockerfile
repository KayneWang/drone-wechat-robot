FROM plugins/base:multiarch

LABEL maintainer="Kayne Wang <w.zengkai@foxmail.com>"

LABEL org.label-schema.version=latest
LABEL org.label-schema.vcs-url="https://github.com/KayneWang/drone-wechat-robot.git"
LABEL org.label-schema.name="Drone Wechat Robot"
LABEL org.label-schema.schema-version="1.0"

ADD release/linux/amd64/drone-wechat-robot /bin/
ENTRYPOINT ["/bin/drone-wechat-robot"]