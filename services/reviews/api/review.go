package api

import (
	"net/http"
	"strconv"

	"github.com/chri5bot/reviews-pya/services/reviews/models"
	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)

type reviewBody struct {
	UserID  uuid.UUID `json:"user_id"`
	StoreID uuid.UUID `json:"store_id"`
	Score   uint64    `json:"score"`
	Opinion string    `json:"opinion"`
}

// CreateReview creates a review
func (a *API) CreateReview(c *gin.Context) {
	body := &reviewBody{}
	if err := c.ShouldBindJSON(body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.UserID == (uuid.UUID{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The user is required"})
		return
	}

	if body.StoreID == (uuid.UUID{}) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The store is required"})
		return
	}

	if body.Score == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The score is required"})
		return
	}

	if body.Score > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The score must be between 1 and 5"})
		return
	}

	if body.Opinion == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "The opinion is required"})
		return
	}

	i, err := strconv.ParseUint(c.Param("purchaseID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result := a.db.Where("purchase_id = ? AND status = ?", i, models.ApprovedStatus).First(&models.Review{}); !result.RecordNotFound() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unique score per purchase"})
		return
	}

	review, err := models.NewReview(&models.NewReviewConfig{
		PurchaseID: i,
		UserID:     body.UserID,
		StoreID:    body.StoreID,
		Opinion:    body.Opinion,
		Score:      body.Score,
		Status:     models.ApprovedStatus,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if result := a.db.Create(&review); result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusCreated, review)
}
