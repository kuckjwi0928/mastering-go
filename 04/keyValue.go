package main

import "fmt"

type Element struct {
	Name    string
	SurName string
	Id      string
}

var DATA = make(map[string]Element)

func ADD(k string, n Element) bool {
	if len(k) < 1 {
		return false
	}

	if LOOKUP(k) == nil {
		DATA[k] = n
		LOG(fmt.Sprintf("Add Data %v", DATA[k]))
		return true
	}

	return false
}

func DELETE(k string) bool {
	if LOOKUP(k) != nil {
		prev := DATA[k]
		delete(DATA, k)
		LOG(fmt.Sprintf("Deleted Data %v", prev))
		return true
	}
	return false
}

func LOOKUP(k string) *Element {
	_, ok := DATA[k]
	if ok {
		n := DATA[k]
		return &n
	}
	return nil
}

func CHANGE(k string, n Element) bool {
	prev := DATA[k]
	DATA[k] = n
	LOG(fmt.Sprintf("Changed Data %v to %v", prev, n))
	return true
}

func PRINT() {
	for k, v := range DATA {
		fmt.Printf("key: %s value: %v\n", k, v)
	}
}

func LOG(message string) {
	fmt.Println("[LOG]:", message)
}

func main() {
	ADD("123", Element{"kuckhwan", "cho", "kuckjwi"})
	PRINT()
	CHANGE("123", Element{"slave", "tester", "nope"})
	PRINT()
	DELETE("123")
}
