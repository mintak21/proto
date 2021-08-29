FROM golang:1.17-buster

WORKDIR /app

RUN apt-get update -y \
  && apt-get install -y \
  # protobuf-compiler \
  curl \
  unzip \
  && apt-get clean \
  && rm -rf /var/lib/apt/lists/*

ENV PROTOC_VERSION=3.17.3

RUN tmpdir=$(mktemp -d) \
  && curl -Ls -o ${tmpdir}/protoc.zip https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip \
  && unzip ${tmpdir}/protoc.zip -d /usr/local/  \
  && rm -rf /usr/local/readme.txt \
  && rm -rf ${tmpdir}

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.27 \
  && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1 \
  && go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5
