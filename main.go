package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type JsonData struct {
	Name string `json:Name,omitempty`
	Age  string `json:Age,omitempty`
}

func handler(ctx context.Context, sqsEvent events.SQSEvent) error {
	for _, message := range sqsEvent.Records {

		var body JsonData

		byte_body := []byte(message.Body)
		json.Unmarshal(byte_body, &body)

		log.Println(message.MessageId)
		log.Println(body)
		log.Printf("The message %s for event source %s = %s \n", message.MessageId, message.EventSource, message.Body)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
