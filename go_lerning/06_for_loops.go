package main

import (
	"fmt"
)
func main(){

	var a []int = []int {1, 3 , 7 ,5 , 12 , 17 ,37 , 92 ,3 , 17}
	for i := 0 ; i < len(a) ; i++ {
		fmt.Println(" element =", a[i])
	}
	for i, elem1 := range a {
		fmt.Printf("the index is %d : the value is %d \n" , i ,elem1)
	}
	for i , elem1 := range a {
		for j , elem2 := range a {
			if elem1 == elem2 && i < j {
				fmt.Println("the values that existe two times are " , elem1)
			}
		}
	}
}
