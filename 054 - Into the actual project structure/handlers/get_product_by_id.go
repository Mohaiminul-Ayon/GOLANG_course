package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

//GET /products/{productID}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("productID")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valied product id", 400)
		return
	}
	ProductList := database.GetProductList()
	for _, product := range ProductList {
		if product.Id == pId {
			util.SendData(w, product, 200)
			return
		}
	}
	util.SendData(w, "Data pai nai", 404)
}
