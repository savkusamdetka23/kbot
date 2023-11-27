#.DEFAULT_GOAL := linux

APP=$(shell basename ${shell git remote get-url origin})
REGISTRY=gcr.io
PROJECT=translatebot-405321
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

# Define the supported platforms
PLATFORMS = linux windows arm macos
TARGETARCH ?= amd64
linux: TARGETOS=linux TARGETARCH=amd64
windows: TARGETOS=windows TARGETARCH=amd64
arm: TARGETOS=linux TARGETARCH=arm64 
macos: TARGETOS=darwin TARGETARCH=amd64

TAG=${REGISTRY}/${PROJECT}/${APP}:${VERSION}-${TARGETARCH}

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get: 
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -v -o kbot -ldflags "-X="github.com/savkusamdetka23/kbot/cmd.appVersion=${VERSION}

image: 
	docker build . -t ${TAG}

push:
	docker push ${TAG}

clean:
	rm -rf kbot
	docker rmi -f ${TAG}


# Define OS and architecture params for tasks
$(PLATFORMS): %: build

# Make the tasks phony
.PHONY: $(PLATFORMS) build