package product

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func (h *Handler)GetProducts(w http.ResponseWriter, r *http.Request) { //response writer-w and r-  response
	productList := database.List()
	util.SendData(w, productList, 200)
}
