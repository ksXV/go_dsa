package datastructures

type Queue[T any] struct {
	head   *Node[T]
	tail   *Node[T]
	Length int
}
type QueueOperations interface {
	Enque(value int)
	Deque() *int
	Peek() *int
}

func (q *Queue[T]) Enque(value T) {
	node := &Node[T]{Value: value}
	if q.head == nil {
		q.head = node
		q.tail = node
		q.Length = 1
		return
	}
	q.tail.Next = node
	q.tail = node
	q.Length++
}
func (q *Queue[T]) Deque() *T {
	if q.head == nil {
		return nil
	}
	temp := q.head
	q.head = q.head.Next
	temp.Next = nil
	q.Length--
	if q.Length == 0 {
		q.tail = nil
	}
	return &temp.Value

}
func (q *Queue[T]) Peek() *T {
	if q.head == nil {
		return nil
	}
	return &q.head.Value
}
