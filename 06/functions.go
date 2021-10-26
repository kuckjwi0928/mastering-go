package main

import "fmt"

func main() {
	square := func(s int) int {
		return s * s
	}
	fmt.Println(square(2))
	double := func(s int) int {
		return s + s
	}
	fmt.Println(double(2))
	d, s := doubleSquare(2)
	fmt.Println(d, s)
}

func doubleSquare(x int) (int, int) {
	return x * 2, x * x
}
