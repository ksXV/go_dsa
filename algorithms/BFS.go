package algorithms

import "go_dsa/datastructures"

func BFS(head datastructures.BinaryNode, needle int) bool {
	queue := datastructures.Queue[datastructures.BinaryNode]{}
	queue.Enque(head)

	for queue.Length > 0 {
		next := queue.Deque()
		if needle == next.Value {
			return true
		}
		if next.Left != nil {
			queue.Enque(*next.Left)
		}
		if next.Right != nil {
			queue.Enque(*next.Right)
		}
	}
	return false
}
