#!/usr/bin/env bash

set -ue
cd $(dirname "$0")/..

# sample
directory="sample"
protoc -I=. \
  -I ${GOPATH}/src \
  -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
  --go_out=./${directory} \
  --go_opt=module=github.com/mintak21/proto/${directory} \
  --validate_out="lang=go:./${directory}" \
  --validate_opt=module=github.com/mintak21/proto/${directory} \
  --doc_out=./${directory}/doc --doc_opt=markdown,doc.md \
  ./${directory}/proto/*.proto # -I ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
