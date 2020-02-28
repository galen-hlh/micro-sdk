# micro-sdk


### 安装依赖

- [安装protoc](https://github.com/golang/protobuf)
- [安装prototool](https://github.com/uber/prototool/blob/dev/docs/install.md)
- [安装go-micro](https://github.com/micro/protoc-gen-micro)
- [安装grpc_php_plugin](https://blog.csdn.net/myeye520/article/details/103923752)
```bash
git clone -b $(curl -L https://grpc.io/release) https://github.com/grpc/grpc   grpc
cd grpc
git submodule update --init
make grpc_php_plugin

# 把生成的文件加入环境变量
./grpc/bins/opt/grpc_php_plugin
```




### 生成sdk
```bash
prototool generate
```