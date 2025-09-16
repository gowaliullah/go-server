package rest

import (
	"net/http"

	"github.com/gowaliullah/ecommerce/rest/handlers"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// mux.Handle("GET /habib", manager.With(http.HandlerFunc(handlers.Test)))

	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct), middleware.AuthenticateJwt))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreateProduct), middleware.AuthenticateJwt))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct), middleware.AuthenticateJwt))

	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.Login)))
}
