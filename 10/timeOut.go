package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	fmt.Println(time.Now())

	go func() {
		time.Sleep(4 * time.Second)
		c1 <- "c1 OK!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "c2 OK!"
	}()

	for {
		select {
		case res := <-c1:
			fmt.Println(time.Now())
			fmt.Println(res)
		case res := <-c2:
			fmt.Println(time.Now())
			fmt.Println(res)
		case <-time.After(2 * time.Second):
			fmt.Println(time.Now())
			fmt.Println("timeout")
			return
		}
	}
}
