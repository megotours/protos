#!/bin/bash

for dir in proto/*; do
  if [ -d "$dir" ]; then
    name=$(basename "$dir")
    protoc -I "$dir" "$dir"/*.proto --go_out="./gen/go/$name" --go_opt=paths=source_relative --go-grpc_out="./gen/go/$name" --go-grpc_opt=paths=source_relative
  fi
done
