package main

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var client *messaging.Client

const (
	configLocation = "FIREBASE_CONFIG"
	firebaseURL    = "FIREBASE_URL"
	pushTopic      = "notification_events"
	projectID = "heartbeatpublisher"
)

func initFirebase() {
	apiCredsConfig := os.Getenv(configLocation)
	url := os.Getenv(firebaseURL)
	credentials := option.WithCredentialsFile(apiCredsConfig)
	endpoint := option.WithEndpoint(url)
	conf := &firebase.Config{
		ProjectID: projectID,
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

func sendPushNotification(event []byte) {
	title, body := ParseEvent(event)
	data := &messaging.Notification{
		Title: title,
		Body:  body,
	}
	msg := &messaging.Message{
		Notification: data,
		Topic:        pushTopic,
	}
	client.Send(context.Background(), msg)
	fmt.Printf("Processed %v %v", title, body)
}
