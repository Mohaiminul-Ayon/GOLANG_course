package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

//GET /products/{productID}
func GetProductByID(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valied product id", 400)
		return
	}


	ProductList := database.GetId(pId)
	if ProductList!=nil{
		util.SendData(w, ProductList, 200)
		return
	}
	util.Senderror(w,"Data pai nai",404)
}
