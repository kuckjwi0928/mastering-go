package main

import (
	"fmt"
)

func main() {
	d1()
	d2()
	d3()
}

func d1() {
	for i := 3; i > 0; i-- {
		defer fmt.Println(i, " ")
	}
}

func d2() {
	for i := 3; i > 0; i-- {
		defer func() {
			fmt.Println(i, " ")
		}()
	}
}

func d3() {
	for i := 3; i > 0; i-- {
		defer func(n int) {
			fmt.Println(n, " ")
		}(i)
	}
}
