#!/usr/bin/env bash

set -ue
cd $(dirname "$0")/..

# sample
protoc -I=. \
  --go_out=./sample \
  --go_opt=module=github.com/mintak21/proto/sample \
  --doc_out=./sample/doc --doc_opt=markdown,doc.md \
  ./sample/proto/*.proto
