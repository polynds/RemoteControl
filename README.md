# RemoteControl

## 功能点
- 实现键盘方向键控制主机
- 实现远程屏幕查看和控制
    - 基于流媒体的屏幕显示
- 实现远程触控板
- 基于websocket实现指令通讯

## 运行
- Mac需要高权限才能运行
```shell
sudo ./RemoteControl-darwin-amd64
```


## 编译
- 需要安装gcc
- Windows需要安装mingw
- https://juejin.cn/post/6844903944808824845
- Mac 下编译， Linux 或者 Windows 下去执行
```shell
# linux 下去执行
CGO_ENABLED=1  GOOS=linux  GOARCH=amd64  go build main.go
# Windows 下去执行
CGO_ENABLED=1 GOOS=windows  GOARCH=amd64  go  build  main.go
```
- Linux 下编译 ， Mac 或者 Windows 下去执行
```shell
# Mac  下去执行
CGO_ENABLED=1 GOOS=darwin  GOARCH=amd64  go build main.go
# Windows 下执行
CGO_ENABLED=1 GOOS=windows  GOARCH=amd64  go build main.go
```
- Windows 下执行 ， Mac 或 Linux 下去执行
```shell
# Mac 下执行
SET  CGO_ENABLED=1
SET GOOS=darwin
SET GOARCH=amd64
go build main.go

# Linux 去执行
SET CGO_ENABLED=1
SET GOOS=linux
SET GOARCH=amd64
go build main.go

```