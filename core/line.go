package core

import (
	"dictionary/models"
	"log"
	"os"

	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
)

func Send(t string, d models.Output) error {
	bot, err := messaging_api.NewMessagingApiAPI(
		os.Getenv("LINE_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	var messages []messaging_api.MessageInterface
	if d.Noun != nil {
		noun := messaging_api.TextMessage{
			Text: "Noun: " + *d.Noun,
		}
		messages = append(messages, noun)
	}

	if d.Verb != nil {
		verb := messaging_api.TextMessage{
			Text: "Verb: " + *d.Verb,
		}

		messages = append(messages, verb)
	}

	if d.Adverb != nil {
		adverb := messaging_api.TextMessage{
			Text: "Adverb: " + *d.Adverb,
		}

		messages = append(messages, adverb)
	}

	if d.Adjective != nil {
		adjective := messaging_api.TextMessage{
			Text: "Adjective: " + *d.Adjective,
		}

		messages = append(messages, adjective)
	}

	if d.Interjection != nil {
		interjection := messaging_api.TextMessage{
			Text: "Interjection: " + *d.Interjection,
		}

		messages = append(messages, interjection)
	}

	if d.Error != nil {
		e := messaging_api.TextMessage{
			Text: d.Error.Message,
		}

		messages = append(messages, e)
	}

	log.Printf("%v", messages)

	_, err = bot.ReplyMessage(
		&messaging_api.ReplyMessageRequest{
			ReplyToken: t,
			Messages:   messages,
		},
	)

	if err != nil {
		log.Print(err)
		return err
	}

	log.Println("Sent text reply.")

	return nil
}
