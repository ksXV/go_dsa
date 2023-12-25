package datastructures

import "errors"

type LinkedList struct {
	Head   *Node
	Tail   *Node
	Length int
}

func (l LinkedList) checkForBoundaries(idx int) {
	if idx > l.GetLength() {
		panic("Index is bigger that the linked list length")
	}
	return
}

func (l LinkedList) GetLength() int {
	return l.Length
}
func (l LinkedList) Get(idx int) int {
	curr := l.Head
	l.checkForBoundaries(idx)

	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	return curr.Value

}
func (l *LinkedList) Append(item int) {
	l.Length++
	node := &Node{Value: item}

	if l.Tail == nil {
		l.Head = node
		l.Tail = node
		return
	}
	l.Tail.Next = node
	node.Prev = l.Tail
	l.Tail = node

}
func (l *LinkedList) InsertAt(item, idx int) {
	curr := l.Head
	l.checkForBoundaries(idx)
	if l.GetLength() == idx {
		l.Append(item)
	} else if idx == 0 {
		l.Prepend(item)
	}
	l.Length++
	node := &Node{Value: item}

	for i := 0; i < idx; i++ {
		curr = curr.Next
	}
	node.Next = curr
	node.Prev = curr.Prev
	curr.Prev = node
	curr.Prev.Next = curr

}
func (l *LinkedList) Remove(item int) (int, error) {
	curr := l.Head
	for i := 0; i < l.Length; i++ {
		if curr.Value == item {
			break
		}
		curr = curr.Next
	}
	if curr == nil {
		return -1, errors.New("couldn't find the value")
	}
	l.Length--
	if l.Length == 0 {
		value := l.Head.Value
		l.Head = nil
		l.Tail = nil
		return value, nil
	}
	if curr.Prev != nil {
		curr.Prev.Next = curr.Next
	}
	if curr.Next != nil {
		curr.Next.Prev = curr.Prev
	}
	if curr == l.Head {
		l.Head = curr.Next
	}
	if curr == l.Tail {
		l.Head = curr.Prev
	}
	curr.Prev = nil
	curr.Next = nil
	return curr.Value, nil

}
func (l *LinkedList) Prepend(item int) {
	node := &Node{Value: item}
	l.Length++

	if l.Head == nil {
		l.Head = node
		l.Tail = node
		return
	}
	node.Next = l.Head
	l.Head.Prev = node
	l.Head = node
}
