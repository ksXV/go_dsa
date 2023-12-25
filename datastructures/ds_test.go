package datastructures_test

import (
	"go_dsa/datastructures"
	"testing"
)

func TestQueue(t *testing.T) {
	queue := datastructures.Queue{}
	queue.Enque(10)
	queue.Enque(20)
	queue.Enque(30)
	wantedDeque := 10
	dequeuedItem := queue.Deque()
	if wantedDeque != *dequeuedItem {
		t.Fatalf("Deque returned %v instead of %v", *dequeuedItem, wantedDeque)
	}
	wantedPeek := 20
	peekedItem := queue.Peek()
	if wantedPeek != *peekedItem {
		t.Fatalf("Peek returned %v instead of %v", *peekedItem, wantedPeek)
	}
}
func TestStack(t *testing.T) {
	stack := datastructures.Stack{}
	stack.Push(30)
	stack.Push(20)
	stack.Push(10)
	wantedPop := 10
	popedItem := stack.Pop()
	if wantedPop != *popedItem {
		t.Fatalf("Pop returned %v instead of %v", *popedItem, wantedPop)
	}
	wantedPeek := 20
	peekedItem := stack.Peek()
	if wantedPeek != *peekedItem {
		t.Fatalf("Peek returned %v instead of %v", *peekedItem, wantedPeek)
	}
}
