package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args

	var sum int64 = 0

	for i := 1; i < len(arguments); i++ {
		n, _ := strconv.ParseInt(arguments[i], 0, 64)
		sum = n + sum
	}

	fmt.Println("sum:", sum)
}
