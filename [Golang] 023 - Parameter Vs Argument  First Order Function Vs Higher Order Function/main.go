package main

import "fmt"
func sum(){
	add(2,3)
}

func add(p int, q int){//perameter a , b
	c:=p+q
	fmt.Println(c)
}
func call()func(p int, q int){
	return add
}

func processOperation(a int,b int ,op func(x int ,y int) ){
	op(a,b)
}


func main(){
	// sum()
	// add(3,4)//argument 3,4
	summ:=call()//eta call funtion ke call korse jeta summ e add function ke return korse then oi add function == summ hoi gese so sum kei add hishebe use kora gese
	
	summ(2,3)
	processOperation(2,5,add)
	
	
	

	
}