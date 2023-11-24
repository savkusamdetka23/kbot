APP=$(shell basename ${shell git remote get-url origin})
REGISTRY=maxeem23
VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux #windows
ARCHITECTURE=$(shell dpkg --print-architecture)
TARGETARC=arm64 #arm64 amd64
format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

get: 
	go get

build:
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARC} go build -v -o kbot -ldflags "-X="github.com/savkusamdetka23/kbot/cmd.appVersion=${VERSION}

image:
	docker build . -t ${REGISTRY}/${APP}:${VERSION}-${TARGETARC}

push:
	docker push ${REGISTRY}/${APP}:${VERSION}-${TARGETARC}

clean:
	rm -rf kbot