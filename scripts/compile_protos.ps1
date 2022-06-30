$Env:OUTPUT_PATH = 'services/'
$Env:IMPORT_PATH = 'proto/'

protoc --proto_path=$Env:IMPORT_PATH --go-grpc_out=$Env:OUTPUT_PATH --go_out=$Env:OUTPUT_PATH proto/auth.proto
protoc --proto_path=$Env:IMPORT_PATH --go-grpc_out=$Env:OUTPUT_PATH --go_out=$Env:OUTPUT_PATH proto/request.proto
protoc --proto_path=$Env:IMPORT_PATH --go-grpc_out=$Env:OUTPUT_PATH --go_out=$Env:OUTPUT_PATH proto/response.proto