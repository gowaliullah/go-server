package user

import (
	"net/http"

	"github.com/gowaliullah/ecommerce/rest/handlers"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

func func(h *Handler) RegisterRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("POST /users", manager.With(http.HandlerFunc(handlers.CreateUser)))
	mux.Handle("POST /users/login", manager.With(http.HandlerFunc(handlers.Login)))
}
