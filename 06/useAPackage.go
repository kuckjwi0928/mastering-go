package main

import (
	"./aPackage"
	"fmt"
)

func main() {
	fmt.Println("Using aPacakge!")
	aPackage.A()
	aPackage.B()
	fmt.Println(aPackage.Myconstant)
}
