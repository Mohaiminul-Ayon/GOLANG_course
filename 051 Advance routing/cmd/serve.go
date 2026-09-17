package cmd

import (
	"ecommerce/manager"
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	managerInstance:= manager.NewManager()

	// manager.Use(middleware.Logger,middleware.Hudai,middleware.CorseWithPreflight)

	mux := http.NewServeMux()

	// globalRouter := middleware.CorseWithPreflight(mux)
	
	// wrppedMux := manager.WrappedRouter(
	// 	mux,
	// 	middleware.Logger,
	// 	middleware.Hudai,
	// 	middleware.CorseWithPreflight,
	// )

	// globalMiddlewares := []manager.Middleware{
	// 	middleware.CorseWithPreflight,
	// 	middleware.Hudai,
	// 	middleware.Logger,
	// }

	managerInstance.Use(
		middleware.Preflight,
		middleware.Cors,
		middleware.Logger,
	)

	wrappedMux := managerInstance.WrappedRouter(mux)
	
	initRouts(mux,managerInstance)
	
	fmt.Println("Server running on : 8080")

	err := http.ListenAndServe(":8080", wrappedMux) //"Faild"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}