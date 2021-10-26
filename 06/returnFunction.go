package main

import "fmt"

func main() {
	i := funReturnFun()
	j := funReturnFun()
	fmt.Println("i1:", i())
	fmt.Println("i2:", i())
	fmt.Println("j1:", j())
	fmt.Println("j2:", j())
}

func funReturnFun() func() int {
	i := 0
	return func() int {
		i++
		return i * i
	}
}
