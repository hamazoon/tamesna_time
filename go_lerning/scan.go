package main

import (
	"fmt"
	"os"
	"bufio"
)
func main(){
	// Create a new scanner that reads from standard input (console)
	scanner := bufio.NewScanner (os.Stdin)
	// Print the prompt for the user
	fmt.Println("enter your name :"  )
	// Scan the next input from the user
	scanner.Scan()
	input := scanner.Text()
	// Print the result with correct formatting
	fmt.Printf("your name is %q\n" , input)

}
