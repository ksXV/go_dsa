package algorithms

type WeightedAdjacencyList = [][]GraphEdge

type GraphEdge struct {
	To     int
	Weight int
}

func walkDFSOnGraph(graph WeightedAdjacencyList, curr, needle int, seen []bool, path *[]int) bool {

	if seen[curr] {
		return false
	}

	seen[curr] = true

	//pre
	*path = append(*path, curr)
	if curr == needle {
		*path = append(*path, curr)
		return true
	}

	//recurse
	list := graph[curr]

	for i := 0; i < len(list); i++ {
		edge := list[i]
		if walkDFSOnGraph(graph, edge.To, needle, seen, path) {
			return true
		}
	}

	//post
	*path = (*path)[:len(*path)-1]
	return false
}

func DFSonGraph(graph WeightedAdjacencyList, source, needle int) []int {
	seen := make([]bool, len(graph))
	path := []int{}

	walkDFSOnGraph(graph, source, needle, seen, &path)
	return path

}
