package main

import (
	"fmt"
)
func main () {
	var num1 = 3
	num2 := 4
	var num3 = num1
	if num1 < num2{
		fmt.Println("yes")
	}else {
		fmt.Println("no")
	}
	if num2 > num3 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if num1 == num3 {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if (num1 <= num2) && (num2 > num1) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if (num1 >= num2) || (num2 < num3) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
	if !((num1 <= num3) || (num2 >= num1)) {
		fmt.Println("yes")
	} else {
		fmt.Println("no")
	}
}
