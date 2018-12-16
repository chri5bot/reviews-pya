package api

func (a *API) initRoutes() {
	a.handler.GET("reviews/", nil)
}
