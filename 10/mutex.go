package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	m  sync.Mutex
	v1 int
)

func change(i int) {
	m.Lock()
	time.Sleep(time.Second)
	v1 = v1 + 1
	if v1%10 == 0 {
		v1 = v1 - 10*i
	}
	m.Unlock()
}

func read() int {
	m.Lock()
	a := v1
	m.Unlock()
	return a
}

func main() {
	var waitGroup sync.WaitGroup
	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(i int) {
			defer waitGroup.Done()
			change(i)
			fmt.Printf("-> %d ", read())
		}(i)
	}
	waitGroup.Wait()
	fmt.Printf("-> %d ", read())
}
