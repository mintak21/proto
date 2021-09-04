FROM golang:1.17-buster

WORKDIR /app

RUN apt-get update -y \
  && apt-get install --no-install-recommends -y \
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
  && go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5 \
  && go install github.com/envoyproxy/protoc-gen-validate@v0.6 \
  && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.5.0 \
  && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@v2.5.0

# ソースパスに"validate/validate.proto"が必要なためgo installではなくgit cloneしている
RUN mkdir -p ${GOPATH}/src/github.com/envoyproxy \
  && git clone https://github.com/envoyproxy/protoc-gen-validate.git -b v0.6.1 ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
WORKDIR ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate
RUN make build

# ソースパスに"google/api/annotations.proto"が必要なためgo installではなくgit cloneしている
RUN mkdir -p ${GOPATH}/src/github.com/googleapis \
  && git clone https://github.com/googleapis/googleapis.git ${GOPATH}/src/github.com/googleapis
# WORKDIR ${GOPATH}/src/github.com/googleapis/googleapis
# RUN make install


WORKDIR ${GOPATH}/src/github.com/mintak21/proto
