package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gowaliullah/ecommerce/config"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func Start(cnf config.Config) {
	mux := http.NewServeMux()
	manager := middleware.NewManager()

	manager.Use(
		// middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	wrappedMux := manager.WrapMux(mux)
	// initRoutes(mux, manager)

	addr := ":" + strconv.Itoa(cnf.HttpPort) // type casting (int64 to string)

	fmt.Println("server running on PORT:", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
