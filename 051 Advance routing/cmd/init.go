package cmd

import (
	handlers "ecommerce/handlers"
	"ecommerce/manager"
	"net/http"
)

func initRouts(mux *http.ServeMux, manager *manager.Manager) {

	
	mux.Handle(
		"GET /Ayon",
		manager.With(
				http.HandlerFunc(handlers.Test),
			),
		)
	mux.Handle(
		"GET /route",
		manager.With(
			http.HandlerFunc(handlers.Test),
			),
		)
	mux.Handle(
		"GET /products",
		manager.With(
				http.HandlerFunc(handlers.GetProducts),
			),
		)
	mux.Handle(
		"POST /products", 
		manager.With(
				http.HandlerFunc(handlers.CreateProduct),
			),
		)
	mux.Handle(
		"GET /products/{productID}",
		manager.With(
				http.HandlerFunc(handlers.GetProductByID),
			),
		)



}
