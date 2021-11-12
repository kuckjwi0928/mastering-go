package main

import (
	"fmt"
)

func main() {
	a()
	fmt.Println("main() ended")
}

func a() {
	fmt.Println("Inside a()")
	defer func() {
		if c := recover(); c != nil {
			fmt.Println("Recover inside a()")
			err, ok := c.(error)
			if ok {
				fmt.Println(err)
			}
		}
	}()
	fmt.Println("About to call b()")
	b()
	fmt.Println("b() exited!")
	fmt.Println("Exiting a()")
}

func b() {
	fmt.Println("Inside b()")
	panic("Panic in b()")
	fmt.Println("Exiting b()")
}
