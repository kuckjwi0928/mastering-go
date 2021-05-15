package main

import "fmt"

func main() {
	// for
	for i := 0; i < 100; i++ {
		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Println(i, " ")
	}

	fmt.Println()

	// while
	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Println(i, " ")
		i--
	}

	// do-while
	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}
		fmt.Println(i, " ")
		i++
	}

	// array := [size]type{elements}
	anArray := [5]int{0, 1, -1, 2, -2}
	// index and value
	for i, value := range anArray {
		fmt.Println("index:", i, "value:", value)
	}

	// only value
	for _, value := range anArray {
		fmt.Println("value:", value)
	}

	// length loop
	for i := 0; i < len(anArray); i++ {
		fmt.Println(anArray[i])
	}
}
