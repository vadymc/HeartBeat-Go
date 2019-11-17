package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	if os.Getenv("FIREBASE_ENABLED") == "true" {
		initFirebase()
	}
	router := gin.Default()
	router.POST("/v1/events/", createEvent)
	router.Run(":11001")
}
