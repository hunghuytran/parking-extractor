package main

import (
	"challenge/web_api/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.POST("/parking", handlers.GetParking)
	router.Run(":8080")
}
