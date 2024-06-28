package main

import "fmt"

type linkedList struct {
	head *Node
}
type Node struct {
	data int
	next *Node
}

func (list *linkedList) insert(value int) {

	newNode := &Node{data: value, next: nil}

	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current.next != nil {
		current = current.next
	}
	current.next = newNode

	// linklst.head = newNode
}

func (list *linkedList) del(value int) {

	// prev := (*Node)(nil)

	// if &temp == nil {
	// 	return
	// }
	// if prev == nil {
	// 	*list.head = *temp.next
	// } else {
	// 	prev.next = temp.next
	// }

	if list.head == nil {
		return
	}
	if list.head.data == value {
		list.head = list.head.next
		return
	}
	prev := list.head
	for prev.next != nil && prev.next.data != value {
		prev = prev.next
	}
	if prev.next == nil {
		return
	}
	prev.next = prev.next.next

}

func (list *linkedList) display() {
	if list.head == nil {
		fmt.Println("nil")
		return
	}
	current := list.head
	for current.next != nil {
		fmt.Print(current.data, " -> ")
		current = current.next
	}

	fmt.Printf("%d->nil", current.data)
}

func main() {
	var list linkedList
	list.insert(10)
	list.insert(12)

	list.display()
	fmt.Println("\nDeleting value  2 \n")
	list.del(10)
	list.del(12)
	fmt.Println("After deletion \n")
	list.display()
}
