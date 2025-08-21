package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/middleware"
)

func Serve() {

	manager := middleware.NewManager()
	mux := http.NewServeMux()

	// muxRouter := middleware.CorsWithPreflight(mux)
	wrappedMux := manager.WrapMux(mux, middleware.Logger, middleware.Hudai)

	initRoutes(mux, manager)

	fmt.Println("server running on: 8080")
	err := http.ListenAndServe(":8080", wrappedMux)
	if err != nil {
		fmt.Println("Error: ", err)
	}

}
