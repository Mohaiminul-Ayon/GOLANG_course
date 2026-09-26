package rest

import (
	"ecommerce/rest/handler"
	"ecommerce/rest/middlewares"
	"net/http"
)

func InitRouts(mux *http.ServeMux, manager *middleware.Manager) {

	
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
		"GET /products/{id}",
		manager.With(
				http.HandlerFunc(handlers.GetProductByID),
			),
		)
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
				http.HandlerFunc(handlers.UpdateProduct),
			),
		)
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
				http.HandlerFunc(handlers.DeleteProduct),
			),
		)
	mux.Handle(
		"POST /user"
		manager.With(
				http.HandlerFunc(handlers.CreateUser),
			),
		)



}
