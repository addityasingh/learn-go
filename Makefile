configure-protobuf:
	export GOPATH=$(pwd)
	export PATH=$PATH:$GOPATH/bin
	protoc --go_out=plugins=grpc:. src/github.com/addityasingh/protobuf/pb/server.proto