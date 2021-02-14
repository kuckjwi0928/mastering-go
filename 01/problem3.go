package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var f *os.File
	f = os.Stdin
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if strings.ToUpper(scanner.Text()) == "STOP" {
			fmt.Println("Stopping!")
			break
		}
		n, _ := strconv.ParseInt(scanner.Text(), 0, 64)
		fmt.Println(">", n)
	}
}
