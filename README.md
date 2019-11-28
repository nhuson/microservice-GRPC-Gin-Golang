# microservice-GRPC-Gin-Golang
The repository implement Gin gonic combined with GRPC to build architect microservice. each service is sidecar (private database, instance)

# Setup Protobuf
go get -u github.com/amsokol/protoc-gen-gotag
protoc -I pb pb/user.proto --go_out=plugins=grpc:pb/. --proto_path=../../../

# Docs
- Create .env for each service.
- Client service provider restful api at http://localhost:8080
