package main

import "fmt"

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	
	sum(1, 2)
	sum(1, 2, 3)
	sum(3, 4, 7, 12)

	
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
