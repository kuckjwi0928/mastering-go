package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	client := http.Client{
		Timeout: time.Microsecond,
	}
	data, err := client.Get("https://google.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer data.Body.Close()
	_, _ = io.Copy(os.Stdout, data.Body)
}
