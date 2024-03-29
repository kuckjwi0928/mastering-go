package main

import "fmt"

func main() {
	var myInt interface{} = 123
	k, ok := myInt.(int)
	if ok {
		fmt.Println("Success:", k)
	}

	v, ok := myInt.(float64) // Prevents panic
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("Failed without panicking!")
	}

	i := myInt.(int)
	fmt.Println("No checking:", i)

	j := myInt.(bool) // panic
	fmt.Println(j)
}
