package cmd

import (
	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/rest"
	"github.com/gowaliullah/ecommerce/rest/handlers/product"
	"github.com/gowaliullah/ecommerce/rest/handlers/user"
)

func Serve() {
	cnf := config.GetConfig()

	productHandler := product.NewHandler()
	userHandler := user.NewHandler()

	server := rest.NewServer(cnf, userHandler, productHandler)
	server.Start()
}
