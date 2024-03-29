package main

import "fmt"

func main() {
	iMap := make(map[string]int)
	iMap["k1"] = 12
	iMap["k2"] = 13
	fmt.Println("iMap:", iMap)

	anotherMap := map[string]int{
		"k1": 12,
		"k2": 13,
	}
	fmt.Println("anotherMap:", anotherMap)

	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	delete(anotherMap, "k1")
	fmt.Println("anotherMap:", anotherMap)

	_, ok := iMap["k1"]

	if ok {
		fmt.Println("Exists!")
	} else {
		fmt.Println("Does Not exist")
	}

	for k, v := range iMap {
		fmt.Println(k, v)
	}
}
