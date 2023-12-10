#.DEFAULT_GOAL := linux

APP=$(shell basename -s .git ${shell git remote get-url origin})
REGISTRY_DOCKERHUB=maxeem23
REGISTRY_GCR=gcr.io
PROJECT_GCR=translatebot-405321
REGISTRY_GHCR=ghcr.io
PROJECT_GHCR=savkusamdetka23
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

# Define the supported platforms
PLATFORMS = linux windows arm macos
TARGETARCH ?= amd64
TARGETOS ?= linux
linux: TARGETOS=linux TARGETARCH=amd64
windows: TARGETOS=windows TARGETARCH=amd64
arm: TARGETOS=linux TARGETARCH=arm64 
macos: TARGETOS=darwin TARGETARCH=amd64

TAG=${APP}:${VERSION}-${TARGETOS}-${TARGETARCH}
TAG_GCR=${REGISTRY_GCR}/${PROJECT_GCR}/${TAG}
TAG_GHCR=${REGISTRY_GHCR}/${PROJECT_GHCR}/${TAG}
TAG_DOCKERHUB=${REGISTRY_DOCKERHUB}/${TAG}

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
#	docker build . -t ${TAG_GCR}
#	docker build . -t ${TAG_DOCKERHUB}
	docker build . -t ${TAG_GHCR}

push:
#	docker push ${TAG_GCR}
#	docker push ${TAG_DOCKERHUB}
	docker push ${TAG_GHCR}

clean:
	rm -rf kbot
#	docker rmi -f ${TAG_GCR}
#	docker rmi -f ${TAG_DOCKERHUB}
	docker rmi -f ${TAG_GHCR}


# Define OS and architecture params for tasks
$(PLATFORMS): %: build

# Make the tasks phony
.PHONY: $(PLATFORMS) build format lint test get image push clean
