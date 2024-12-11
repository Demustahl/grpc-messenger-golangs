#!/bin/bash

PROTO_DIR=./api/proto
OUT_DIR=./api/generated
WEB_OUT_DIR=./web/src

# Создаем выходные папки, если их нет
mkdir -p $OUT_DIR
mkdir -p $WEB_OUT_DIR

# Генерация Go-кода
protoc -I $PROTO_DIR \
  --go_out=$OUT_DIR --go_opt=paths=source_relative \
  --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
  $PROTO_DIR/*.proto

# Генерация кода для gRPC-Web (JS/TS)
protoc -I $PROTO_DIR \
  --grpc-web_out=import_style=commonjs,mode=grpcwebtext:$WEB_OUT_DIR \
  $PROTO_DIR/*.proto

echo "gRPC-код успешно сгенерирован!"
