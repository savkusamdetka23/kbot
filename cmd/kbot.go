/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bregydoc/gtranslate"
	"github.com/fre5h/transliteration-go"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	telebot "gopkg.in/telebot.v3"
)

var (
	TeleToken = os.Getenv("TELE_TOKEN")
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

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			// Log the received text
			log.Printf("Received Text: %s", m.Text())

			// Ensure we have non-empty text
			if m.Text() != "" {
				// Check for a specific command
				switch strings.ToLower(m.Text()) {
				case "/start":
					// Send a welcome message and provide the current Kyiv time
					welcomeMessage := "СЛАВА УКРАЇНІ! I'm @savkusamdetka23_bot. Welcome! 😊 You can text me anything and I will try to translate it to Ukrainiand and provide transcription"
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
					// Translate any incoming text message from English to Ukrainian
					translatedText, err := gtranslate.Translate(m.Text(), language.English, language.Ukrainian)
					if err != nil {
						log.Printf("Error translating text: %v", err)
						return err
					}
					transliteratedText := transliteration.UkrToLat(translatedText)
					err = m.Send(fmt.Sprintf("Translation: %s\nTransliteration: %s", translatedText, transliteratedText))
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
