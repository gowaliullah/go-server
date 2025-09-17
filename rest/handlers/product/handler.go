package product

import (
	"github.com/gowaliullah/ecommerce/repo"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

type Handler struct {
	middlewares *middleware.Middlewares
	productRepo repo.ProductRepo
}

func NewHandler(middlewares *middleware.Middlewares, productRepo repo.ProductRepo) *Handler {
	return &Handler{
		middlewares: middlewares,
		productRepo: productRepo,
	}
}
