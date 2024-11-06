#!/bin/bash

PROJECT_DIR="greet"
PROTO_DIR="proto"
BIN_DIR="bin"
SERVER_DIR="server"
CLIENT_DIR="client"
MODULE="github.com/Clement-Jean/grpc-go-course"
PROTO_FILENAME="greet.proto"

echo "building ${PROJECT_DIR}"

protoc -I${PROJECT_DIR}/${PROTO_DIR} --go_out=. --go_opt=module=${MODULE} --go-grpc_out=. --go-grpc_opt=module=${MODULE} ${PROJECT_DIR}/${PROTO_DIR}/${PROTO_FILENAME}
