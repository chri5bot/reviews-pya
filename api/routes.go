package api

func (a *API) initRoutes() {
	a.handler.GET("stores/:storeID/reviews", a.ListReviews)
	a.handler.GET("purchases/:purchaseID/reviews", a.GetReview)
	a.handler.POST("purchases/:purchaseID/reviews", a.CreateReview)
	a.handler.DELETE("reviews/:reviewID", a.DeleteReview)
}
