setup:
	mkdir -p "./delivery/proto"


protoc: setup
	protoc -I=./ --go_out=./delivery/proto  --go-grpc_out=require_unimplemented_servers=false:./delivery/proto  ./shared/proto/*.proto
