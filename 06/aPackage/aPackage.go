package aPackage

import "fmt"

const Myconstant = 123     // public
const privateConstant = 21 // private

func A() {
	fmt.Println("This is function A!")
}

func B() {
	fmt.Println("privateConstant:", privateConstant)
}
