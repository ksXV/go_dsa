package algorithms

func LinearSearch(item int, arr []int) int {
	for i := 0; i < len(arr); i++ {
		if arr[i] == item {
			return i
		}
	}
	return -1
}
