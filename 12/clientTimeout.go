package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"time"
)

var timeout = time.Millisecond * 500

func TimeoutContext(ctx context.Context, network, host string) (net.Conn, error) {
	conn, err := net.DialTimeout(network, host, timeout)
	if err != nil {
		return nil, err
	}
	_ = conn.SetDeadline(time.Now().Add(timeout))
	return conn, nil
}

func main() {
	t := http.Transport{
		DialContext: TimeoutContext,
	}
	client := http.Client{
		Transport: &t,
	}
	data, err := client.Get("https://google.com")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer data.Body.Close()

	_, err = io.Copy(os.Stdout, data.Body)
	if err != nil {
		fmt.Println(err)
	}
}
