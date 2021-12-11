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

	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	for {
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println(err)
			return
		}
		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting connection!")
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
