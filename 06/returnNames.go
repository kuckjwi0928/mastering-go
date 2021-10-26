package main

import "fmt"

func main() {
	fmt.Println(namedMinMax(5, 10))
	fmt.Println(minMax(5, 10))
}

func namedMinMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return
}

func minMax(x, y int) (min, max int) {
	if x > y {
		min = y
		max = x
	} else {
		min = x
		max = y
	}
	return max, min
}
