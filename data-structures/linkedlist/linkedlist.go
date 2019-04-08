package linkedlist

type LinkedList struct {
	head *Node
	tail *Node
}

type Node struct {
	data int
	next *Node
}

func InitLinkedList() *LinkedList {
	var list *LinkedList
	return list
}


