package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdateProduct(w http.ResponseWriter, r *http.Request) {

	productID := r.PathValue("id")
	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please give me a valied product id", 400)
		return
	}

	var newProduct database.Product
	//decoder
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid Json", 400)
		return
	}
	newProduct.Id = pId

	database.Update(newProduct)

	util.SendData(w, "successfully updated", 200)

}
