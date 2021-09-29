# Go mod 依赖管理
- Go.mod是Golang1.11版本新引入的官方包管理工具用于解决之前没有地方记录依赖包具体版本的问题，方便依赖包的管理。
- 设置GO111MODULE
  - GO111MODULE有三个值：off, on和auto（默认值）。
  - window : go 
  - linux: export GO111MODULE=on
- go mod命令

查看 gin历史版本
go list -m -versions github.com/gin-gonic/gin
更换版本
go mod edit -require="github.com/gin-gonic/gin@v1.6.0"  //@后是版本号
go mod tidy
查看项目依赖的包
go list -m all

# go 代理地址设置
目前发现的几个不错的goproxy
阿里云
配置如下：

export GOPROXY=https://mirrors.aliyun.com/goproxy/
nexus社区提供的
配置如下：
export GOPROXY=https://gonexus.dev
goproxy.io 的
配置如下：
export GOPROXY=https://goproxy.io/
基于athens的公共服务
配置如下：
export GOPROXY=https://athens.azurefd.net

官方提供的(jfrog,golang)
export GOPROXY=https://gocenter.io
export GOPROXY=https://proxy.golang.org

七牛云赞助支持的
export GOPROXY=https://goproxy.cn

- 终端执行
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/

go env -w  GOPROXY=https://proxy.golang.org,direct

```
go get github.com/go-redis/redis: module github.com/go-redis/redis: Get https://proxy.golang.org/github.com/go-redis/redis/@v/list: dial tcp 142.251.42.241:443: connectex: A c
onnection attempt failed because the connected party did not properly respond after a period of time, or established connection failed because connected host has failed to res
pond.
```
处理：
```
go env -w GOPROXY=https://goproxy.cn
```