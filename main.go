package main

import (
	"context"
	"dictionary/core"
	"dictionary/models"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func handleRequest(ctx context.Context, data json.RawMessage) {
	var cb models.Input
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
