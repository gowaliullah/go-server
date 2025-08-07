package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/handlers"
	"github.com/gowaliullah/ecommerce/middleware"
)

func Serve() {

	mux := http.NewServeMux()

	muxRouter := middleware.GlobalRouter(mux)

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))

	fmt.Println("server running on: 8080")

	err := http.ListenAndServe(":8080", muxRouter)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
