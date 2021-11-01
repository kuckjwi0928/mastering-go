package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func woreByword(file string) error {
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

		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for _, word := range words {
			fmt.Println(word)
		}
	}

	return nil
}

func main() {
	err := woreByword("./08/test.txt")
	if err != nil {
		fmt.Println(err)
	}
}
