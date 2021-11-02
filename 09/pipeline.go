package main

import (
	"fmt"
	"math/rand"
	"time"
)

var CLOSEA = false
var DATA = make(map[int]bool)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func first(min, max int, out chan<- int) {
	for {
		if CLOSEA {
			close(out)
			return
		}
		fmt.Println("first out!")
		out <- random(min, max)
	}
}

func second(out chan<- int, in <-chan int) {
	for x := range in {
		fmt.Print("second receive! ", x, " ")
		fmt.Println()
		_, ok := DATA[x]
		if ok {
			CLOSEA = true
		} else {
			DATA[x] = true
			fmt.Println("second out!")
			out <- x
		}
	}
	close(out)
}

func third(in <-chan int) {
	var sum int
	sum = 0
	for x := range in {
		fmt.Println("third receive!")
		sum = sum + x
	}
	fmt.Println(sum)
}

func main() {
	A := make(chan int)
	B := make(chan int)

	rand.Seed(time.Now().UnixNano())

	go first(10, 20, A)
	go second(B, A)
	third(B)
}
