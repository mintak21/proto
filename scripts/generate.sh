#!/usr/bin/env bash

set -ue
cd $(dirname "$0")/..

# sample
directory="sample"
protoc -I. \
  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
  -I ${GOPATH}/src/github.com/googleapis \
  --go_out=./${directory}/go \
  --go_opt=module=github.com/mintak21/proto/${directory} \
  --go-grpc_out=./${directory}/go \
  --go-grpc_opt=module=github.com/mintak21/proto/${directory} \
  --validate_out="lang=go:./${directory}/go" \
  --validate_opt=module=github.com/mintak21/proto/${directory} \
  --doc_out=./${directory}/doc --doc_opt=markdown,doc.md \
  ./${directory}/proto/*.proto # --go-grpc_opt=paths=github.com/mintak21/proto/${directory} \
