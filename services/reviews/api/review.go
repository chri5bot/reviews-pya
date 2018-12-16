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

	purchaseID, err := strconv.ParseUint(c.Param("purchaseID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parsing purchase id"})
		return
	}

	if result := a.db.Where("purchase_id = ?", purchaseID).First(&models.Review{}); !result.RecordNotFound() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Unique score per purchase"})
		return
	}

	review, err := models.NewReview(&models.NewReviewConfig{
		PurchaseID: purchaseID,
		UserID:     body.UserID,
		StoreID:    body.StoreID,
		Opinion:    body.Opinion,
		Score:      body.Score,
	})
	if err != nil {
		c.Error(err)
		return
	}

	if result := a.db.Create(&review); result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusCreated, review)
}

// DeleteReview deletes a review
func (a *API) DeleteReview(c *gin.Context) {
	ID, err := uuid.FromString(c.Param("reviewID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := a.db.Where("id = ?", ID).Delete(&models.Review{}); result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, "Review Deleted")
}

// GetReview retrieve a review
func (a *API) GetReview(c *gin.Context) {
	pur, err := strconv.ParseUint(c.Param("purchaseID"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Parsing purchase id"})
		return
	}

	var review models.Review
	if result := a.db.Where("purchase_id = ?", pur).First(&review); result.Error != nil {
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, review)
}

// ListReviews retrieves all reviews
func (a *API) ListReviews(c *gin.Context) {
	storeID, err := uuid.FromString(c.Param("storeID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	query := a.db

	start := c.Query("start")
	if start != "" {
		query = query.Where("updated_at >= ?", start)
	}

	end := c.Query("end")
	if end != "" {
		query = query.Where("updated_at <= ?", end)
	}

	reviews := make([]*models.Review, 0)
	if result := query.Where("store_id = ?", storeID).Find(&reviews); result.Error != nil {
		if result.RecordNotFound() {
			c.JSON(http.StatusNotFound, gin.H{"error": "Records not found"})
			return
		}
		c.Error(result.Error)
		return
	}

	c.JSON(http.StatusOK, reviews)
}
