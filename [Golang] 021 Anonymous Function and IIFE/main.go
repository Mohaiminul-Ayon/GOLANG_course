package main

import "fmt"




func main(){
	fmt.Println("start")
	func(a int,b int){//anonymous function
		c:=a+b
		fmt.Println(c)
	}(5,7)//IIFE

	a:=10


}