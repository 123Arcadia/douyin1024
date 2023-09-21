# douyin1024

=======

## 抖音项目


工程无其他依赖，直接编译运行即可
```shell
go mod tidy
```

```shell
go build && ./simple-demo
```

## 配置地址

在`config.go`中配置 `Ipv4Address`

可以在cmd查看中输入查看 IPv4 地址

```shell
ipconfig
```

注意: 

在APP中可以使用模拟器打开，在首次打开(未登录)"我的"，双击2次，可以打开APP的服务端地址，填写自己的
 ipv4Addr 和 port 

如果是登录状态，打开"我的"可以在"高级设置"中填写
