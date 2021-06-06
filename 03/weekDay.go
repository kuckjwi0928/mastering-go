package main

import (
	"fmt"
	"time"
)

func main() {
	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println(Sunday == time.Sunday)
	fmt.Println(Monday == time.Monday)
}
