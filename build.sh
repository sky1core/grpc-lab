#!/bin/bash

# Define where to find the proto files and where to output the generated files
PROTO_DIR="./proto"
GEN_DIR="./proto/gen"

# Create the directories for the generated files
mkdir -p ${GEN_DIR}/go
mkdir -p ${GEN_DIR}/openapiv2

# Run protoc
protoc -I ${PROTO_DIR} \
  --go_out ${GEN_DIR}/go --go_opt paths=source_relative \
  --go-grpc_out ${GEN_DIR}/go --go-grpc_opt paths=source_relative \
  --grpc-gateway_out ${GEN_DIR}/go --grpc-gateway_opt paths=source_relative \
  --openapiv2_out ${GEN_DIR}/openapiv2 \
  ${PROTO_DIR}/*.proto

echo "Code generation complete."
