package main

import "fmt"

func main() {
	x := 12.2
	fmt.Println(getPtr(&x))
}

func getPtr(v *float64) float64 {
	return *v * *v
}
