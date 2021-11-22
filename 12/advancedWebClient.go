package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
)

func main() {
	c := &http.Client{
		Timeout: 15 * time.Second,
	}
	req, err := http.NewRequest("GET", "https://google.com", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	httpData, err := c.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Status Code:", httpData.Status)
	header, _ := httputil.DumpResponse(httpData, false)
	fmt.Print(string(header))
	characterSet := strings.SplitAfter(httpData.Header.Get("Content-Type"), "charset=")
	fmt.Println(characterSet)
	length := 0
	var buffer [1024]byte
	r := httpData.Body
	for {
		n, err := r.Read(buffer[0:])
		if err != nil {
			break
		}
		length = length + n
	}
	fmt.Println("Calculated response data length:", length)
}
