package main

import (
	"fmt"
)

func main() {
	var arr [5]string // declaration of an empty array of strings
	fmt.Println("arr =", arr)

	var arr1 [5]int // declaration of an empty array of integers
	fmt.Println("arr1 =", arr1)

	arr1[3] = 92   // assign 92 to the fourth element of arr1
	arr1[1] = 37   // assign 37 to the second element of arr1
	fmt.Println("arr1 =", arr1)
	fmt.Println("length of arr1 =", len(arr1))
	fmt.Println("arr1[3] =", arr1[3]) // print the element arr1[3]

	arr2 := [3]int{23, 12, 7}
	fmt.Println("arr2 =", arr2)
	fmt.Println("length of arr2 =", len(arr2)) // print length of arr2

	for i := 0; i < len(arr2); i++ { // corrected for loop syntax
		fmt.Println("the elements of arr2 =", arr2[i])
	}

	arr2D := [3][2]int{{1, 2}, {3, 4}, {5, 6}} // two-dimensional array
	fmt.Println("arr2D =", arr2D)
	fmt.Println("arr2D[2][1] =", arr2D[2][1])
}
