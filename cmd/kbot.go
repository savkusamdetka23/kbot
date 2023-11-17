package cmd

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/bregydoc/gtranslate"
	"github.com/fre5h/transliteration-go"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken    = os.Getenv("TELE_TOKEN")
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
			Token:  TeleToken,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			log.Fatalf("Please check TELE_TOKEN env variable. %s", err)
			return
		}

		// Custom function to check if the text contains Latin characters
		hasLatin := func(text string) bool {
			match, _ := regexp.MatchString("[a-zA-Z]", text)
			return match
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			// Log the received text
			log.Printf("Received Text: %s", m.Text())

			// Ensure we have non-empty text
			if m.Text() != "" {
				// Check for a specific command
				switch strings.ToLower(m.Text()) {
				case "/start":
					// Send a welcome message and provide the current Kyiv time
					welcomeMessage := "СЛАВА УКРАЇНІ! I'm @savkusamdetka23_bot. Welcome! 😊 You can text me anything, and I will try to translate it to Ukrainian and provide transcription."
					currentTime := time.Now().In(time.FixedZone("Europe/Kiev", 2*60*60))
					timeMessage := fmt.Sprintf("The current time in Kyiv is: %s", currentTime.Format("15:04:05"))
					err := m.Send(welcomeMessage)
					if err != nil {
						log.Printf("Error sending welcome message: %v", err)
						return err
					}
					err = m.Send(timeMessage)
					if err != nil {
						log.Printf("Error sending time message: %v", err)
						return err
					}

				default:
					// Determine the script of the input text and translate accordingly
					var translatedText string
					if hasLatin(m.Text()) {
						// Input has Latin characters, translate to Ukrainian and provide transliteration
						translatedText, err = gtranslate.Translate(m.Text(), EnglishTag, UkrainianTag)
						if err == nil && strings.EqualFold(UkrainianTag.String(), "uk") {
							translatedText = fmt.Sprintf("%s\nTransliteration: %s", translatedText, transliteration.UkrToLat(translatedText))
						}
					} else {
						// Input has Cyrillic characters, translate to English
						translatedText, err = gtranslate.Translate(m.Text(), UkrainianTag, EnglishTag)
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
