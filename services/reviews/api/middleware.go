package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func errorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastError := c.Errors.ByType(gin.ErrorTypePrivate).Last()

		if lastError != nil {
			log.Println(lastError)
			c.JSON(http.StatusInternalServerError, gin.H{"error": lastError.Error()})
		}
	}
}
