package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
)

func GetProducts(w http.ResponseWriter, r *http.Request) { //response writer-w and r-  response
	util.SendData(w, database.ProductList, 200)
}