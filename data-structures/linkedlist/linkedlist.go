package linkedlist

import (
	//"fmt"
	"strconv"
)

type LinkedList struct {
	head *Node
	length int
}

type Node struct {
	data int
	next *Node
}

func InitLinkedList() *LinkedList {
	return &LinkedList{nil, 0}
}

// inserts node at the end of the list
func (ll *LinkedList) Insert(data int) {
	node := createNode(data)
	if ll.length == 0 {
		ll.head = node
	} else if ll.length == 1 {
		ll.head.next = node
	} else {
		temp := ll.head
		for temp.next != nil {
			temp = temp.next
		}
		temp.next = node
	}
	ll.length++
}

func (ll *LinkedList) Update(data int) {
}

func (ll *LinkedList) Delete(data int) {
}

func createNode(data int) *Node {
	return &Node{data, nil}
}

func (ll *LinkedList) GetHead() int {
	return ll.head.data
}

func (ll *LinkedList) PrintList() string {
	if ll.length == 0 {
		return "List is empty"
	} else if ll.length == 1 {
		return string(ll.head.data)
	} else {
		var list string
		temp := ll.head
		for temp.next != nil {
			list += strconv.Itoa(temp.data) + ", "
			temp = temp.next
		}
		list += strconv.Itoa(temp.data)
		return list
	}
}
