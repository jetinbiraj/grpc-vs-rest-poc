#!/usr/bin/env bash

set -e

PROTO_DIR=proto
OUT_DIR=proto

protoc \
  --proto_path="$PROTO_DIR" \
  --go_out=paths=source_relative:"$OUT_DIR" \
  --go-grpc_out=paths=source_relative:"$OUT_DIR" \
  "$PROTO_DIR"/common.proto