package main 

import (
	"fmt"
	"strings"
)

func main()  {
	name := "Hamza"
	arr := []string {"appel","tea","camp"}
	fmt.Println(len(name))
	fmt.Println(strings.ToUpper(name))
	fmt.Println(strings.ToLower(name))
	fmt.Println(strings.HasPrefix(name, "Ha"))
	fmt.Println(strings.HasSuffix(name, "za"))
	fmt.Println(strings.Join(arr , "-"))
	fmt.Println(strings.Repeat(name, 4))
	fmt.Println(strings.Contains(name, "mz"))
	fmt.Println(strings.Index(name, "za"))
	fmt.Println(strings.Count(name, "a"))
	fmt.Println(strings.Replace(name, "a", "o", 2))
	fmt.Println(strings.Compare("hamza", "hamza"))
	fmt.Println(strings.Compare("Hamza", "hamza"))
	fmt.Println(strings.ContainsAny(name, "z"))
	fmt.Println(strings.ContainsAny(name, "f"))

}
