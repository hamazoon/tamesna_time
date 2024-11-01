package main

import "fmt"

func main () {
	var mp map[string]int = map[string]int {	//create a map 
		"ahmad" : 17,
		"john" : 12,
		"zaid" : 19,
	}
	//mp1 = make(map[string]int)  other way to cearte a map 
	
	fmt.Println("my map =" ,mp)
	
	fmt.Println("value of ahmad =" ,mp["ahmad"])

	mp["john"] = 15 //modify the value of john key 
	fmt.Println("my map =" ,mp)

	mp["hamza"] = 20 	// adding new element 
	fmt.Println("my map =" ,mp)

	delete(mp, "zaid")	// delete an element 
	fmt.Println("my map =" ,mp)

	val, ok := mp["khalid"]
	if ok {
		fmt.Println("value is ", val)
	} else {
		fmt.Println("the value is NON")
	}
	val1, ok := mp["hamza"]
	if ok {
		fmt.Println("value is ", val1)
	} else {
		fmt.Println("the value is NON")
	}

}
