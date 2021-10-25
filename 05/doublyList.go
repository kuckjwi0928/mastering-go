package main

import "fmt"

type DoublyListNode struct {
	Value    int
	Previous *DoublyListNode
	Next     *DoublyListNode
}

var doublyListRoot = new(DoublyListNode)

func addDoublyListNode(t *DoublyListNode, v int) int {
	if doublyListRoot == nil {
		t = &DoublyListNode{v, nil, nil}
		doublyListRoot = t
		return 0
	}

	if v == t.Value {
		fmt.Println("Node already exists: ", v)
		return -1
	}

	if t.Next == nil {
		temp := t
		t.Next = &DoublyListNode{v, temp, nil}
		return -2
	}
	return addDoublyListNode(t.Next, v)
}

func removeDoublyListNode(t *DoublyListNode, v int) bool {
	currNode := t

	if sizeDoublyList(t) == 1 {
		doublyListRoot = nil
		return true
	}

	if doublyListRoot.Value == v {
		temp := doublyListRoot.Next
		temp.Previous = nil
		doublyListRoot = temp
		return true
	}

	for currNode != nil {
		if currNode.Value == v {
			currNode.Previous.Next = currNode.Next
			if currNode.Next == nil {
				return true
			}
			currNode.Next.Previous = currNode.Previous
			return true
		}
		currNode = currNode.Next
	}
	return false
}

func traverseDoublyList(t *DoublyListNode) {
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

func reverseDoublyList(t *DoublyListNode) {
	if t == nil {
		fmt.Println("-> Empty list!")
		return
	}
	temp := t
	for t != nil {
		temp = t
		t = t.Next
	}
	for temp.Previous != nil {
		fmt.Printf("%d -> ", temp.Value)
		temp = temp.Previous
	}
	fmt.Printf("%d -> ", temp.Value)
	fmt.Println()
}

func sizeDoublyList(t *DoublyListNode) int {
	if t == nil {
		return 0
	}
	n := 0
	for t != nil {
		n++
		t = t.Next
	}
	return n
}

func lookupDoublyListNode(t *DoublyListNode, v int) bool {
	if doublyListRoot == nil {
		return false
	}
	if v == t.Value {
		return true
	}
	if t.Next == nil {
		return false
	}
	return lookupDoublyListNode(t.Next, v)
}

func main() {
	fmt.Println(doublyListRoot)
	doublyListRoot = nil
	traverseDoublyList(doublyListRoot)
	addDoublyListNode(doublyListRoot, 1)
	addDoublyListNode(doublyListRoot, 10)
	addDoublyListNode(doublyListRoot, 5)
	addDoublyListNode(doublyListRoot, 0)
	traverseDoublyList(doublyListRoot)
	removeDoublyListNode(doublyListRoot, 0)
	traverseDoublyList(doublyListRoot)
	removeDoublyListNode(doublyListRoot, 5)
	traverseDoublyList(doublyListRoot)
	removeDoublyListNode(doublyListRoot, 1)
	traverseDoublyList(doublyListRoot)
	removeDoublyListNode(doublyListRoot, 10)
	traverseDoublyList(doublyListRoot)
	addDoublyListNode(doublyListRoot, 5)
	fmt.Println("Size:", sizeDoublyList(doublyListRoot))
	reverseDoublyList(doublyListRoot)
}
