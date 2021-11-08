package main

import (
	"fmt"
	"sync"
)

var mutex sync.Mutex

func main() {
	var w sync.WaitGroup

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)

	go func() {
		defer w.Done()
		function()
	}()
	w.Add(1)
	w.Wait()
}

func function() {
	mutex.Lock()
	fmt.Println("Locked!")
}
