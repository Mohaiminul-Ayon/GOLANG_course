package main

import (
	"fmt"
	"net/http"
)

func HelloHandeler(w http.ResponseWriter,r *http.Request){//response writer-w and r-  response
	fmt.Fprintln(w, "Hello World")
}
func adminHandeler(w http.ResponseWriter,r *http.Request){//response writer-w and r-  response
	fmt.Fprintln(w, "I'm Ayon. I'm a great Engineer,I'm Software engineer")
}

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("/hello",HelloHandeler)
	mux.HandleFunc("/admin",adminHandeler)

	fmt.Println("Server running on : 3000")

	err:= http.ListenAndServe(":3000",mux)//"Faild"
	if err!= nil{
		fmt.Println("Error starting the server",err)
	}
}