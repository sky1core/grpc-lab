#!/bin/bash

# Define where to find the proto files and where to output the generated files
PROTO_DIR="./proto"
GEN_DIR="./proto/gen"

# Create the directories for the generated files
mkdir -p ${GEN_DIR}/go
mkdir -p ${GEN_DIR}/openapiv2
mkdir -p ${GEN_DIR}/protoset

# Run protoc
protoc -I proto \
  --go_out ${GEN_DIR}/go --go_opt paths=source_relative \
  --go-grpc_out ${GEN_DIR}/go --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ${GEN_DIR}/go --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ${GEN_DIR}/openapiv2 \
  ${PROTO_DIR}/memo.proto

# protoset 파일 생성
protoc -I ${PROTO_DIR} \
  --descriptor_set_out=${GEN_DIR}/protoset/memo.protoset \
  --include_imports \
  ${PROTO_DIR}/memo.proto

echo "Code generation complete."
