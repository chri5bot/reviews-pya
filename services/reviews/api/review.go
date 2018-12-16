package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateReview creates a review
func (a *API) CreateReview(c *gin.Context) {
	c.JSON(http.StatusCreated, nil)
}
