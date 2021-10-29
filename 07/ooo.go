package main

import "fmt"

type root struct {
	o1 o1
	o2 o2
}

type o1 struct {
	XX int
	YY int
}

type o2 struct {
	AA string
	XX int
}

func (o o1) A() {
	fmt.Println("Function A() for o1")
}

func (o o2) A() {
	fmt.Println("Function A() for o2")
}

func main() {
	var i root
	i.o1.A()
	i.o2.A()
}
