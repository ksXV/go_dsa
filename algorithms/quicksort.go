package algorithms

func partition(arr []int, lo, hi int) int {
	pivot := arr[hi]
	idx := lo - 1
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			idx++
			temp := arr[i]
			arr[i] = arr[idx]
			arr[idx] = temp
		}
	}
	idx++
	arr[hi] = arr[idx]
	arr[idx] = pivot

	return idx
}

func qs(arr []int, lo, hi int) {
	if lo >= hi {
		return
	}
	pivotIdx := partition(arr, lo, hi)

	qs(arr, lo, pivotIdx-1)
	qs(arr, pivotIdx+1, hi)

}

func QuickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
