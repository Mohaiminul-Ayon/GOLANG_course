package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func DeleteProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valied product id", 400)
		return
	}

	database.Delet(pId)

	util.SendData(w, "successfully deleted", 201)

}
