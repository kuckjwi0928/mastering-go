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

	if root.Value > v {
		temp := root
		root = &ListNode{v, temp}
		return -2
	}

	prevNode := t
	currNode := t

	for currNode != nil {
		if v == currNode.Value {
			fmt.Println("Node already exist:", v)
			return -1
		}
		if currNode.Value > v {
			temp := currNode
			prevNode.Next = &ListNode{v, temp}
			return -2
		}
		prevNode = currNode
		currNode = currNode.Next
	}

	prevNode.Next = &ListNode{v, nil}

	return -2
}

func removeNode(t *ListNode, v int) (bool, int) {
	prevNode := t
	currNode := t
	for currNode != nil {
		if currNode.Value == v {
			prevNode.Next = currNode.Next
			return true, currNode.Value
		}
		prevNode = currNode
		currNode = currNode.Next
	}
	return false, -1
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
	addNode(root, 1)
	addNode(root, 2)
	addNode(root, 3)
	addNode(root, 4)
	addNode(root, 5)
	addNode(root, 0)
	addNode(root, -1)
	traverseLinkedList(root)
	addNode(root, 10)
	addNode(root, 5)
	addNode(root, 45)
	addNode(root, 5)
	addNode(root, 5)
	removed, value := removeNode(root, 5)
	if removed {
		fmt.Println("Node Removed", value)
	}
	traverseLinkedList(root)
	addNode(root, 100)
	traverseLinkedList(root)
	if lookUpLinkedListNode(root, 100) {
		fmt.Println("node exists!")
	}
	fmt.Println(size(root))
}
