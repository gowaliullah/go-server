package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/middleware"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()

	manager.Use(
		// middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	fmt.Println("server running on: 8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
