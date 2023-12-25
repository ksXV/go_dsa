package algorithms

func BinarySearch(item int, arr []int) int {
	left := 0
	right := len(arr) - 1
	for left <= right {
		middle := (left + right) / 2
		if arr[middle] == item {
			return middle
		}
		if arr[middle] > item {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return -1
}
