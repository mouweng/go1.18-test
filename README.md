# go1.18 新特性测试

- [Go1.18新特性](https://jishuin.proginn.com/p/763bfbd73a5e)

## 新特性
- 支持泛型 Generics
- 支持Workspaces
- 新增strings API(clone、cut)


## 升级go1.18
- 查看当前开发机伤的golang版本
```shell
> go version
go version go1.16.3 linux/amd64
```
- 删除旧版本
```shell
> which go
/usr/local/go/bin/go
> sudo rm -rf /usr/local/go
```
- 到 go 下载页，下载最新版： https://go.dev/dl/
```shell
# 安装后，确认一下版本即可：
> go version
go version go1.18 linux/amd64
```