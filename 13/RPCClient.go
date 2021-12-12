package main

import (
	"fmt"
	"net/rpc"
	"sharedRPC"
)

func main() {
	c, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	args := sharedRPC.MyFloats{A1: 16, A2: -0.5}

	var reply float64

	err = c.Call("MyInterface.MultiPly", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Multiply): %f\n", reply)

	err = c.Call("MyInterface.Power", args, &reply)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("Reply (Power): %f\n", reply)
}
