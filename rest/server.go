package rest

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gowaliullah/basic-ecommerce/config"
	"github.com/gowaliullah/basic-ecommerce/rest/handlers/product"
	"github.com/gowaliullah/basic-ecommerce/rest/handlers/user"
	middleware "github.com/gowaliullah/basic-ecommerce/rest/middlewares"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
}

func NewServer(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
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
