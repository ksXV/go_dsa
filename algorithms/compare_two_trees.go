package algorithms

import "go_dsa/datastructures"

func Compare(a *datastructures.BinaryNode, b *datastructures.BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}
	if a.Value != b.Value {
		return false
	}
	return Compare(a.Left, b.Left) && Compare(a.Right, b.Right)
}
