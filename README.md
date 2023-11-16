# kbot
devops application from scratch
 go mod init github.com/savkusamdetka23/kbot
 go install github.com/spf13/cobra-cli@latest
 cobra-cli init
 cobra-cli add version
 go run main.go help
 go run main.go version
 cobra-cli add kbot
 go build -ldflags "-X="github.com/savkusamdetka23/kbot/cmd.appVersion=1.0.0
 ./kbot version