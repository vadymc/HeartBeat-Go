package main

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func createEvent(c *gin.Context) {
	body := getBody(c)
	go sendPushNotification(body)
	c.Status(http.StatusOK)
}

func getBody(c *gin.Context) string {
	buf := new(bytes.Buffer)
	buf.ReadFrom(c.Request.Body)
	body := buf.String()
	return body
}