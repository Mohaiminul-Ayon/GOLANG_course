package rest

import (
	"ecommerce/rest/middlewares"
	"ecommerce/config"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Start(cnf config.Config) {
	managerInstance := middleware.NewManager()
	mux := http.NewServeMux()

	managerInstance.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := managerInstance.WrappedRouter(mux)

	InitRouts(mux, managerInstance)
	address := ":" + strconv.Itoa(cnf.HttpPort) //typecasting
	fmt.Println("Server running on port", address)

	err := http.ListenAndServe(address, wrappedMux)
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}
}
