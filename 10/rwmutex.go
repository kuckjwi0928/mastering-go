package main

import (
	"fmt"
	"sync"
)

var Password = secret{password: "kuckjwi"}

type secret struct {
	RWM      sync.RWMutex
	M        sync.Mutex
	password string
}

func Change(c *secret, pass string) {
	c.RWM.Lock()
	fmt.Println("LChange")
	c.password = pass
	c.RWM.Unlock()
}

func show(c *secret) string {
	c.RWM.RLock()
	fmt.Print("show ")
	defer c.RWM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Print("showWithLock ")
	defer c.M.Unlock()
	return c.password
}

func main() {
	var waitGroup sync.WaitGroup

	for i := 0; i < 15; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Print("Pass: ", show(&Password))
			fmt.Println()
			fmt.Print("Pass: ", showWithLock(&Password))
			fmt.Println()
		}()
		go func() {
			waitGroup.Add(1)
			defer waitGroup.Done()
			Change(&Password, "123456")
		}()
	}

	waitGroup.Wait()
}
