package main

import (
	"fmt"
	"time"
)

func main() {
	t := "12:10:31"
	// 15 = hour
	// 04 = minute
	// 05 = second
	d, err := time.Parse("15:04:05", t)

	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Hour(), d.Minute(), d.Second())
	} else {
		fmt.Println(err)
	}
}
