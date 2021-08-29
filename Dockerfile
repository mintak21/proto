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
  && go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@v1.5 \
  && go install github.com/envoyproxy/protoc-gen-validate@v0.6

# ソースパスに"validate/validate.proto"が必要なためgo installではなくgit cloneしている
RUN mkdir -p ${GOPATH}/src/github.com/envoyproxy \
  && git clone https://github.com/envoyproxy/protoc-gen-validate.git -b v0.6.1 ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate \
  && (cd ${GOPATH}/src/github.com/envoyproxy/protoc-gen-validate && make build)
