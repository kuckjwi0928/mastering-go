package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Epoch time:", time.Now().Unix())
	t := time.Now()
	fmt.Println(t, t.Format(time.RFC3339))
	fmt.Println(t.Weekday(), t.Day(), t.Month(), t.Year())

	time.Sleep(time.Second)
	t1 := time.Now()
	fmt.Println("Time differenct:", t1.Sub(t))
}
