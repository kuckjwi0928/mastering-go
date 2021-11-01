package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
	var err error

	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		for _, x := range line {
			fmt.Println(string(x))
		}
	}

	return nil
}

func main() {
	err := charByChar("./08/test.txt")
	if err != nil {
		fmt.Println(err)
	}
}
