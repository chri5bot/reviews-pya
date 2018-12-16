package api

func (a *API) initRoutes() {
	a.handler.POST("purchases/:purchaseID/reviews", a.CreateReview)
}
