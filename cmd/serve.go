package cmd

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/middleware"
)

func Serve() {

	cnf := config.GetCConfig()
	fmt.Printf("Service Name: %s, Version: %s, Http Port: %d\n", cnf.ServiceName, cnf.Version, cnf.HttpPort)

	mux := http.NewServeMux()
	manager := middleware.NewManager()

	manager.Use(
		// middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	wrappedMux := manager.WrapMux(mux)
	initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort) // type casting (int64 to string)

	fmt.Println("server running on PORT:", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
