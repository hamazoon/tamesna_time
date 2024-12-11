package main 

import "fmt"

func fizzbuzz(n int)  {
	for i := 1 ; i <= n ; i++ {
		if (i % 3 == 0) && !(i % 5 == 0) {
			fmt.Printf("Fizz ,")
		} else if (i % 5 == 0) &&!(i % 3 == 0) {
			fmt.Printf("Buzz ,")
		} else if (i % 3 == 0 ) && (i % 5 == 0) {
			fmt.Printf("Fizz Buzz ,")
		} else {
			fmt.Printf("%v ,", i)
		}
	}
	fmt.Println()
}

func main()  {
	fizzbuzz(17)
	fizzbuzz(30)
}
