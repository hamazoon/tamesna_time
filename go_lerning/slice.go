package main

import "fmt"

func main(){

	var x []int = []int {3, 7, 12, 17, 37, 92}
	var s []int = x[1:3]// inisialation of a slice 

	fmt.Println("slice from index 1 to 3 = " , s)
	fmt.Println("the  length of s =" , len(s))// print the length of s 

	fmt.Println("the capasity of s = " , cap(s))// print the capasity of s 
	fmt.Println("s[1:cap(s)] = ", s[1:cap(s)])
	
	n := append(x , 10) //adding an itme to x 
	fmt.Println("the new n = " , n)
}
