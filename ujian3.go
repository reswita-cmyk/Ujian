package main

import "fmt"

type Node struct {
	Name string
	Next *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (ll *LinkedList) add(nameValue string) {
	newNode := &Node{Name: nameValue }

	if ll.Head == nil {
		ll.Head = newNode
		ll.Tail = newNode
		return
	}

	ll.Tail.Next = newNode
	ll.Tail = newNode
	
}

func (ll *LinkedList) delete() {
	if ll.Tail == nil {
		return
	}

	if ll.Tail == ll.Head {
		ll.Head = nil
		ll.Tail = nil
		return
	}

	ptr := ll.Head
	for ptr.Next != ll.Tail { // 5 - 4 - 3 - 2 - 1
		ptr = ptr.Next
	}
	ptr.Next = nil
	ll.Tail = ptr
}

func (ll *LinkedList) Display() {
	ptr := ll.Head
	for ptr != nil {
		fmt.Println(ptr.Name)
		ptr = ptr.Next
	}
}

func main() {
	arenaPertempuran := LinkedList{}

	arenaPertempuran.add("Yusuf")
	arenaPertempuran.add("Kevin")
	arenaPertempuran.add("Mario")
	arenaPertempuran.add("Haqqi")
	arenaPertempuran.add("Anicode")

	arenaPertempuran.delete()

	arenaPertempuran.Display()
}




