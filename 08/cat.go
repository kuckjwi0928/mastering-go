package main

import (
	"bufio"
	"io"
	"os"
)

func printFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		_, _ = io.WriteString(os.Stdout, scanner.Text())
		_, _ = io.WriteString(os.Stdout, "\n")
	}
	return nil
}

func main() {
	filename := "./08/data.txt"
	_ = printFile(filename)
}
