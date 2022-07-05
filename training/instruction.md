brew install protobuf

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

export PATH=$PATH:~/go/bin

protobuf:
    protoc -I=$PWD --go_out=$PWD $PWD/chat.proto
grpc:
    protoc --go_out=./generated --go_opt=paths=source_relative --go-grpc_out=./generated --go-grpc_opt=paths=source_relative chat.proto
