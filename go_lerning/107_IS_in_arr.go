package main 

import "fmt"

func main()  {
	a := []int {13,17,12,5,7,9,3}
	fmt.Println(is_in(a ,3))
	b := []int {7,4,8,9,10,14,45}
	fmt.Println(is_in(b ,3))
	c := []int {13,8,-7, -4,-8}
	fmt.Println(is_in(c ,-8))
}

func is_in(arr []int , n int) bool {
	for i := 0 ; i < len(arr) ; i++ {
		if n == arr[i] {
			return true
		}
	}
	return false 
}
