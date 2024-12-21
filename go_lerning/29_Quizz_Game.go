package main 

import "fmt"

func main()  {
	fmt.Println("Welcome to the Quizz Game!")
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello! %v , Welcome to the Game\n" , name)
	score := 0
	var age uint 
	fmt.Println("Enter your age: ")
	fmt.Scan(&age)
	if age >= 18 {
		fmt.Println("Yeah you can play! ")
		score++
	} else {
		fmt.Println("Sorry you can't play!")
		return
	}
	fmt.Printf("Question 1 \nWhat do you prefer cold tea or hot coffee: ")
	var answer string
	var answer2 string
	fmt.Scan(&answer, &answer2)

	if answer + " " + answer2 == "hot coffee" {
		fmt.Println("Good choice")
		score++
	} else {
		fmt.Println("You can take an other choice")
	}
	
	fmt.Printf("Question 2 \nHow many days in a week: \n")
	var days uint
	fmt.Scan(&days)

	if days == 7 {
		fmt.Println("Correct answer")
		score++
	} else {
		fmt.Println("you Can take another chace.")
	}
	persent := (float64(score)/ 3.0) * 100
	fmt.Printf("Your score is %v%%!" , persent)
}
