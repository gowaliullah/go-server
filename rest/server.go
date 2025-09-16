package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gowaliullah/ecommerce/config"
	"github.com/gowaliullah/ecommerce/rest/handlers/product"
	"github.com/gowaliullah/ecommerce/rest/handlers/user"
	middleware "github.com/gowaliullah/ecommerce/rest/middlewares"
)

type Server struct {
	productHandler *product.Handler
	userHandler    *user.Handler
	cnf            config.Config
}

func NewServer(
	productHandler *product.Handler,
	userHandler *user.Handler,
	cnf config.Config,
) *Server {
	return &Server{
		productHandler: productHandler,
		userHandler:    userHandler,
		cnf:            cnf,
	}
}

func (server *Server) Start() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()

	manager.Use(
		// middleware.Cors,
		middleware.Preflight,
		middleware.Logger,
	)
	wrappedMux := manager.WrapMux(mux)
	// initRoutes(mux, manager)

	server.productHandler.RegisterRoutes(mux, manager)
	server.userHandler.RegisterRoutes(mux, manager)

	addr := ":" + strconv.Itoa(server.cnf.HttpPort) // type casting (int64 to string)

	fmt.Println("server running on PORT:", addr)
	err := http.ListenAndServe(addr, wrappedMux)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
}
