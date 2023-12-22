FROM quay.io/projectquay/golang:1.20 AS builder

WORKDIR /go/src/app
ARG TARGETOS=linux 
ARG TARGETARCH=amd64
COPY . .
RUN make build GOOS=${TARGETOS} GOARCH=${TARGETARCH}

FROM scratch
WORKDIR /
COPY --from=builder /go/src/app/kbot .
COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT [ "./kbot", "start" ]