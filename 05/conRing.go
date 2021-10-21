package main

import (
	"container/ring"
	"fmt"
)

var sz = 10

func main() {
	myRing := ring.New(sz + 1)

	fmt.Println(myRing)
	fmt.Println(myRing.Len())

	for i := 0; i < myRing.Len()-1; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}

	myRing.Value = 2

	fmt.Println(myRing)

	sum := 0
	myRing.Do(func(x interface{}) {
		t := x.(int)
		sum = sum + t
	})

	fmt.Println(sum)

	myRing.Do(func(x interface{}) {
		t := x.(int)
		fmt.Print(t, " -> ")
	})
}
