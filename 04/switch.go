package main

import "fmt"

func main() {
	number := 1

	switch {
	case number < 0:
		fmt.Println("Less than zero!")
	case number > 0:
		fmt.Println("Bigger than zero!")
	default:
		fmt.Println("Zero!")
	}

	name := "kuckjwi"
	switch name {
	case "kuckjwi":
		fmt.Println("Hi kuckjwi!")
		fallthrough
	default:
		fmt.Println("Who are you?")
	}
}
