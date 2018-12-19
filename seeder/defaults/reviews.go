package defaults

import (
	"github.com/chri5bot/reviews-pya/models"
)

// Reviews default reviews
var Reviews = []models.Review{
	models.Review{
		ID:         Review1,
		UserID:     User1,
		StoreID:    Store1,
		PurchaseID: Purchase1,
		Score:      5,
		Opinion:    "Esta bien",
		CreatedAt:  Date,
		UpdatedAt:  Date,
	},
	models.Review{
		ID:         Review2,
		UserID:     User2,
		StoreID:    Store2,
		PurchaseID: Purchase2,
		Score:      1,
		Opinion:    "Esta mal",
		CreatedAt:  Date,
		UpdatedAt:  Date,
	},
}
