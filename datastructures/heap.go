package datastructures

type Heap struct {
	Length int
	data   []int
}

func (h *Heap) Insert(val int) {
	h.data[h.Length] = val
	h.heapifyUp(h.Length)
	h.Length = h.Length + 1
}
func (h *Heap) Delete() int {
	if h.Length == 0 {
		return -1
	}

	out := h.data[0]
	h.Length--
	if h.Length == 0 {
		h.data = []int{}
		h.Length--
		return out
	}
	h.data[0] = h.data[h.Length]
	h.heapifyDown(0)
	return out
}
func parent(i int) int {
	return (i - 1) / 2
}
func leftChild(i int) int {
	return i*2 + 1
}
func rightChild(i int) int {
	return i*2 + 2
}
func (h *Heap) heapifyUp(idx int) {
	if idx == 0 {
		return
	}
	p := parent(idx)
	val := h.data[idx]
	parentValue := h.data[p]
	if parentValue > val {
		h.data[idx] = parentValue
		h.data[p] = val
		h.heapifyUp(p)
	}
}

func (h *Heap) heapifyDown(idx int) {

	lIdx := leftChild(idx)
	rIdx := rightChild(idx)

	if lIdx >= h.Length || idx >= h.Length {
		return
	}
	lV := h.data[lIdx]
	rV := h.data[rIdx]
	v := h.data[idx]

	if lV > rV && v > rV {
		h.data[idx] = rV
		h.data[rIdx] = v
		h.heapifyDown(rIdx)
	} else if lV < rV && v > lV {
		h.data[idx] = lV
		h.data[lIdx] = v
		h.heapifyDown(lIdx)
	}

}
