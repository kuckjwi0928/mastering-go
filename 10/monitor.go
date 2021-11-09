package main

import (
	"fmt"
	"sync"
)

var readValue = make(chan int)
var writeValue = make(chan int)

func set(newValue int) {
	writeValue <- newValue
}

func get() int {
	return <-readValue
}

func monitor() {
	var value int
	for {
		select {
		case newValue := <-writeValue:
			value = newValue
			fmt.Printf("%d ", value)
		case readValue <- value:
		}
	}
}

func main() {
	n := 10
	fmt.Printf("Going to create %d random numbers\n", n)

	go monitor()

	var w sync.WaitGroup

	for r := 0; r < n; r++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			set(i)
		}(r)
	}

	w.Wait()

	fmt.Println()
	fmt.Printf("Last value %d\n", get())
}
