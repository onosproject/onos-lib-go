export CGO_ENABLED=1
export GO111MODULE=on

.PHONY: build

ONOS_PROTOC_VERSION := v0.6.9

build: # @HELP build the Go binaries (default)
build:
	go build github.com/onosproject/onos-lib-go/pkg/...

build-tools:=$(shell if [ ! -d "./build/build-tools" ]; then cd build && git clone https://github.com/onosproject/build-tools.git; fi)
include ./build/build-tools/make/onf-common.mk

test: # @HELP run the unit tests and source code validation  producing a golang style report
test: build deps license_check linters
	go test -race github.com/onosproject/onos-lib-go/pkg/...

jenkins-test:  # @HELP run the unit tests and source code validation producing a junit style report for Jenkins
jenkins-test: build deps license_check linters
	TEST_PACKAGES=github.com/onosproject/onos-lib-go/pkg/... ./build/build-tools/build/jenkins/make-unit

protos: # @HELP compile the protobuf files (using protoc-go Docker)
	docker run -it -v `pwd`:/go/src/github.com/onosproject/onos-lib-go \
		-w /go/src/github.com/onosproject/onos-lib-go \
		--entrypoint build/bin/compile-protos.sh \
		onosproject/protoc-go:${ONOS_PROTOC_VERSION}

publish: # @HELP publish version on github and dockerhub
	./build/build-tools/publish-version ${VERSION}

jenkins-publish: jenkins-tools # @HELP Jenkins calls this to publish artifacts
	./build/build-tools/release-merge-commit
	./build/build-tools/build/docs/push-docs

all: test

clean:: # @HELP remove all the build artifacts
	go clean -testcache github.com/onosproject/onos-lib-go/...
