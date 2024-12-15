package main 

import "fmt"

func fibonacci(num int) int {
	if num <= 1 {
		return num
	}
	return fibonacci(num - 1) + fibonacci(num - 2)
}

func main()  {
	a := fibonacci(5)
	b := fibonacci(11)
	c := fibonacci(16)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	
}
