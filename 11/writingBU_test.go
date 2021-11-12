package main

import (
	"fmt"
	"os"
	"testing"
)

var ERR error

func Benchmark_Create(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Create("./output", 1, 10000000)
	}
	ERR = err
	err = os.Remove("./output")
	if err != nil {
		fmt.Println(err)
	}
}

func Benchmark_Create2(b *testing.B) {
	var err error
	for i := 0; i < b.N; i++ {
		err = Create("./output", 2, 10000000)
	}
	ERR = err
	err = os.Remove("./output")
	if err != nil {
		fmt.Println(err)
	}
}
