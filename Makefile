VERSION=$(shell git describe --tags --abbrev=0)-$(shell git rev-parse --short HEAD)
TARGETOS=linux
ARCHITECTURE=$(shell dpkg --print-architecture)

install: 
	go get

format:
	gofmt -s -w ./

lint:
	golint

test:
	go test -v

build:
	CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${ARCHITECTURE} go build -v -o kbot -ldflags "-X="github.com/savkusamdetka23/kbot/cmd.appVersion=${VERSION}

clean:
	rm -rf kbot