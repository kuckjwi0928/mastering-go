package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/trace"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGc:", mem.NumGC)
	fmt.Println("-----")
}

func main() {
	f, err := os.Create("./11/trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()

	var mem runtime.MemStats

	for i := 0; i < 3; i++ {
		s := make([]byte, 50000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}
	printStats(mem)

	for i := 0; i < 5; i++ {
		s := make([]byte, 100000000)
		if s == nil {
			fmt.Println("Operation failed")
		}
	}
	printStats(mem)
}
