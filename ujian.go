package main

import (
	"fmt"
)

// soal 1

type CircularQueue struct {
	Size int
	Head int
	Tail int
	Data[] int
}

func NewCircularQueue(size int) CircularQueue {
	return CircularQueue{
		Size: size,
		Head: -1,
		Tail: -1,
		Data: make([]int, size),
	}
}

func (cq CircularQueue) IsEmpty() bool {
	return cq.Head == -1
}

func (cq CircularQueue) IsFull() bool {
	nextTailPosition := (cq.Tail + 1) % cq.Size
	return nextTailPosition == cq.Head
}

func (cq *CircularQueue) Enqueue(value int) {
	if cq.IsFull() {
		fmt.Println("Antrean Penuh")
		return 
	}

	if cq.IsEmpty() {
		cq.Head = 0
		cq.Tail = 0
	} else {
		cq.Tail = (cq.Tail + 1) % cq.Size
	}

	cq.Data[cq.Tail] = value
}

func (cq *CircularQueue) Dequeue() int {
	if cq.IsEmpty() {
		return -1
	}

	outData := cq.Data[cq.Head]
	if cq.Head == cq.Tail {
		cq.Tail = -1
		cq.Head = -1
	} else {
		cq.Head = (cq.Head + 1) % cq.Size
	}
	return outData

}

func main() {
	antreanLogin := NewCircularQueue(10)

	antreanLogin.Enqueue(999)
	antreanLogin.Enqueue(100)
	antreanLogin.Enqueue(200)
	antreanLogin.Enqueue(300)
	antreanLogin.Enqueue(400)
	antreanLogin.Enqueue(500)
	antreanLogin.Enqueue(600)
	antreanLogin.Enqueue(700)
	antreanLogin.Enqueue(800)
	antreanLogin.Enqueue(900)

	
    fmt.Println(antreanLogin.Dequeue())
	fmt.Println(antreanLogin.Dequeue())
	fmt.Println(antreanLogin.Dequeue())
	fmt.Println(antreanLogin.Dequeue())
	fmt.Println(antreanLogin.Dequeue())

	fmt.Println(antreanLogin.Data)
}


