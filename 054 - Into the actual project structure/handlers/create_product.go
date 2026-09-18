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
productList:=database.GetProductList()
newProduct.Id = len(productList)+1
database.Store(newProduct)
// productList = append(productList, newProduct)

//encoder
util.SendData(w,newProduct,201)


}
