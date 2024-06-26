# GIM 服务端（golang）
基于go+websocket实现的即时通讯系统

## 项目简介

GIM 是一个网页版即时聊天系统，界面简约、美观、操作简单且容易进行二次开发，更加适合于开发人员进行技术交流，摒弃了微信、QQ等一些大众化的功能，更加注重于交流学习。

##### 使用技术
##### 1、后端

- Golang 1.21+
- MySQL 5.7+
- Redis 5.0+
- Minio
- Nsq
##### 2、前端

- Vue3
- Naive UI
- Ts
##### 功能介绍

- 支持基本的注册登录
- 支持常见的增删改查
- 支持 WebSocket 通信
- 支持私聊及群聊以及房间聊天场景
- 支持服务水平扩展
- 支持聊天消息类型有 文本、代码块、图片及其它类型文件
- 支持聊天消息撤回、删除或批量删除、转发消息（逐条转发、合并转发）及群投票功能
- 支持编写个人笔记
- 支持绑定远程代码仓库
- 支持容器化部署
- 支持k8s容器编排

[查看前端代码](https://github.com/Ci-77/GIM)


## 项目安装

1. 下载源码

```git
$ git clone https://github.com/Ci-77/Go-chat.git
```

1. 拷贝项目根目录下 config.example.yaml 文件为 config.yaml 并正确配置相关参数

``` bash
$ cp config.example.yaml config.yaml # 请务必正确配置相关参数
```

3. 安装依赖包

``` bash
$ go mod tidy
```

4. 安装相关依赖命令行工具

``` bash
$ make install
```

5. 初始化数据库

``` bash
$ go run ./cmd/gim migrate
```

6. 开发环境下启动服务

``` bash
# 打开两个终端，分别运行下面两个命令

$ go run ./cmd/gim http      # 本地启动 http 服务
$ go run ./cmd/gim commet    # 本地启动 websocket 服务
```

7. 编译后运行

``` bash
$ make build                   # 执行编译命令

# 执行后可在 ./bin 目录下看到 lumenim
```
