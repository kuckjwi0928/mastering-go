package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("You ar using ", runtime.Compiler, " ")
	fmt.Println("on a", runtime.GOARCH)
	fmt.Println("with Go version", runtime.Version())
}
