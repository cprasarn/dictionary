package main

import (
	"context"
	"dictionary/core"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type Source struct {
	Type    string `json:"type"`
	GroupId string `json:"groupId"`
	UserId  string `json:"userId"`
}

type Message struct {
	ID   *string `json:"id"`
	Type string  `json:"type"`
	Text string  `json:"text"`
}
type Event struct {
	Type            string           `json:"type"`
	Message         *Message         `json:"message"`
	ReplyToken      string           `json:"replyToken"`
	Source          *Source          `json:"source"`
	Mode            *string          `json:"mode"`
	DeliveryContext *DeliveryContext `json:"deliveryContext"`
}

type Input struct {
	Destination *string `json:"destination"`
	Events      []Event `json:"events"`
}

func handleRequest(ctx context.Context, data json.RawMessage) {
	var cb Input
	err := json.Unmarshal(data, &cb)
	if err != nil {
		log.Printf("Failed to unmarshal data: %v", err)
		return
	}

	for _, event := range cb.Events {
		log.Printf("/dictionary called: %+v...\n", event)

		switch event.Type {
		case "message":
			switch event.Message.Type {
			case "text":
				response, err := core.GetDictionary(event.Message.Text)
				if err != nil {
					log.Printf("Failed to get dictionary: %v", err)
				}

				err = core.Send(event.ReplyToken, *response)
				if err != nil {
					log.Printf("Failed to send to line: %v", err)
				}
			default:
				log.Printf("Unsupported message content: %v\n", event.Message)
			}
		default:
			log.Printf("Unsupported message: %v\n", event)
		}
	}
}

func main() {
	lambda.Start(handleRequest)
}
