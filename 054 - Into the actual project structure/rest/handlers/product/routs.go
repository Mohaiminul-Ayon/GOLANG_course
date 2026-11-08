package product

import (
	middleware "ecommerce/rest/middlewares"
	"net/http"
)

func (h *Handler) RegisterRouts(mux *http.ServeMux, manager *middleware.Manager) {

	mux.Handle(
		"GET /products",
		manager.With(
			http.HandlerFunc(h.GetProducts),
		),
	)
	mux.Handle(
		"POST /products",
		manager.With(
			http.HandlerFunc(h.CreateProduct),
			h.middlewares.AuthenticateJwt,
		),
	)
	mux.Handle(
		"GET /products/{id}",
		manager.With(
			http.HandlerFunc(h.GetProductByID),
		),
	)
	mux.Handle(
		"PUT /products/{id}",
		manager.With(
			http.HandlerFunc(h.UpdateProduct),
		),
	)
	mux.Handle(
		"DELETE /products/{id}",
		manager.With(
			http.HandlerFunc(h.DeleteProduct),
			h.middlewares.AuthenticateJwt,

		),
	)

}
