package main

import "fmt"

func main() {
	sq := returnPtr(10)
	fmt.Println(*sq)
	fmt.Println(sq)
}

func returnPtr(x int) *int {
	y := x * x
	return &y
}
