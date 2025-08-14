package cmd

import (
	"fmt"
	"net/http"

	"github.com/gowaliullah/ecommerce/handlers"
	"github.com/gowaliullah/ecommerce/middleware"
)

func Serve() {

	mux := http.NewServeMux()

	mux.Handle("GET /habib", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.Test))))

	mux.Handle("GET /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProducts))))
	mux.Handle("POST /products", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.CreateProduct))))
	mux.Handle("GET /products/{id}", middleware.Hudai(middleware.Logger(http.HandlerFunc(handlers.GetProductById))))

	muxRouter := middleware.GlobalRouter(mux)
	fmt.Println("server running on: 8080")

	err := http.ListenAndServe(":8080", muxRouter)

	if err != nil {
		fmt.Println("Error: ", err)
	}

}
