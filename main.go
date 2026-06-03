package main

import "fmt"

type Node struct {
	Data int
	Next *Node
}

// Insert at Head
func insertHead(head *Node, data int) *Node {
	newNode := &Node{Data: data}

	newNode.Next = head

	return newNode
}

// Insert at Tail
func insertTail(head *Node, data int) *Node {
	newNode := &Node{Data: data}

	if head == nil {
		return newNode
	}

	temp := head

	for temp.Next != nil {
		temp = temp.Next
	}

	temp.Next = newNode

	return head
}

// Insert After a Given Value
func insertAfter(head *Node, target int, data int) *Node {
	temp := head

	for temp != nil {
		if temp.Data == target {
			newNode := &Node{Data: data}

			newNode.Next = temp.Next
			temp.Next = newNode

			return head
		}

		temp = temp.Next
	}

	return head
}

// Delete by Value
func deleteValue(head *Node, value int) *Node {

	if head == nil {
		return nil
	}

	// Delete head
	if head.Data == value {
		return head.Next
	}

	temp := head

	for temp.Next != nil && temp.Next.Data != value {
		temp = temp.Next
	}

	if temp.Next != nil {
		temp.Next = temp.Next.Next
	}

	return head
}

// Print List
func printList(head *Node) {
	temp := head

	for temp != nil {
		fmt.Print(temp.Data, " -> ")
		temp = temp.Next
	}

	fmt.Println("nil")
}

func main() {

	var head *Node

	// Insert at Head
	head = insertHead(head, 30)
	head = insertHead(head, 20)
	head = insertHead(head, 10)

	fmt.Println("After Insert Head:")
	printList(head)

	// Insert at Tail
	head = insertTail(head, 40)
	head = insertTail(head, 50)

	fmt.Println("After Insert Tail:")
	printList(head)

	// Insert in Middle
	head = insertAfter(head, 20, 25)

	fmt.Println("After Insert 25 after 20:")
	printList(head)

	// Delete Value
	head = deleteValue(head, 30)

	fmt.Println("After Delete 30:")
	printList(head)
}