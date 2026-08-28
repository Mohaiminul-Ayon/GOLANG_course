package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Product struct{
	Id int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Price float64 `json:"price"`
	ImgUrl string `json:"imageUrl"`

}

var productList[]Product 

// func HelloHandeler(w http.ResponseWriter,r *http.Request){//response writer-w and r-  response
// 	fmt.Fprintln(w, "Hello World")
// }
// func adminHandeler(w http.ResponseWriter,r *http.Request){//response writer-w and r-  response
// 	fmt.Fprintln(w, "I'm Ayon. I'm a great Engineer,I'm Software engineer")
// }

func HandleCorse(w http.ResponseWriter){
	w.Header().Set("Access-Control-Allow-Origin","*")
	w.Header().Set("Access-Control-Allow-Methods","GET,PUT,PATCH,DELETE,POST,OPTION")
	w.Header().Set("Access-Control-Allow-Headers","Content-Type, Ayon")
	w.Header().Set("Content-Type","application/json")
}

func HandleOption(w http.ResponseWriter,r *http.Request){
	if r.Method == "OPTIONS"{
		w.WriteHeader(200)
	}
}

func SendData(w http.ResponseWriter,data interface{},statusCode int){
	w.WriteHeader(statusCode)
	encoder := json.NewEncoder(w)
	encoder.Encode(data)
	
}



//GET
func getProducts(w http.ResponseWriter,r *http.Request){//response writer-w and r-  response
	HandleCorse(w)
	HandleOption(w,r)
	SendData(w,productList,200)
}

//POST
func createProduct(w http.ResponseWriter,r *http.Request){
HandleCorse(w)
HandleOption(w,r)


// description
// imageUrl
// price
// title

var newProduct Product
//decoder
decoder := json.NewDecoder(r.Body)
err := decoder.Decode(&newProduct)

if err!=nil {
	fmt.Println(err)
	http.Error(w,"Plz give me valid Json",400)
	return 
}



newProduct.Id = len(productList)+1
productList = append(productList, newProduct)

SendData(w,newProduct,201)
//encoder

}




func main(){
	mux := http.NewServeMux()

	// mux.HandleFunc("/hello",HelloHandeler)  
	// mux.HandleFunc("/admin",adminHandeler)
	mux.Handle("GET /products", http.HandlerFunc(getProducts))
    mux.Handle("OPTIONS /products", http.HandlerFunc(getProducts))
    mux.Handle("POST /create-products",http.HandlerFunc(createProduct) )
    mux.HandleFunc("OPTIONS /create-products", createProduct)

	fmt.Println("Server running on : 8080")

	err:= http.ListenAndServe(":8080",mux)//"Faild"
	if err!= nil{
		fmt.Println("Error starting the server",err)
	}
}



func init(){
	prd1:= Product{
		Id: 1,
		Title: "Orange",
		Description: "orange is grat ,I love orange",
		Price: 100,
		ImgUrl:"https://imgs.search.brave.com/XjEbXAKyW6pGkpxfKTvQT4-lxVvto-jaamT9fEF3sNk/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9pbWcu/bWFnbmlmaWMuY29t/L3ByZW1pdW0tcHNk/L2p1aWN5LW9yYW5n/ZS1zbGljZXMtZnJl/c2gtY2l0cnVzLWZy/dWl0LWdyZWVuLWxl/YXZlc184NDQ0My0z/NzE4NS5qcGc_c2Vt/dD1haXNfaHlicmlk/Jnc9NzQwJnE9ODA" ,
	}
	prd2:= Product{
		Id: 2,
		Title: "Apple",
		Description: "Apple is grat ,I love Apple",
		Price: 110,
		ImgUrl:"https://imgs.search.brave.com/WLzUepLPLVUHhWI8M6qCegqd6vzh2QLl_DiYDzXXtNc/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly93d3cu/cG5ncGxheS5jb20v/d3AtY29udGVudC91/cGxvYWRzLzYvU2xp/Y2VkLUFwcGxlLUZy/dWl0LVRyYW5zcGFy/ZW50LVBORy1wcmV2/aWV3LndlYnA" ,
	}
	prd3:= Product{
		Id: 3,
		Title: "mango",
		Description: "mango is grat ,I love mango",
		Price: 120,
		ImgUrl:"https://imgs.search.brave.com/b7swo_iOimkgNgtL-PckAjVMKEjDlw3b7SSecWz8qtQ/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9pbWcu/bWFnbmlmaWMuY29t/L3ByZW1pdW0tdmVj/dG9yL3JpcGUtZGVs/aWNpb3VzLWN1dC1m/cnVpdC1tYW5nby1s/ZW1vbi1zbGljZXMt/cG5nLXRyYW5zcGFy/ZW50LWJhY2tncm91/bmRfMTIzNjkyNy0x/MjkxOC5qcGc_Z2E9/R0ExLjEuMjA4NjA1/NTg5NC4xNzg1OTQw/NjgwJnNlbXQ9YWlz/X3Rlc3RfYiZ3PTc0/MCZxPTgw" ,
	}
	prd4:= Product{
		Id: 4,
		Title: "Kiwi",
		Description: "Kiwi is grat ,I love Kiwi",
		Price: 600,
		ImgUrl:"https://imgs.search.brave.com/IGcKTp7P2-Rhs6TeAjc3l0PDbTc60inyWWbVrUhsGdE/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9wbmcu/cG5ndHJlZS5jb20v/cG5nLXZlY3Rvci8y/MDI2MDUxNy9vdXJt/aWQvcG5ndHJlZS1h/LWhhbHZlZC1raXdp/LWZydWl0LXdpdGgt/Z3JlZW4tanVpY2Ut/c3BsYXNoaW5nLWFy/b3VuZC1pdC1vbi1w/bmctaW1hZ2VfMTkz/NDI3Mzkud2VicA" ,
	}
	prd5:= Product{
		Id: 5,
		Title: "Banana",
		Description: "Banana is grat ,I love Banana",
		Price: 80,
		ImgUrl:"https://imgs.search.brave.com/6ebB1a3qXuQle2AQekeSu18gie1JXNf92SU_DD-OuBo/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly90aHVt/Ym5haWwuaW1nYmlu/LmNvbS8yNC81LzIz/L2ltZ2Jpbi1iYW5h/bmEtZnJ1aXQtcmlw/ZW5pbmctZm9vZC1i/YW5hbmEtZnJ1aXQt/dVl5RHhnMURDN0Fq/d3JScWhGSEQ5d0xl/Nl90LmpwZw" ,
	}
	
	prd6 := Product{
		Id: 6,
		Title: "guava",
		Description: "guava is grat ,I love guava",
		Price: 80,
		ImgUrl:"https://imgs.search.brave.com/dtvi0KIWPtvn4Mt1lSdfzWcBr9K-dL7bA1p26PjM3K8/rs:fit:500:0:1:0/g:ce/aHR0cHM6Ly9zdGF0/aWMudmVjdGVlenku/Y29tL3N5c3RlbS9y/ZXNvdXJjZXMvdGh1/bWJuYWlscy8wNzIv/NTI5LzU2Ny9zbWFs/bC9mcmVzaC1yaXBl/LWd1YXZhLWZydWl0/LXdob2xlLWFuZC1o/YWx2ZWQtdHJvcGlj/YWwtcHJvZHVjZS1o/ZWFsdGh5LWVhdGlu/Zy1wbmcucG5n" ,
	}

	productList = append(productList, prd1)
	productList = append(productList, prd2)
	productList = append(productList, prd3)
	productList = append(productList, prd4)
	productList = append(productList, prd5)
	productList = append(productList, prd6)
}


//middleware

  