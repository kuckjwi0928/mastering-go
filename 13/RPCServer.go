package main

import (
	"fmt"
	"math"
	"net"
	"net/rpc"
	"sharedRPC"
)

type MyInterface struct{}

func Power(x, y float64) float64 {
	return math.Pow(x, y)
}

func (m *MyInterface) MultiPly(arg *sharedRPC.MyFloats, reply *float64) error {
	*reply = arg.A1 * arg.A2
	return nil
}

func (m *MyInterface) Power(arg *sharedRPC.MyFloats, reply *float64) error {
	*reply = Power(arg.A1, arg.A2)
	return nil
}

func main() {
	err := rpc.Register(new(MyInterface))

	if err != nil {
		fmt.Println(err)
		return
	}

	t, err := net.ResolveTCPAddr("tcp4", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
