package main 

import "fmt"

func main () {
	s := [] int {5,13,3,17,19,7,12}
sorting(s)
}
func sorting (arr[] int) {
	n := len(arr)
	for i := 0 ; i < n-1 ; i++ {
		flag := 0
		for j := 0 ; j < n-i-1 ; j++ {
			if arr[j] > arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
				flag = 1
			}
		}
		if flag == 0 {
			break
		}
	} 
	fmt.Println(arr)
}
