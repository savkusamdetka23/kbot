package cmd

import (
	"context"
	"fmt"
	"github.com/bregydoc/gtranslate"
	"github.com/fre5h/transliteration-go"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"os"
	"strings"
	"time"
	"unicode"

	"github.com/hirosassa/zerodriver"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
	telebot "gopkg.in/telebot.v3"
)

var (
	TELE_TOKEN   = os.Getenv("TELE_TOKEN")
	EnglishTag   = language.English
	UkrainianTag = language.Ukrainian
	MetricsHost  = os.Getenv("METRICS_HOST")
)

func initMetrics(ctx context.Context) {

	// Create a new OTLP Metric gRPC exporter with the specified endpoint and options
	exporter, _ := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(MetricsHost),
		otlpmetricgrpc.WithInsecure(),
	)

	// Define the resource with attributes that are common to all metrics.
	// labels/tags/resources that are common to all metrics.
	resource := resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(fmt.Sprintf("kbot_%s", appVersion)),
	)

	// Create a new MeterProvider with the specified resource and reader
	mp := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(resource),
		sdkmetric.WithReader(
			// collects and exports metric data every 10 seconds.
			sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(10*time.Second)),
		),
	)

	// Set the global MeterProvider to the newly created MeterProvider
	otel.SetMeterProvider(mp)

}

func pmetrics(ctx context.Context, payload string) {
	// Get the global MeterProvider and create a new Meter with the name "kbot_light_signal_counter"
	meter := otel.GetMeterProvider().Meter("kbot_light_signal_counter")

	// Get or create an Int64Counter instrument with the name "kbot_light_signal_<payload>"
	counter, _ := meter.Int64Counter(fmt.Sprintf("kbot_light_signal_%s", payload))

	// Add a value of 1 to the Int64Counter
	counter.Add(ctx, 1)
}

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
		logger := zerodriver.NewProductionLogger()

		kbot, err := telebot.NewBot(telebot.Settings{
			URL:    "",
			Token:  TELE_TOKEN,
			Poller: &telebot.LongPoller{Timeout: 10 * time.Second},
		})

		if err != nil {
			logger.Fatal().Str("Error", err.Error()).Msg("Please check TELE_TOKEN")
			return
		} else {
			logger.Info().Str("Version", appVersion).Msg("kbot started")
		}

		kbot.Handle(telebot.OnText, func(m telebot.Context) error {
			// Log the received text
			logger.Info().Str("Payload", m.Text()).Msg(m.Message().Payload)

			payload := m.Message().Payload
			pmetrics(context.Background(), payload)

			// Ensure we have non-empty text
			if m.Text() != "" {
				// Check for a specific command
				switch strings.ToLower(m.Text()) {
				case "/start":
					// Send a welcome message
					welcomeMessage := "СЛАВА УКРАЇНІ! I'm @savkusamdetka23_bot. Welcome! 😊 You can text me anything or forward other Telegram messages and posts, and I will translate your message to Ukrainian and provide transcription or translate it from Ukrainian to English."
					err := m.Send(welcomeMessage)
					if err != nil {
						logger.Printf("Error sending welcome message: %v", err)
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
						logger.Printf("Error translating text: %v", err)
						return err
					}

					err = m.Send(fmt.Sprintf("Translation: %s", translatedText))
					if err != nil {
						logger.Printf("Error sending translation: %v", err)
					}
				}
			}

			return nil
		})

		kbot.Start()
	},
}

func init() {
	ctx := context.Background()
	initMetrics(ctx)
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
