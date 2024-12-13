package main 

import "fmt"

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
func main ()  {
	a := BaseToDec("123" , 4)
	b := BaseToDec("10101" , 2)
	c := BaseToDec("11" , 16)
	d := BaseToDec("12" , 8)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
