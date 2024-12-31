package main

import "fmt"

func GCD(a, b int) int {
	for b != 0 {
		temp := a
		a = b 
		b = temp % b
	}
	return a
}

func main()  {
	c := GCD(24, 36)
	d := GCD(8, 15)
	g := GCD(120, 24)
	e := GCD(13, 19)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(g)
	fmt.Println(e)
}
