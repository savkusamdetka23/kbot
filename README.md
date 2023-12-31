# kbot - devops application from scratch

Hi! This repo is for telegram bot which can translate your messages or forwarded posts to Ukrainian and provide transliteration for it or translate it from Ukrainian to English.


# Tech stack

 - Code: Goland
 - Golang framework for telegram bots-  [**telebot.v3**](https://gopkg.in/telebot.v3)
 - Framework for CLI interfaces: [**cobra**](https://github.com/spf13/cobra)
 - Other packages:
[**gtranslate**](https://github.com/bregydoc/gtranslate),
[**transliteration-go**](https://github.com/fre5h/transliteration-go),
[**language**](https://golang.org/x/text/language).

## Start using

To start using this bot you need to send messages to [**savkusamdetka23_bot**](https://t.me/savkusamdetka23_bot).

### Initialization
Just press start or type `/start`.

![Alt text](img/image.png)


### Translate to Ukrainian
After initialization you can start typing or forwarding posts and messages in English to receive a translation in Ukrainian with the transliteration.

![Alt text](img/image-1.png)


### Translate to English
Also you can post or forward messages in Ukrainian to receive tranlation in English.

![Alt text](img/image-23.png)


## Workflow diagram

![image](https://github.com/savkusamdetka23/kbot/assets/10897695/9bb66be6-728f-46d3-a5e8-8b1d8fdc04b3)

![image](https://github.com/savkusamdetka23/kbot/assets/10897695/8183dbce-76a2-4531-a86a-d3860fb193a5)

![image](https://github.com/savkusamdetka23/kbot/assets/10897695/b889ca7a-5467-46d0-8a59-aa934911ac89)


# Using gitleaks pre-commit hook
To install pre-commit git-hook, you can just execute next commands
```
 cp  git-hooks/pre-commit .git/hooks/ 
 git config core.hooksPath .git/hooks
 #Linux
 #chmod +x .git/hooks/pre-commit
 #Windows
 #icacls .git\hooks\pre-commit /grant Everyone:RX
```

To enable gitleaks check
```
 git config hooks.gitleaks true
```

To disable gitleaks check
```
 git config hooks.gitleaks false
```