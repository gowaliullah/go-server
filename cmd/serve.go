package cmd

import (
	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/repo"
	"github.com/gowaliullah/ecommerce/rest"
	"github.com/gowaliullah/ecommerce/rest/handlers/product"
	"github.com/gowaliullah/ecommerce/rest/handlers/user"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	productRepo := repo.NewProductRepo()
	userRepo := repo.NewUser()

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
