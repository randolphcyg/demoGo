## demoGo

go-zero http服务 内部启动一个grpc客户端，连接java dubbo3服务

## 1. 初始化

编写修改 demoGo.api

后面若更新api文件，直接在项目根目录执行此命令刷新项目

```shell
goctl api go --api demoGo.api --dir . --style goZero
```

## 2. 用proto文件生成go-zero代码桩

为什么说是go-zero代码桩呢？

因为go-zero基于grpc封装，除了生成原生的pb stub，还会生成自己封装的pb。

我这边选择新开一个文件夹，用goctl rpc命令生成一个rpcX项目，
然后将client和pb文件夹复制到`/javastub`文件夹内即可。

写了个脚本 `./genStub.sh`执行一下 然后javastub文件夹下的client将pb的依赖修改下(删除了再加)

```shell
goctl rpc protoc demo.proto --go_out=./tmp --go-grpc_out=./tmp --zrpc_out=./tmp  --style goZero
```
生成的代码桩会在/tmp目录下，将 /tmp/democlient、/tmp/demopb 复制到/javastub文件夹下，client记得修改依赖引用位置。

demo.proto
```shell
syntax = "proto3";

package demopb;

option go_package="./demopb";

// 请求体
message HelloReq {
  string name = 1;
}

// 响应体
message HelloResp {
  string message = 1;
}

// demo 客户端
service demoClient {
  rpc SayHello (HelloReq) returns (HelloResp) {}
}
```