package product

import (
	"net/http"

	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

func (h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle("GET /products", manager.With(http.HandlerFunc(h.GetProducts)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(h.UpdateProduct), middleware.AuthenticateJwt))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(h.CreateProduct), middleware.AuthenticateJwt))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(h.GetProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(h.DeleteProduct), middleware.AuthenticateJwt))
}
