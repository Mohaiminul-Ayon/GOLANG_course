package main

import "fmt"

func calculate()(result int){
	result=0

	fmt.Println("1st :",result)
	show:=func(){
		result = result+10
		fmt.Println("defer :",result)
	}

	defer show()

	result =5
	fmt.Println("2nd :",result)

	return
}
func calc()int{
	result:=0

	fmt.Println("first : ",result)
	show:=func(){
		result = result+10
		fmt.Println("defer :",result)
	}

	defer show()

	result =5
	fmt.Println("2nd : ",result)

	return result
}

func sum(a int, b int)(s int){
	s=a+b
	return
}

func main(){
	a:=calculate()
	fmt.Println("A result :",a)
	b:=calc()
	fmt.Println("B result :",b)
}








//defer 1

// func def(){
// 	money:= 10

// 	fmt.Println(money)

// 	money+=10
// 	defer fmt.Println("defer 1st :",money)
	
// 	money=5
// 	fmt.Println(money)
	
// 	money+=10
// 	defer fmt.Println("defer 2nd :",money)


// }

// func call(){
// def()
// }


// func main(){
// 	call()
// }