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
```
docker ps -a|grep d72c2c74e975|awk '{print $1}'|xargs docker rm
```
```
docker stop $(docker ps -qa)
docker ps -a|awk '{print $1}'|xargs docker rm

docker-compose up
docker-compose run shippy-cli-consignment

docker-compose run user-cli ./user-cli   --name="Ewan Valentine"  --email="ewan.valentine89@gmail.com"   --password="Testing123"   --company="BBC"

$ docker run -d -p 5432:5432 postgres
$ docker run -d -p 27017:27017 mongo

docker pull microhq/micro
```

```
curl -X POST -H 'Content-Type: application/json' \
    -d '{ "service": "shippy.auth", "method": "UserService.Create", "request": {  "email": "majc@gmail.com", "password": "testing123", "name": "Ewan Valentine", "company": "BBC"  } }' \
    http://localhost:8080/rpc

curl -X POST -H 'Content-Type: application/json' \
    -d '{ "service": "shippy.auth", "method": "UserService.GetAll", "request": {  } }' \
    http://localhost:8080/rpc


curl -XPOST -H 'Content-Type: application/json' \
    -d '{ "service": "shippy.auth", "method": "UserService.Auth", "request":  { "email": "majc@gmail.com", "password": "testing123" } }' \
    http://localhost:8080/rpc


curl -XPOST -H 'Content-Type: application/json' \
    -H 'Token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiOTE0ZTY0N2MtZjg4YS00OGEwLWFhNzUtMDg1ZGEyNzljZDBhIiwibmFtZSI6IkV3YW4gVmFsZW50aW5lIiwiY29tcGFueSI6IkJCQyIsImVtYWlsIjoibWFqY0BnbWFpbC5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCQ4dTl3VmV4d3huVEV6WEl5R3pKaGxPSzJ4WlhYY1VRekRjT203TUp2ZUNJdU0vQ1huY1VmaSJ9LCJleHAiOjE1NjY0ODA4NjksImlzcyI6ImdvLm1pY3JvLnNydi51c2VyIn0.O8FBeMimFN4B5Vt0fY7W-toOmAeOC2K9vLMpmHHENYE' \
    -d '{
      "service": "shippy.service.consignment",
      "method": "ShippingService.CreateConsignment",
      "request": {
        "description": "This is a test",
        "weight": "500",
        "containers": []
      }
    }' --url http://localhost:8080/rpc

```

micro api --address :8888 --namespace hello
