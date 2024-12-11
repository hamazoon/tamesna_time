package main

import (
	"fmt"
)
func main () {
	var num1 = 3
	num2 := 4
	var num3 = num1
	if num1 < num2{
		num2 += num3 
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("yes")
	}else {
		num1 -= num2
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("no")
	}
	if num2 > num3 {
		num3 *= num2
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("yes")
	} else {
		num2 /= num1
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("no")
	}
	if num1 == num3 {
		num3 %= num1
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("yes")
	} else {
		num1 += num2
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("no")
	}
	if (num1 <= num2) && (num2 > num1) {
		num2 -= num3
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("yes")
	} else {
		num3 *= num1
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("no")
	}
	if (num1 >= num2) || (num2 < num3) {
		num1 /= num2
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("yes")
	} else {
		num2 %=num3
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("no")
	}
	if !((num1 <= num3) || (num2 >= num1)) {
		num3 += num1
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("yes")
	} else {
		num1 -= num2
		fmt.Printf("num1 = %v , num2 = %v num3 %v \n" , num1, num2,num3 )
		fmt.Println("no")
	}
}
