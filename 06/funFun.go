package main

import "fmt"

func main() {
	fmt.Println(funFun(function1, 3))
	fmt.Println(funFun(function2, 3))
	fmt.Println(funFun(func(i int) int { return i * i * i }, 3))
}

func function1(i int) int {
	return i + i
}

func function2(i int) int {
	return i * i
}

func funFun(f func(int) int, v int) int {
	return f(v)
}
