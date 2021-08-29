export DOCKER_BUILDKIT=1
IMAGE=ghcr.io/mintak21/proto-builder:latest
BUILD_PATH=/app

.PHONY: build prepare

build:
	docker build . -f ./Dockerfile -t ${IMAGE}

prepare:
	docker pull ${IMAGE} || make build

generate: prepare
	docker run --rm -it -v ${PWD}:${BUILD_PATH} ${IMAGE} bash -c ./scripts/generate.sh

push: build
	docker push ${IMAGE}
