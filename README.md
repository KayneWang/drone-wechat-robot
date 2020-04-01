## drone-wechat-robot

The WeChat for Work Robot notifications for Drone plugin. You can reference [the doc](https://work.weixin.qq.com/help?doc_id=13376) to know how to use it.

## Build

```shell
$ sh build.sh
```

## Environment Reference

You can look at drone pipelines [doc](https://docs.drone.io/pipeline/environment/reference/)

## Usage

```yml
kind: pipeline
type: docker
name: default

steps:
- name: notify
  image: kaynewang/drone-wechat-robot
  settings:
    msgtype: text
    key: your own robot key
    content: hello world
    mentioned_list: @all,kaynewang
    mentioned_mobile_list: @all,kaynewang
```

```yml
kind: pipeline
type: docker
name: default

steps:
- name: notify
  image: kaynewang/drone-wechat-robot
  settings:
    msgtype: markdown
    key: your own robot key
    content: "实时新增用户反馈<font color=\"warning\">132例</font>，请相关同事注意。\n
         >类型:<font color=\"comment\">用户反馈</font>
         >普通用户反馈:<font color=\"comment\">117例</font>
         >VIP用户反馈:<font color=\"comment\">15例</font>"
```