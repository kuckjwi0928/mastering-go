package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	v1 := "123"
	v2 := 123
	v3 := "Have a nice day"
	v4 := "abc"

	fmt.Println(v1 ,v2 ,v3 ,v4)
	_, _ = io.WriteString(os.Stdout, "kuckjwi")
}
