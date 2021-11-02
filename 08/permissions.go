package main

import (
	"fmt"
	"os"
)

func main() {
	info, _ := os.Stat("./08/dataFile.gob")
	mode := info.Mode()
	fmt.Println(mode)
	fmt.Println("mode is", mode.String()[1:len(mode.String())])
}
