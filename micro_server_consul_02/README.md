# go_micro - v3
## github 地址
https://github.com/asim/go-micro

## 概述
Go Micro 是一个分布式系统开发框架

在架构之外，它默认实现了consul作为服务发现（2019年源码修改了默认使用mdns），通过http进行通信，通过protobuf和json进行编解码

## 安装
```
# go-micro v3 版本
go get github.com/asim/go-micro
```

## micro 与 go-micro 关系
使用 go-micro 框架开始编写服务。运行时，使用 micro 运行时管理您的服务


```
go-micro共分为两块：
	- micro：使用 micro 运行时管理您的服务。如同 beego 框架中的bee
	- go-micro：就是微服务框架
```


