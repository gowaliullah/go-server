package cmd

import (
	"fmt"
	"os"

	"github.com/gowaliullah/basic-ecommerce/config"
	"github.com/gowaliullah/basic-ecommerce/infra/db"
	"github.com/gowaliullah/basic-ecommerce/product"
	"github.com/gowaliullah/basic-ecommerce/repo"
	"github.com/gowaliullah/basic-ecommerce/rest"
	prdHandler "github.com/gowaliullah/basic-ecommerce/rest/handlers/product"
	usrHandler "github.com/gowaliullah/basic-ecommerce/rest/handlers/user"
	middleware "github.com/gowaliullah/basic-ecommerce/rest/middlewares"
	"github.com/gowaliullah/basic-ecommerce/user"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// repos
	productRepo := repo.NewProductRepo(*dbCon)
	userRepo := repo.NewUser(dbCon)

	// domains
	usrSrv := user.NewService(userRepo)
	prdSrc := product.NewService(productRepo)

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdHandler.NewHandler(middlewares, prdSrc)
	userHandler := usrHandler.NewHandler(cnf, usrSrv)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
