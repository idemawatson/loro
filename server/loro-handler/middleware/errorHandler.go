package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func ErrorHandler(c *gin.Context) {
	c.Next()

	err := c.Errors.ByType(gin.ErrorTypePublic).Last()
	if err != nil {
		logger, zaperr := zap.NewProduction()
		if zaperr != nil {
			log.Fatal(err.Error())
		}
		logger.Error("Error",
			zap.Error(err.Err),
		)

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}
}
