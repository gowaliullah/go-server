package cmd

import (
	"fmt"
	"os"

	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/infra/db"
	"github.com/gowaliullah/ecommerce/repo"
	"github.com/gowaliullah/ecommerce/rest"
	"github.com/gowaliullah/ecommerce/rest/handlers/product"
	"github.com/gowaliullah/ecommerce/rest/handlers/user"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(dbCon)

	productRepo := repo.NewProductRepo(*dbCon)
	userRepo := repo.NewUser(dbCon)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := product.NewHandler(middlewares, productRepo)
	userHandler := user.NewHandler(cnf, userRepo)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
