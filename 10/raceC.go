package main

import (
	"fmt"
	"sync"
)

var GlobalMutex sync.Mutex

func main() {
	var waitGroup sync.WaitGroup

	k := make(map[int]int)
	k[1] = 12

	for i := 0; i < 10; i++ {
		waitGroup.Add(1)
		go func(_i int) {
			defer waitGroup.Done()
			GlobalMutex.Lock()
			k[_i] = _i
			GlobalMutex.Unlock()
		}(i)
	}

	k[2] = 10
	waitGroup.Wait()

	fmt.Printf("k = %#v\n", k)
}
