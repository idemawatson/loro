package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getListenPort() string {
	port := os.Getenv("PORT")
	if port != "" {
		return ":" + port
	}
	return ":3000"
}

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello world!"})
	})

	router.Run(getListenPort())
}
