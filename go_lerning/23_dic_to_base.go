package main

import "fmt"

func DecToBase(dec , b int) string  {
	const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	var res string
	for dec	> 0 {
		mod := dec % b 
		res = string(charset[mod]) +res
		dec = dec / b 
	}
	return res
}

func main()  {
	a := DecToBase(15 ,2)
	b := DecToBase(17 ,8)
	c := DecToBase(156 ,60)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
