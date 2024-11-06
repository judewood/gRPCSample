#!/bin/bash

PROJECT_DIR="calculator"
PROTO_DIR="proto"
BIN_DIR="bin"
SERVER_DIR="server"
CLIENT_DIR="client"
MODULE="github.com/judewood/gRPCSample"
PROTO_FILENAME="calculator.proto"

echo "building ${PROJECT_DIR}"

protoc -I${PROJECT_DIR}/${PROTO_DIR} --go_out=. --go_opt=module=${MODULE} --go-grpc_out=. --go-grpc_opt=module=${MODULE} ${PROJECT_DIR}/${PROTO_DIR}/${PROTO_FILENAME}
 echo protoc -I${PROJECT_DIR}/${PROTO_DIR} --go_out=. --go_opt=module=${MODULE} --go-grpc_out=. --go-grpc_opt=module=${MODULE} ${PROJECT_DIR}/${PROTO_DIR}/${PROTO_FILENAME}