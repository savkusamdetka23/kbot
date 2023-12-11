package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/bregydoc/gtranslate"
	"github.com/fre5h/transliteration-go"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	telebot "gopkg.in/telebot.v3"
)

var (
	teletoken    = os.Getenv("teletoken")
	EnglishTag   = language.English
	UkrainianTag = language.Ukrainian
)

var kbotCmd = &cobra.Command{
	Use:     "kbot",
	Aliases: []string{"start"},
	Short:   "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("kbot %s started", appVersion)

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  teletoken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check teletoken env variable. %s", err)
			return
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			// Log the received text
			log.Printf("Received Text: %s", m.Text())

			// Ensure we have non-empty text
			if m.Text() != "" {
				// Check for a specific command
				switch strings.ToLower(m.Text()) {
				case "/start":
					// Send a welcome message
					welcomeMessage := "СЛАВА УКРАЇНІ! I'm @savkusamdetka23_bot. Welcome! 😊 You can text me anything or forward other Telegram messages and posts, and I will translate your message to Ukrainian and provide transcription or translate it from Ukrainian to English."
					err := m.Send(welcomeMessage)
					if err != nil {
						log.Printf("Error sending welcome message: %v", err)
						return err
					}

				default:
					// Determine the script of the input text and translate accordingly
					var translatedText string

					// Count the number of Cyrillic and Latin characters
					cyrillicCount, latinCount := countCharacters(m.Text())

					if cyrillicCount > latinCount {
						// Input has more Cyrillic characters, translate to English
						translatedText, err = gtranslate.Translate(m.Text(), UkrainianTag, EnglishTag)
					} else {
						// Input has more Latin characters, translate to Ukrainian and provide transliteration
						translatedText, err = gtranslate.Translate(m.Text(), EnglishTag, UkrainianTag)
						if err == nil && strings.EqualFold(UkrainianTag.String(), "uk") {
							translatedText = fmt.Sprintf("%s\nTransliteration: %s", translatedText, transliteration.UkrToLat(translatedText))
						}
					}

					if err != nil {
						log.Printf("Error translating text: %v", err)
						return err
					}

					err = m.Send(fmt.Sprintf("Translation: %s", translatedText))
					if err != nil {
						log.Printf("Error sending translation: %v", err)
					}
				}
			}

			return nil
		})

		kbot.Start()
	},
}

func init() {
	rootCmd.AddCommand(kbotCmd)
	// Add other initialization logic here...
}

// countCharacters counts the number of Cyrillic and Latin characters in a given text.
func countCharacters(text string) (cyrillicCount, latinCount int) {
	for _, runeValue := range text {
		switch {
		case unicode.Is(unicode.Cyrillic, runeValue):
			cyrillicCount++
		case unicode.Is(unicode.Latin, runeValue):
			latinCount++
		}
	}
	return cyrillicCount, latinCount
}
