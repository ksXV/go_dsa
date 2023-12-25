package datastructures

type Node[T any] struct {
	Next  *Node[T]
	Prev  *Node[T]
	Value T
}

type BinaryNode struct {
	Value int
	Left  *BinaryNode
	Right *BinaryNode
}
