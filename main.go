package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

type Event struct {
	Event string `json:"event"`
}

func HandleRequest(ctx context.Context, event Event) (string, error) {
	sendPushNotification(event.Event)
	return fmt.Sprintf("Processed event %v", event.Event), nil
}

func main() {
	lambda.Start(HandleRequest)
}
