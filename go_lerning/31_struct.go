package main

import "fmt"

type adress struct {
	city string
	street string
	home_num int
	phone_num int
}

type student struct {
	name string
	s_num int
	age int
	department string
	adress
}

func (s student) PrintInfo()  {
	fmt.Println("from PrintInfo" , s)
}

func main()  {
	var hamza_adress = adress{city: "tamsna", street: "najah", home_num: 56, phone_num: 0634325735}
	var hamza = student{name: "hamza", s_num: 5478, age: 32, department: "go_devloper" ,adress: hamza_adress}
	fmt.Println(hamza)
	fmt.Println(hamza.age)
	hamza.PrintInfo()

}
