package main 

import "fmt"

func Sum(list []int) int  {
	s := 0
	for _, i := range list {
		s += i
	}
	return s
}

func R_Sum(arr[]int) int {
	if len(arr) == 0 {
		return 0
	}
	return arr[0] + R_Sum(arr[1:])

}

func main()  {
	a := []int {2,5,7,9,17,12}
	fmt.Println(Sum(a))
	b := []int {3,7,8,6,-7,-5}
	fmt.Println(Sum(b))
	c := []int {3,5,9,7,4}
	fmt.Println(R_Sum(c))
	d := []int {4,6,9,6,-7,-5,0}
	fmt.Println(R_Sum(d))
}
