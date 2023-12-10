#.DEFAULT_GOAL := linux

APP=$(shell basename -s .git ${shell git remote get-url origin})
REGISTRY=gcr.io
REGISTRY_DOCKERHUB=maxeem23
PROJECT=translatebot-405321
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

# Define the supported platforms
PLATFORMS = linux windows arm macos
TARGETARCH ?= amd64
TARGETOS ?= linux
linux: TARGETOS=linux TARGETARCH=amd64
windows: TARGETOS=windows TARGETARCH=amd64
arm: TARGETOS=linux TARGETARCH=arm64 
macos: TARGETOS=darwin TARGETARCH=amd64

TAG=${REGISTRY}/${PROJECT}/${APP}:${VERSION}-${TARGETARCH}
TAG_DOCKERHUB=${REGISTRY_DOCKERHUB}/${APP}:${VERSION}-${TARGETARCH}

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
#	docker build . -t ${TAG}
	docker build . -t ${TAG_DOCKERHUB}

push:
#	docker push ${TAG}
	docker push ${TAG_DOCKERHUB}

clean:
	rm -rf kbot
	docker rmi -f ${TAG}


# Define OS and architecture params for tasks
$(PLATFORMS): %: build

# Make the tasks phony
.PHONY: $(PLATFORMS) build format lint test get image push clean
