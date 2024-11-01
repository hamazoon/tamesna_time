package main

import "fmt"

func main () {
test(1 , "thomas")
test(3 , "mohamed")
test(5 , "hamza")

cunt := sum(5,17)
fmt.Println("the result is " ,cunt)

count1 , count2 := add(12,17)
fmt.Println( "functoin add " , count1 , " : ",count2) 
}

func test(x int , s string) { 	// a function that take two values 
	fmt.Println("test function" , x  ,s)
	
}
func sum (a int , b int ) int { 	//a function that return one value 
	c := a + b 
	return c
}
func add (x,y int) (int ,int ) { 	//a function that return two values
	defer fmt.Println("defer is here ") 	//defer steatment  
	z := x+y
	w := x*y 
	fmt.Println("the last word in the function ")
return z , w 
}
