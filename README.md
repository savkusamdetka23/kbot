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

 https://github.com/tucnak/telebot

#align tabulation
 gofmt -s -w ./
  go build -ldflags "-X="github.com/savkusamdetka23/kbot/cmd.appVersion=1.0.1
 ./kbot start

# @BotFather official bot manager
/newbot

read -s TELE_TOKEN
echo $TELE_TOKEN
export TELE_TOKEN
./kbot start