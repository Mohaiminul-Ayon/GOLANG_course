package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
)

//POST
func CreateProduct(w http.ResponseWriter,r *http.Request){
var newProduct database.Product
//decoder
decoder := json.NewDecoder(r.Body)
err := decoder.Decode(&newProduct)
if err!=nil {
	fmt.Println(err)
	http.Error(w,"Plz give me valid Json",400)
	return 
}

newProduct.Id = len(database.ProductList)+1
database.ProductList = append(database.ProductList, newProduct)

//encoder
util.SendData(w,newProduct,201)


}
