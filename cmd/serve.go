package cmd

import (
	"fmt"
	"os"

	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/infra/db"
	"github.com/gowaliullah/ecommerce/repo"
	"github.com/gowaliullah/ecommerce/rest"
	prdHandler "github.com/gowaliullah/ecommerce/rest/handlers/product"
	usrHandler "github.com/gowaliullah/ecommerce/rest/handlers/user"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
	"github.com/gowaliullah/ecommerce/user"
)

func Serve() {
	cnf := config.GetConfig()

	dbCon, err := db.NewConnection(cnf.DB)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(dbCon)

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

	middlewares := middleware.NewMiddlewares(cnf)

	productHandler := prdHandler.NewHandler(middlewares, productRepo)
	userHandler := usrHandler.NewHandler(cnf, usrSrv)

	server := rest.NewServer(cnf, productHandler, userHandler)
	server.Start()
}
