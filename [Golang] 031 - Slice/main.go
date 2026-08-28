package main

import "fmt"

func main(){
//Topic
//1.Slice From an existing Array
//2.Slice From an Slice 
/*
	arr:=[6]string{"This","is","a","Go","Interview","Question"}
	fmt.Println(arr)

	s:=arr[1:4] //[is a go]
	//len=3 ; capacity= 5 (jei ghor theke dhorbo tar last porjonto ) ; ptr= *arr[1] 1st ghor er address
	fmt.Println(s)

	fmt.Println(len(s))
	fmt.Println(cap(s))
	
	s1:=s[1:2]//a
	fmt.Println(s1)
	fmt.Println(len(s1)) // 1
	fmt.Println(cap(s1))//4 keno? = cause ekhane address rakhe arr er adress ke point kore tai capacity arr er hisbe hishab hoy slice er hishabe na


*/
//3.slice litaral
	s:=[]int {1,2,5}//slice literal
	fmt.Println(s)
//4.slice with make lenght
//5.slice with make lenght and capacity

s1:=make([]int/*type*/,3/*length*/,5/*capacity*/)
// s1:=make([]int/*type*/,3/*length*/)
fmt.Println(s1)
fmt.Println(len(s1))
fmt.Println(cap(s1))


//6.empty slice and initial

var s2[]int //empty slice

fmt.Println(s2)
fmt.Println(len(s2))
fmt.Println(cap(s2))



s2=append(s2,1,2,3)

fmt.Println(s2)
fmt.Println(len(s2))
fmt.Println(cap(s2))
	


	



}







/*
note
=============

slice litaral : jei array te size define kora nai
*/