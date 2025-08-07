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
	mux.Handle("GET /hello", http.HandlerFunc(handlers.HelloHandlar))
	mux.Handle("GET /about", http.HandlerFunc(handlers.AboutHandlar))
	mux.Handle("POST /add-product", http.HandlerFunc(handlers.CreateProduct))

	fmt.Println("server running on: 8080")

	err := http.ListenAndServe(":8080", muxRouter)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
