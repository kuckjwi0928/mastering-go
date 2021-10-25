package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var i int64 = 1
	var LENGTH int64 = 10
	SEED := time.Now().Unix()
	rand.Seed(SEED)
	startChar := "!"
	for {
		myRand := rand.Intn(94-0) + 94
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		if i == LENGTH {
			break
		}
		i++
	}
	fmt.Println()
}
