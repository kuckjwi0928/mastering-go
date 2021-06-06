package main

import (
	"fmt"
)

type People struct {
	name string
	age  int
}

func main() {
	slice := make([]People, 0)

	slice = append(slice, People{"kuckjwi", 30})
	slice = append(slice, People{"isis", 29})

	fmt.Println(slice)

	_map := make(map[string]int)

	for _, v := range slice {
		_map[v.name] = v.age
	}

	for k, v := range _map {
		fmt.Println(k, v)
	}
}
