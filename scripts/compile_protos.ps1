$Env:OUTPUT_PATH = 'pkg/services/'
$Env:IMPORT_PATH = 'pkg/proto/'

protoc --proto_path=$Env:IMPORT_PATH --go-grpc_out=$Env:OUTPUT_PATH --go_out=$Env:OUTPUT_PATH pkg/proto/*.proto