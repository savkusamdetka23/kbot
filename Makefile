#.DEFAULT_GOAL := linux

APP=$(shell basename ${shell git remote get-url origin})
REGISTRY=gcr.io
PROJECT=translatebot-405321
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)

# Define the supported platforms
PLATFORMS = linux windows arm macos
TAG=${REGISTRY}/${PROJECT}:${VERSION}-${TARGETARC}
# Define the task names
TARGETOS=linux 
TARGETARC=amd64 
linux: TARGETOS=linux TARGETARC=amd64 
windows: TARGETOS=windows TARGETARC=amd64
arm: TARGETOS=linux TARGETARC=arm64
macos: TARGETOS=darwin TARGETARC=amd64

# Export TARGETARC for subsequent targets
export TARGETARC

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get: 
	go get

build: format get
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARC} go build -v -o kbot -ldflags "-X="github.com/savkusamdetka23/kbot/cmd.appVersion=${VERSION}
export TARGETARC

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