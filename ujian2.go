package main

import (
	"fmt"
)

type Node struct {
	Name string
	Priotiry int
	Next *Node
	Prev *Node

}

type Queue struct {
	Head *Node
	Tail *Node
}

func (q *Queue) Enqueue(nameValue string, priorityValue int) {
	newNode := &Node{Name: nameValue, Priotiry: priorityValue }

	if q.Head == nil { // ketika antreann kosong
		q.Head = newNode // orang paling depan
		q.Tail = newNode // orang paling belakang
		return
	}

	if priorityValue > q.Head.Priotiry { // ketika nilai prioritas pengantri baru > antran depan
		q.Head.Prev = newNode // berubah, sebelum antrean depan, sekrang antrain baru
		newNode.Next = q.Head
		q.Head = newNode  // node baru  = 5, node head = 4.  4 - 3 - 2 - 1 => 5 - 4 - 3 - 2 - 1
		return
	}

	// 5 -> <- 4 - 2 - 1
	// 5 -> 4   3.   2. prioritas = 3 . Ptr == 2
	//
	// 5
	// 4
	// 2
	// 4 harus hubung ke node baru (next)
	// 2 harus hubung ke node baru (prev)
	// node baru hubung ke 4 dan 2 (prev next)

	ptr := q.Head // tunjuk antrean ke - x

	for ptr != nil && ptr.Priotiry > priorityValue { // ketika antreann ditunjuk tidak kosong dan nilai prioritas nya besar dari pengantri baru
		ptr = ptr.Next
	}

	if ptr == nil { // ketika pointer menunjuk ke nil
		q.Tail.Next = newNode
		newNode.Prev = q.Tail
		q.Tail = newNode
		return
	} 

	if ptr.Prev == nil {
		q.Head.Prev = newNode 
		newNode.Next = q.Head
		q.Head = newNode
		return
	}
	

	// ketika newnode diselipkan / pointer tidak nil
	tempPrev := ptr.Prev

	ptr.Prev = newNode
	tempPrev.Next = newNode
	newNode.Prev = tempPrev
	newNode.Next = ptr
	
}


func (q *Queue) Dequeue() {
	if q.Head == nil {
		return
	}

	if q.Head == q.Tail {
		q.Head = nil
		q.Tail = nil
		return
	} else {
		tempNext := q.Head.Next
		q.Head.Next = nil
		tempNext.Prev = nil
		q.Head = tempNext
		return
	}
}

func (q *Queue) Display() {
	ptr := q.Head
	for ptr != nil {
		fmt.Println(ptr.Name)
		ptr = ptr.Next
	}
}

func main() {
	indomaretQueue := Queue{}

	indomaretQueue.Enqueue("Kevin", 1) 
	indomaretQueue.Enqueue("Haqqi", 5)

	indomaretQueue.Display() // Haqqi -> <- Kevin
}
