package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	s := []byte("Data to write\n")
	f1, err := os.Create("./08/data.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f1.Close()
	fmt.Fprintf(f1, string(s))

	f2, err := os.Create("./08/data2.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	defer f2.Close()
	n, err := f2.WriteString("kuckjwi")
	fmt.Printf("wrote %d bytes\n", n)

	f3, err := os.Create("./08/data3.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	w := bufio.NewWriter(f3)
	n, _ = w.WriteString("isis")
	fmt.Printf("wrote %d bytes\n", n)
	w.Flush()

	f4 := "./08/data4.txt"
	err = ioutil.WriteFile(f4, []byte("kuckjwi kuckjwi"), 0644)
	if err != nil {
		fmt.Println(err)
		return
	}

	f5, err := os.Create("./08/data5.txt")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	n, err = io.WriteString(f5, "kuckjwi5555555")
	if err != nil {
		fmt.Println("Cannot create file", err)
		return
	}
	fmt.Printf("wrote %d bytes\n", n)
}
