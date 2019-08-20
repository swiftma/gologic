## go rpc example

参考：
https://ewanvalentine.io/microservices-in-golang-part-1/


安装 go rpc

https://grpc.io/docs/quickstart/go/


下面命令由于被墙会出错

go get -u google.golang.org/grpc

替换方式：

```
mkdir -p $GOPATH/src/google.golang.org/
cd $GOPATH/src/google.golang.org/
git clone --depth=1 https://github.com/grpc/grpc-go.git grpc
mkdir -p $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/google/go-genproto.git genproto

cd grpc
go install
```

```
mkdir -p $GOPATH/src/golang/
cd $GOPATH/src/golang/
git clone https://github.com/golang/protobuf
cd protobuf/protoc-gen-go
go install
```

build docker images requires golang.org/x/sys/unix


https://ewanvalentine.io/microservices-in-golang-part-2/

```
mkdir -p $GOPATH/src/golang.org/x/sys/
git clone https://github.com/golang/sys.git $GOPATH/src/golang.org/x/sys
```

Go-micro integrates as a protoc plugin, in this case replacing the standard gRPC plugin we're currently using.
```
go get -u github.com/micro/protobuf/{proto,protoc-gen-go}
go get github.com/micro/go-micro
```

删除指定image d72c2c74e975 不用的容器实例
docker ps -a|grep d72c2c74e975|awk '{print $1}'|xargs docker rm

docker stop $(docker ps -qa)
docker ps -a|awk '{print $1}'|xargs docker rm

docker-compose up -d
docker-compose run shippy-cli-consignment
