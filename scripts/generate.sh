#!/usr/bin/env bash

# reference. https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

set -ue
cd $(dirname "$0")/..

# sample
directory="sample"
# protoc -I. \
#   -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
#   -I ${GOPATH}/src/github.com/googleapis \
#   --go_out=./${directory}/go \
#   --go_opt=module=github.com/mintak21/proto/${directory} \
#   --go-grpc_out=./${directory}/go \
#   --go-grpc_opt=module=github.com/mintak21/proto/${directory} \
#   --validate_out="lang=go:./${directory}/go" \
#   --validate_opt=module=github.com/mintak21/proto/${directory} \
#   --doc_out=./${directory}/doc --doc_opt=markdown,doc.md \
#   ./${directory}/proto/*.proto # --go-grpc_opt=paths=github.com/mintak21/proto/${directory} \

protoc -I. \
  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
  -I ${GOPATH}/src/github.com/googleapis \
  --go_out=${GOPATH}/src \
  \
  --go-grpc_out=${GOPATH}/src \
  \
  --validate_out="lang=go:${GOPATH}/src" \
  \
  --doc_out=./${directory}/doc --doc_opt=markdown,doc.md \
  ./${directory}/proto/*.proto # --go_opt=paths=source_relative \
# --go-grpc_opt=paths=source_relative \
# --validate_opt=paths=source_relative \
