package main

import "fmt"

func main() {
	const (
		_ int = 1 << iota
		_
		p4_0
		_
		p4_2
		_
		_
		_
		p4_4
		_
		_
		_
		p4_6
	)
	fmt.Println("4^0:", p4_0)
	fmt.Println("4^2:", p4_2)
	fmt.Println("4^4:", p4_4)
	fmt.Println("4^6:", p4_6)
}
