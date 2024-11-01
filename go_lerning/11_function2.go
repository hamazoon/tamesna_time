package main 

import "fmt"

func main () {
	x := test
	x(7)
	y := func(b int){
		fmt.Println("on the run" , b)
	}
	y(5)
	test2(addten)
	test2(multen)
	return_func("hamza")()
}
func test2 (myfunc func ( x int)int ){
	fmt.Println(myfunc (3))
}
func addten (a int) int {
	return a + 10
}
func multen ( b int) int {
	return b * 10
}
func test (a int){
	fmt.Println("test function " , a)
}
func return_func (s string) func() {
	return func (){
		fmt.Println("hello" , s)
	}
}
