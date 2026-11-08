package rest

import (
	"ecommerce/config"
	"ecommerce/rest/handlers/product"
	"ecommerce/rest/handlers/review"
	"ecommerce/rest/handlers/user"
	middleware "ecommerce/rest/middlewares"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

type Server struct {
	cnf            *config.Config
	productHandler *product.Handler
	userHandler    *user.Handler
	reviewHandler  *review.Handler
}

func NewSrver(
	cnf *config.Config,
	productHandler *product.Handler,
	userHandler *user.Handler,
	reviewHandler *review.Handler,
) *Server {
	return &Server{
		cnf:            cnf,
		productHandler: productHandler,
		userHandler:    userHandler,
		reviewHandler:  reviewHandler,
	}
}
func (server *Server) Start() {
	managerInstance := middleware.NewManager()

	managerInstance.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	mux := http.NewServeMux()
	wrappedMux := managerInstance.WrappedRouter(mux)

	server.productHandler.RegisterRouts(mux, managerInstance)
	server.userHandler.RegisterRouts(mux, managerInstance)
	server.reviewHandler.RegisterRouts(mux, managerInstance)

	// InitRouts(mux, managerInstance)
	address := ":" + strconv.Itoa(server.cnf.HttpPort) //typecasting
	fmt.Println("Server running on port", address)

	err := http.ListenAndServe(address, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
