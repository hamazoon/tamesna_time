package main 

import "fmt"

func main ()  {
	a := []int {21,3,37,17,12,5,7,25,19,13}
	fmt.Println(a)
	sorting(a)
	fmt.Println(a)
}
func sorting(arr[]int)  {
	n := len(arr)
	for i := 1 ; i < n ; i++ {
		temp := arr[i]
		j := i-1
		for j >= 0 && arr[j] > temp {
			arr[j+1] = arr[j]
			j--
			fmt.Println(arr)
		}
		arr[j+1] = temp
	}
}
