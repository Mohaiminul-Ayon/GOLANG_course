package cmd

import (
	"ecommerce/config"
	"ecommerce/manager"
	"ecommerce/middleware"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

func Serve() {
	 cnf:= config.Getconfig()
	managerInstance:= manager.NewManager()

	mux := http.NewServeMux()



	managerInstance.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := managerInstance.WrappedRouter(mux)
	
	initRouts(mux,managerInstance)
	address := ":" + strconv.Itoa(cnf.HttpPort)//typecasting
	fmt.Println("Server running on port",address)

	err := http.ListenAndServe(address, wrappedMux) 
	if err != nil {
		fmt.Println("Error starting the server", err)
		os.Exit(1)
	}

}