package main

import "fmt"

type MyArray []int
type Node[T any] struct {
	Data T
	Next *Node[T]
}
type linkedList[T any] struct {
	Head   *Node[T]
	Equals func(a, b T) bool
}

func (list *linkedList[T]) Append(Data T) {
	newNode := &Node[T]{Data: Data, Next: nil}
	if list.Head == nil {
		list.Head = newNode
		return
	}
	current := list.Head
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
}
func (list *linkedList[T]) del(value T) {

	if list.Head == nil {
		return
	}
	if list.Equals(list.Head.Data, value) {
		list.Head = list.Head.Next
		return
	}
	prev := list.Head
	for prev.Next != nil && !list.Equals(prev.Next.Data, value) {
		prev = prev.Next
	}
	if prev.Next == nil {
		return
	}
	prev.Next = prev.Next.Next

}
func (a MyArray) Len() int {
	return len(a)
}

func (list *linkedList[T]) display() {
	if list.Head == nil {
		fmt.Println("nil")
		return
	}
	current := list.Head
	for current.Next != nil {
		fmt.Print(current.Data, " -> ")
		current = current.Next
	}

	fmt.Println(current.Data, "->nil")
}

func main() {
	intEquals := func(a, b int) bool { return a == b }
	strEquals := func(a, b string) bool { return a == b }
	list2 := linkedList[string]{Equals: strEquals}
	list := linkedList[int]{Equals: intEquals}
	list.Append(10)
	list.Append(12)

	list.display()
	fmt.Println("\nDeleting value  2 \n")
	list.del(10)
	list.del(12)
	fmt.Println("After deletion \n")
	list.display()

	list2.Append("Hello")
	list2.Append("World")

	list2.display()
	fmt.Println("\nDeleting world \n")

	list2.del("World")
	fmt.Println("After deletion \n")

	list2.display()

	arrEquals := func(a, b MyArray) bool { return a.Len() == b.Len() }
	list3 := linkedList[MyArray]{Equals: arrEquals}

	list3.Append([]int{5, 2, 4, 1, 9})

	list3.display()

	list3.Append([]int{3, 2, 4, 78})

	list3.display()
}
