package cmd

import (
	"ecommerce/globalrouter"
	"ecommerce/handlers"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()



	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))
	mux.Handle("GET /products/{productID}", http.HandlerFunc(handlers.GetProductByID))



	globalRouter := globalrouter.GlobalRouter(mux)
	fmt.Println("Server running on : 8080")


	err := http.ListenAndServe(":8080", globalRouter) //"Faild"
	if err != nil {
		fmt.Println("Error starting the server", err)
	}

}