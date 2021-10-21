package main

import "fmt"

type StackNode struct {
	Value int
	Next  *StackNode
}

var stackSize = 0
var stack = new(StackNode)

func StackPush(v int) bool {
	if stack == nil {
		stack = &StackNode{v, nil}
		stackSize = 1
		return true
	}
	temp := &StackNode{v, nil}
	temp.Next = stack
	stack = temp
	stackSize++
	return true
}

func StackPop(t *StackNode) (int, bool) {
	if stackSize == 0 {
		return 0, false
	}
	if stackSize == 1 {
		stackSize = 0
		stack = nil
		return t.Value, true
	}
	stack = stack.Next
	stackSize--
	return t.Value, true
}

func traverseStack(t *StackNode) {
	if stackSize == 0 {
		fmt.Println("Empty Stack!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func main() {
	stack = nil
	v, b := StackPop(stack)
	if b {
		fmt.Print(v, " ")
	} else {
		fmt.Println("Pop() failed!")
	}
	StackPush(100)
	traverseStack(stack)
	StackPush(200)
	traverseStack(stack)
	for i := 0; i < 10; i++ {
		StackPush(i)
	}
	for i := 0; i < 15; i++ {
		v, b = StackPop(stack)
		if b {
			fmt.Print(v, " ")
		} else {
			break
		}
	}
	fmt.Println()
	traverseStack(stack)
}
