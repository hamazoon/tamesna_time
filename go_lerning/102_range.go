package main 

import (
	"fmt"
	"strconv"
)
func main()  {
	var my_map = map[string]string {"id":"0231" , "name":"hamza", "age":"32" , "salary":"2000" ,"deparment":"Go"}
	for key,value := range my_map {
		fmt.Println("key is " , key , "value is " , value)
	}
	var a = 12
	var b = 17

	loop: 
	fmt.Println("a is smuller than b")

	if ( a > b ) {
		goto loop
	} else {
		a++
		fmt.Println("a is not smuller than b")
	}
	var x int = 2
	var y float64 = 3.14
	var z string = "12"
	var c string = "17.75"

	var myfloatnum1 = float64(x)
	var myintnum2 = int(y)
	fmt.Printf(" x type is %T \n y type is %T \n z type is %T \n c type is %T \n ", x, y, z, c)
	fmt.Printf("myfloatnum1 type is %T and value is %v \n myintnum2  type is %T and value is  ", myfloatnum1,myfloatnum1, myintnum2 ,myintnum2 )

	var num1, _= strconv.ParseFloat(c, 64)
	var num2, _ = strconv.ParseInt(z, 10, 64)
	fmt.Printf("\nnum1 type is %T and value is %v \n num2 type is %T and value is %v ", num1, num1, num2, num2)
}
