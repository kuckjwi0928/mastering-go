package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	s, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := l.Accept()

	for {
		buffer := make([]byte, 1024)
		if err != nil {
			fmt.Println(err)
			return
		}
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting TCP server!")
			conn.Close()
			return
		}
		fmt.Println(string(buffer[0 : n-1]))
		_, err = conn.Write(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
