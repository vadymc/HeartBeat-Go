package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/v1/events/", createEvent)
	router.Run(":8080")
}
