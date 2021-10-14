package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

var root = new(ListNode)

func addNode(t *ListNode, v int) int {
	if root == nil {
		t = &ListNode{v, nil}
		root = t
		return 0
	}
	if v == t.Value {
		fmt.Println("Node already exist:", v)
		return -1
	}
	if t.Next == nil {
		t.Next = &ListNode{v, nil}
		return -2
	}
	return addNode(t.Next, v)
}

func traverseLinkedList(t *ListNode) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	for t != nil {
		fmt.Printf("%d -> ", t.Value)
		t = t.Next
	}
	fmt.Println()
}

func lookUpLinkedListNode(t *ListNode, v int) bool {
	if root == nil {
		t = &ListNode{v, nil}
		root = t
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookUpLinkedListNode(t.Next, v)
}

func size(t *ListNode) int {
	if t == nil {
		fmt.Println("-> Empty list!")
	}
	i := 0
	for t != nil {
		i++
		t = t.Next
	}
	return i
}

func main() {
	fmt.Println(root)
	root = nil
	root = nil
	traverseLinkedList(root)
	addNode(root, 1)
	addNode(root, -1)
	traverseLinkedList(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	traverseLinkedList(root)
	addNode(root, 100)
	traverseLinkedList(root)
	if lookUpLinkedListNode(root, 100) {
		fmt.Println("node exists!")
	}
	fmt.Println(size(root))
}
