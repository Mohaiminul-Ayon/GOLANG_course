package cmd

import (
	"ecommerce/globalRouter"
	"ecommerce/manager"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager:= manager.NewManager()

	manager.Use(middleware.Logger,middleware.Hudai,middleware.Arekta)

	mux := http.NewServeMux()

	initRouts(mux,manager)

	globalRouter := globalRouter.GlobalRouter(mux)

	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080", globalRouter) //"Faild"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}