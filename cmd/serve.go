package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/middleware"
)

func Serve() {

	manager := middleware.NewManager()
	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()

	initRoutes(mux, manager)

	muxRouter := middleware.GlobalRouter(mux)
	fmt.Println("server running on: 8080")
	err := http.ListenAndServe(":8080", muxRouter)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
