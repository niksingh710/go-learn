package main

import "fmt"

type LinkedList interface {
	getHead() *Node
}

type Node struct {
	val  int
	next *Node
	prev *Node
}

type List struct {
	head *Node
}

type DList struct {
	head *Node
	tail *Node
}

func print(l LinkedList) {
	// curr := (*l).head
	curr := l.getHead()
	arrow := "->"
	if curr.next != nil && curr == curr.next.prev {
		arrow = "<->"
	}
	for curr.next != nil {
		fmt.Printf("%v %v ", curr.val, arrow)
		curr = curr.next
	}
	fmt.Println(curr.val)
}

// NOTE: Go derefrences variable with '.' operator so (*l).head == l.head both will have the same effect

// SINGLE Linked List
func (l *List) add(val int) {
	node := Node{
		val:  val,
		next: nil,
	}
	if l.head == nil {
		// (*l).head = &node
		l.head = &node
		return
	}
	// curr := (*l).head
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &node
}

func (l List) getHead() *Node {
	return l.head
}

// DOUBLE LINKED LIST
func (d *DList) add(val int) {
	node := Node{
		val:  val,
		next: nil,
		prev: nil,
	}
	if d.head == nil {
		d.head = &node
		d.tail = d.head
	}
  node.prev = d.tail
	d.tail.next = &node
	d.tail = d.tail.next
}

func (dl DList) getHead() *Node {
	return dl.head
}
