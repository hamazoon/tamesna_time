package main

import "fmt"

func sum(a , b int)  int {
	sum := a + b 
	return sum
}

func sub(c , d  int)  int {
	sub := c - d 
	return sub 
}

func mul(e , f  int ) int {
	mul := e * f 
	return mul
}

func div (x , y  int) int {
	div := x / y 
	return div 
	
}

func mod(m , n  int) int {
	mod := m % n 
	return mod 
	
}

func main () {
	fmt.Println(sum(5,6))
	fmt.Println(sum(4,-7))
	fmt.Println(sub(4,3))
	fmt.Println(sub(-3,-5))
	fmt.Println(sub(5,-9))
	fmt.Println(mul(3,5))
	fmt.Println(mul(4,-2))
	fmt.Println(mul(-7,-5))
	fmt.Println(div(4,2))
	fmt.Println(div(5,-3))
	fmt.Println(div(-7,-9))
	fmt.Println(mod(6,5))
	fmt.Println(mod(-9,4))
	fmt.Println(mod(-5,-8))
}
