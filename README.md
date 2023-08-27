# toll-calc

# Installing protobuf compiler for mac

brew install protobuff

# Installing GRPC and Protobuffer plugins for golang

1. ProtoBuffers
   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

2. GRPC
   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

3. Set the path
   PATH="${PATH}:${HOME}/go/bin"

4. install the package dependencies
   go get google.golang.org/protobuf/reflect/protoreflect
