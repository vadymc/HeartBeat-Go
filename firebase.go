package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var client *messaging.Client

func init() {
	apiCredsConfig := os.Getenv("FIREBASE_CONFIG")
	url := os.Getenv("FIREBASE_URL")
	credentials := option.WithCredentialsFile(apiCredsConfig)
	endpoint := option.WithEndpoint(url)
	conf := &firebase.Config{
		ProjectID: "heartbeatpublisher",
	}
	app, err := firebase.NewApp(context.Background(), conf, credentials, endpoint)
	if err != nil {
		log.Fatalf("Failed to initialize firebase: %v\n", err)
		return
	}
	client, err = app.Messaging(context.Background())
	if err != nil {
		log.Fatalf("Failed to initialize client: %v\n", err)
		return
	}
}

func sendPushNotification(event string) {
	title, body := parseEvent(event)
	data := &messaging.Notification{
		Title: title,
		Body:  body,
	}
	msg := &messaging.Message{
		Notification: data,
		Topic:        "notification_events",
	}
	client.Send(context.Background(), msg)
	fmt.Printf("Processed %v", event)
}

func parseEvent(event string) (string, string) {
	i := strings.Index(event, "]")
	if i > -1 {
		return event[:i+1], event[i+1:]
	}
	return "", event
}
