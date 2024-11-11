package main

import "fmt"

func main () {
	var x int = 12
	y := x 
	fmt.Printf("value of x is %v and value of y is %v \n" , x , y )
	y = 7 
	fmt.Printf("value of x is %v and value of y is %v \n" , x , y )// immutable varibles 
	
	var a [] int = [] int {3 , 7 , 12}
	b := a 
	b[1] = 37
	fmt.Println(a , b) // mutable varibles 

	var z [] int = [] int {5 , 13 , 17}
	fmt.Println(z)
	
	ubdateSlice(z)
	fmt.Println(z)
}
func ubdateSlice (slice [] int ) {
	slice[1] = 27
}
