package main

import "fmt"

func rev(str string) string {
	var s string
	for i := len(str)-1 ; i >= 0 ; i-- {
		s = s + string(str[i])
	}
	return s
}

func main()  {
	r := rev("cat")
	z := rev("ca t")
	x := rev("hamza")
	
	fmt.Println(r)
	fmt.Println(z)
	fmt.Println(x)
}
