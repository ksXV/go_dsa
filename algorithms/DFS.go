package algorithms

import "go_dsa/datastructures"

func walkDF(head *datastructures.BinaryNode, value int) bool {
	if head == nil {
		return false
	}
	if head.Value == value {
		return true
	}
	if head.Value < value {
		return walkDF(head.Right, value)
	}
	return walkDF(head.Left, value)
}

func DFS(head datastructures.BinaryNode, value int) bool {
	return walkDF(&head, value)
}
