package main

import "fmt"

var a=10

func main(){

	age:=30

	if 18<age{
		a:=18
		fmt.Println(a)
	}
	fmt.Println(a)
}