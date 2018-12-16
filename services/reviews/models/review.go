package models

import (
	"time"

	"github.com/gofrs/uuid"
)

// ApprovedStatus Order Status
const ApprovedStatus = "approved"

// RejectedStatus Order Status
const RejectedStatus = "rejected"

// Review model
type Review struct {
	ID         uuid.UUID  `json:"id"`
	UserID     uuid.UUID  `json:"user_id"`
	StoreID    uuid.UUID  `json:"store_id"`
	PurchaseID uint64     `json:"purchase_id"`
	Score      uint64     `json:"score"`
	Opinion    string     `json:"opinion"`
	Status     string     `json:"status"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"-"`
}

// NewReviewConfig are the params passed to the NewCustomer function
type NewReviewConfig struct {
	UserID     uuid.UUID
	StoreID    uuid.UUID
	PurchaseID uint64
	Score      uint64
	Opinion    string
	Status     string
}

// NewReview creates a new review
func NewReview(config *NewReviewConfig) (*Review, error) {
	ID, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &Review{
		ID:         ID,
		UserID:     config.UserID,
		StoreID:    config.StoreID,
		PurchaseID: config.PurchaseID,
		Score:      config.Score,
		Opinion:    config.Opinion,
		Status:     config.Status,
	}, nil
}
