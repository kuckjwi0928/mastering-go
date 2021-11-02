package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func walk(path string, info os.FileInfo, err error) error {
	fileInfo, err := os.Stat(path)
	if err != nil {
		return err
	}

	mode := fileInfo.Mode()

	if mode.IsRegular() {
		fmt.Println("+", path)
		return nil
	}

	if mode.IsDir() {
		fmt.Println("*", path)
		return nil
	}

	fmt.Println(path)
	return nil
}

func main() {
	err := filepath.Walk("./08", walk)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
