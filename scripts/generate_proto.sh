#!/bin/bash

PROTO_DIR=./api/proto
OUT_DIR=./api/generated

# Создаем выходные папки, если их нет
mkdir -p $OUT_DIR/auth
mkdir -p $OUT_DIR/messenger

# Генерация Go-кода для auth.proto
protoc -I $PROTO_DIR \
  --go_out=$OUT_DIR/auth --go_opt=paths=source_relative \
  --go-grpc_out=$OUT_DIR/auth --go-grpc_opt=paths=source_relative \
  $PROTO_DIR/auth.proto

# Генерация Go-кода для messenger.proto
protoc -I $PROTO_DIR \
  --go_out=$OUT_DIR/messenger --go_opt=paths=source_relative \
  --go-grpc_out=$OUT_DIR/messenger --go-grpc_opt=paths=source_relative \
  $PROTO_DIR/messenger.proto

echo "gRPC-код успешно сгенерирован!"
