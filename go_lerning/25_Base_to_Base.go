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

func BaseToDec(value string, base int) int {
	const charset = "0123456789ABCDEF"
	mul := 1
	res := 0
	
	for i := len(value) - 1 ; i >= 0 ; i-- {
		index := -1 
		for j,char := range charset {
			if char == rune(value[i]) {
				index = j 
				break
			}
		}
		res = res + mul * index
		mul = mul * base
	}
	return res 
}

func BaseToBase(value string , base , nwbase int ) string {
	dec := BaseToDec(value , base)
	return DecToBase(dec , nwbase)
}

func main()  {
	a := BaseToBase("1010" ,2 , 8)
	b := BaseToBase("10",16,2)
	c := BaseToBase("20",16,8)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
