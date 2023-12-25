package algorithms

import "go_dsa/datastructures"

type WeightedAdjacencyMatrix = [][]int

func BFSOnWeightedAdjMatrix(graph WeightedAdjacencyMatrix, source, needle int) []int {
	graphLength := len(graph)
	seen := make([]bool, graphLength)
	prev := make([]int, graphLength)
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}
	seen[source] = true

	queue := datastructures.Queue[int]{}
	queue.Enque(source)

	for {
		curr := queue.Deque()
		if *curr == needle {
			break
		}
		adjs := graph[*curr]
		for i := 0; i < len(adjs); i++ {
			if adjs[i] == 0 {
				continue
			}
			if seen[i] {
				continue
			}

			seen[i] = true
			prev[i] = *curr
			queue.Enque(i)
		}

		if queue.Length == 0 {
			break
		}
	}
	// build it backwards
	curr := needle
	out := []int{}

	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	reverseOut := []int{}
	for i := range out {
		reverseOut = append(reverseOut, out[len(out)-i-1])
	}
	return reverseOut
}
