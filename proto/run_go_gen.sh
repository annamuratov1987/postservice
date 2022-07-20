export GOPROJECT=$GOPATH/src/postservice
protoc --proto_path=$GOPROJECT/proto/ --go_out=plugins=grpc:$GOPROJECT/delivery/grpc loader.proto grud.proto

