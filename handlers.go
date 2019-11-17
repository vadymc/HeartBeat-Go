package main

import (
	"bytes"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	telegramEnabled bool
	firebaseEnabled bool
)

func init() {
	telegramEnabled = os.Getenv("TELEGRAM_ENABLED") == "true"
	firebaseEnabled = os.Getenv("FIREBASE_ENABLED") == "true"
}

func createEvent(c *gin.Context) {
	body := getBody(c)

	if telegramEnabled {
		go sendTelegram(body)
	}

	if firebaseEnabled {
		go sendPushNotification(body)
	}

	c.Status(http.StatusOK)
}

func getBody(c *gin.Context) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	return buf.Bytes()
}
