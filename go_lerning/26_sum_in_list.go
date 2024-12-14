package main 
 
import "fmt"

func SumInList(list []int, sum int) (int, int) {
	for i ,val1 := range list {
		for j , val2 := range list {
			if i == j {
				continue 
			}
			if val1 + val2 == sum {
				return val1 ,val2
			}
		}
	}
	return -1,-1
}

func main()  {
	a := []int {2,3,7,5,9,12,13,17,15,19,23,27,32,37,21,25}
	b ,c := SumInList(a , 50)
	fmt.Println(c,b)
}
