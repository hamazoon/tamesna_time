package main 

import "fmt"

func main()  {
	table := []int {12,3,27,13,17,5,37,7,19,21}
	fmt.Println("Before sorting:" ,table)
	sorting(table)
	fmt.Println("After sorting:" ,table)
}
 func sorting(arr[]int) {
	n := len(arr)
	for i:=0 ; i<n-1 ; i++ {
		min:=i
		for j:=i+1 ; j<n ; j++ {
			if arr[j] < arr[min] {
				min = j 
			}
		}
		if min != i {
			temp := arr[i]
			arr[i] = arr[min]
			arr[min] = temp
			fmt.Println(arr)
		}	
	} 
 }
