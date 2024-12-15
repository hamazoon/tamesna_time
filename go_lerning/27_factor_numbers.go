package main 

import "fmt"

func factor(primes []int , num int ) []int {
	var res []int 
	for _, p := range primes {
		for num % p == 0 {
			res = append(res , p )
			num = num / p 
		}
	}
	if num > 1 {
		res = append(res , num)
	}
	return res
}

func main()  {
	a := factor([]int{2,3,5,7,11,13,17,19} , 9785)
	b := factor([]int{2,3,5,7,9,11,13,17,19} ,120 )
	fmt.Println(a)
	fmt.Println(b)
}
