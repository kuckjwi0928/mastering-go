package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	data, err := http.Get("https://google.com")
	defer data.Body.Close()

	if err != nil {
		fmt.Println(err)
		return
	}

	_, _ = io.Copy(os.Stdout, data.Body)
}
