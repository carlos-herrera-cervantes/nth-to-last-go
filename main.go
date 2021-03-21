package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	length int
	head   *node
	tail   *node
}

// Returns the length of linked list
func (list linkedList) getLength() int {
	return list.length
}

// Prints the elements of linked list
func (list linkedList) printList() {
	for list.head != nil {
		fmt.Printf("%v -> ", list.head.data)
		list.head = list.head.next
	}

	fmt.Println()
}

// Adds new node in the linked list
func (list *linkedList) pushBack(node *node) {
	if list.head == nil {
		list.head = node
		list.tail = node
		list.length++

		return
	}

	list.tail.next = node
	list.tail = node
	list.length++
}

// Returns the kth node to last element.
// The runtime complexity is O(n) where n is the value of k param
func (list linkedList) nthToLast(k int) *node {
	p1 := list.head
	p2 := list.head

	for i := 0; i < k; i++ {
		if p1 == nil {
			return nil
		}

		p1 = p1.next
	}

	for p1 != nil {
		p1 = p1.next
		p2 = p2.next
	}

	return p2
}

func main() {
	node1 := &node{data: 10}
	node2 := &node{data: 3}
	node3 := &node{data: 11}
	node4 := &node{data: 100}
	linkedList1 := linkedList{}

	linkedList1.pushBack(node1)
	linkedList1.pushBack(node2)
	linkedList1.pushBack(node3)
	linkedList1.pushBack(node4)

	kthNode := linkedList1.nthToLast(6)

	if kthNode == nil {
		fmt.Println("There's no value")
	} else {
		fmt.Println("The kth element to last is: ", kthNode.data)
	}
}
