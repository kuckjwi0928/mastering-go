package main

import (
	"fmt"
	"time"
)

func main() {
	myDate := "2021-06-06"
	d, err := time.Parse("2006-02-06", myDate)

	if err == nil {
		fmt.Println("Full:", d)
		fmt.Println("Time:", d.Day(), d.Month(), d.Year())
	} else {
		fmt.Println(err)
	}
}
