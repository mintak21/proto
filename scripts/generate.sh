#!/usr/bin/env bash

# reference. https://developers.google.com/protocol-buffers/docs/reference/go-generated#package

set -ue
cd $(dirname "$0")/..

# sample
directory="sample"
protoc -I. \
  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
  -I ${GOPATH}/src/github.com/googleapis \
  --go_out=${GOPATH}/src \
  --go-grpc_out=${GOPATH}/src \
  --validate_out="lang=go:${GOPATH}/src" \
  --doc_out=./${directory}/doc --doc_opt=markdown,doc.md \
  ./${directory}/proto/*.proto
