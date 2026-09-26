package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

// POST
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var newProduct database.Product
	//decoder
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newProduct)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid Json", 400)
		return
	}

	createdProduct := database.Store(newProduct)

	util.SendData(w, createdProduct, 201)

}
