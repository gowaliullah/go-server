package product

import middleware "github.com/gowaliullah/ecommerce/rest/middlewares"

type Handler struct {
	middlewares *middleware.Middlewares
}

func NewHandler(middlewares *middleware.Middleware) *Handler {
	return &Handler{
		middlewares: middlewares,
	}
}
