package main

import (
	"context"
	"fmt"
	"time"
)

func context1(t int) {
	c1 := context.Background()
	c1, cancel := context.WithCancel(c1)
	defer cancel()
	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c1.Done():
		fmt.Println("context1():", c1.Err())
		return
	case r := <-time.After(time.Duration(t) * time.Second):
		fmt.Println("context1():", r)
		return
	}
}

func context2(t int) {
	c2 := context.Background()
	c2, cancel := context.WithTimeout(c2, time.Duration(t)*time.Second)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c2.Done():
		fmt.Println("context2():", c2.Err())
		return
	case r := <-time.After(time.Duration(2*t) * time.Second):
		fmt.Println("context2():", r)
		return
	}
}

func context3(t int) {
	c3 := context.Background()
	deadline := time.Now().Add(time.Duration(t) * time.Second)
	c3, cancel := context.WithDeadline(c3, deadline)
	defer cancel()

	go func() {
		time.Sleep(4 * time.Second)
		cancel()
	}()

	select {
	case <-c3.Done():
		fmt.Println("context3():", c3.Err())
		return
	case r := <-time.After(time.Duration(t*2) * time.Second):
		fmt.Println("context3():", r)
		return
	}
}

func main() {
	context1(1)
	context2(2)
	context3(2)
}
