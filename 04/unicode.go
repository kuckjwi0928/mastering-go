package main

import (
	"fmt"
	"unicode"
)

func main() {
	const sL = "\x99\x42\x32\x55\x50\x35\x23\x50\x29\x9c"
	for i := 0; i < len(sL); i++ {
		if unicode.IsPrint(rune(sL[i])) {
			fmt.Printf("%c\n", sL[i])
		} else {
			fmt.Println("Not printable!")
		}
	}
}
