# microservice-GRPC-Gin-Golang
The repository implement Gin gonic combined with GRPC to build architect microservice. Each service is a sidecar (private database, instance)

## Features
 Per one service in backend communation via grpc.
Gateway service provide RESTful api to client get data .
![](https://outcrawl.com/static/7b089f424d0abd2c29eb2d51ed362550/3ea43/architecture.jpg)

## The element in one service: 
```
├── handler
│   └── hander.go
├── main.go
├── pb
│   ├── pb.go
│   └── pb.proto
└── repo
    ├── repo.go
    └── entity.go
```    
    
- Run project with dev environment by `docker-compose up -d`. 
- Gateway service run at `http://localhost:8080`

## Requiment setup protobuf
- `go get -u github.com/amsokol/protoc-gen-gotag`
- `protoc -I pb pb/user.proto --go_out=plugins=grpc:pb/. --proto_path=../../../`
