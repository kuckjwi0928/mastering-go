package main

import (
	"fmt"
	"io"
	"os"
)

func readSize(f *os.File, size int) []byte {
	buffer := make([]byte, size)

	n, err := f.Read(buffer)
	if err == io.EOF {
		return nil
	}

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return buffer[0:n]
}

func main() {
	f, err := os.Open("./08/test.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	for {
		readData := readSize(f, 4)
		if readData != nil {
			fmt.Print(" ", string(readData))
		} else {
			break
		}
	}
}
