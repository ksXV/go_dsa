package algorithms

func hasUnvisited(seen []bool, dists []int) bool {
	has := false

	for i := 0; i < len(seen); i++ {
		if !seen[i] && dists[i] < 1<<62 {
			has = true
			break
		}
	}
	return has
}

func getLowestUnvisited(seen []bool, dists []int) int {
	idx := -1
	lowestDistance := 1 << 62
	for i := 0; i < len(seen); i++ {
		if seen[i] {
			continue
		}
		if lowestDistance > dists[i] {
			lowestDistance = dists[i]
			idx = i
		}
	}
	return idx

}

func Dijkstra_list(source, sink int, arr WeightedAdjacencyList) []int {
	seen := make([]bool, len(arr))
	prev := make([]int, len(arr))
	for i := 0; i < len(prev); i++ {
		prev[i] = -1
	}
	seen[source] = true
	dists := make([]int, len(arr))

	for i := 0; i < len(dists); i++ {
		dists[i] = 1 << 62
	}
	dists[source] = 0

	for hasUnvisited(seen, dists) {
		curr := getLowestUnvisited(seen, dists)

		seen[curr] = true

		adjs := arr[curr]

		for i := 0; i < len(adjs); i++ {
			edge := adjs[i]
			if seen[edge.To] {
				continue
			}

			dist := dists[curr] + edge.Weight
			if dist < dists[edge.To] {
				dists[edge.To] = dist
				prev[edge.To] = curr
			}

		}

	}

	out := []int{}
	curr := sink
	for prev[curr] != -1 {
		out = append(out, curr)
		curr = prev[curr]
	}

	out = append(out, source)
	reversedOut := make([]int, len(out))
	for i := 0; i < len(reversedOut); i++ {
		reversedOut[i] = out[len(out)-i-1]
	}
	return reversedOut
}
