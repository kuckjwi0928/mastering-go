package main

import "fmt"

type squ struct {
	X float64
}

type cir struct {
	R float64
}

type rec struct {
	X float64
	Y float64
}

func tellInterface(x interface{}) {
	switch v := x.(type) {
	case squ:
		fmt.Println("This is a square!")
	case cir:
		fmt.Printf("%v is a circle!\n", v)
	case rec:
		fmt.Println("this is a rectangle")
	default:
		fmt.Printf("Unkown type %T\n", v)
	}
}

func main() {
	x := cir{R: 10}
	tellInterface(x)
	y := rec{X: 4, Y: 1}
	tellInterface(y)
	z := squ{X: 4}
	tellInterface(z)
	tellInterface(10)
}
