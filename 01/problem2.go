package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	var sum float64 = 0

	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseFloat(arguments[i], 64)
		sum = n + sum
	}

	fmt.Println("avg:", sum/float64(len(arguments)-1))
}
