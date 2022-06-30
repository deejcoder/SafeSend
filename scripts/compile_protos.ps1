protoc --proto_path=./proto --go-grpc_out=services/auth --go_out=services/auth proto/auth.proto
protoc --proto_path=./proto --go-grpc_out=services/models --go_out=services/models proto/request.proto
protoc --proto_path=./proto --go-grpc_out=services/models --go_out=services/models proto/response.proto