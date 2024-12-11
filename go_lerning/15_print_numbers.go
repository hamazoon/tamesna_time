package main

import "fmt"

func main()  {
	for i := 0 ; i < 10 ; i++ {
		for j := 0 ; j < 10 ; j++ {
			for x := 0 ; x < 10 ; x++ {
				for y := 0 ; y < 10 ; y ++ {
					if i == x && j < y {
						printNumbers(i, j, x, y)
						ifNumbers(i, j, x, y)
					} else if i < x {
						printNumbers(i, j, x, y)
						fmt.Print(" ,")
					}
				}
			}
		}
	}
	
}
func printNumbers(i, j, x, y int)  {
	fmt.Print(i)
	fmt.Print(j)
	fmt.Print(" ")
	fmt.Print(x)
	fmt.Print(y)
	
}
func ifNumbers(i, j, x, y int)  {

	if i == 9 && j == 8{
		fmt.Print("\n")
	} else {
		fmt.Print(" ,")
	}
	
}
