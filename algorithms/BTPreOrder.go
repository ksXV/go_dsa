package algorithms

import "go_dsa/datastructures"

func walkBT(curr *datastructures.BinaryNode, path *[]int) []int {
	if curr == nil {
		return *path
	}
	//recurse
	//pre
	*path = append(*path, curr.Value)
	//recurse
	walkBT(curr.Left, path)
	walkBT(curr.Right, path)
	//post
	return *path
}

func PreOrderSearch(head *datastructures.BinaryNode) []int {
	path := []int{}

	return walkBT(head, &path)
}
