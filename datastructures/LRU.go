package datastructures

type LRU[K comparable, V any] struct {
	length        int
	head          *Node[V]
	tail          *Node[V]
	lookup        map[K]*Node[V]
	reverseLookup map[*Node[V]]K
	Capacity      int
}

func (l LRU[K, V]) createNode(v V) *Node[V] {
	return &Node[V]{Value: v}
}

func (l *LRU[K, V]) Update(key K, value V) {
	node, doesItExist := l.lookup[key]
	if !doesItExist {
		node = l.createNode(value)
		l.length++
		l.prepend(node)
		l.trimCache()

		l.lookup[key] = node
		l.reverseLookup[node] = key

	} else {
		l.detach(node)
		l.prepend(node)
		node.Value = value
	}

}

func (l *LRU[K, V]) Get(key K) *V {
	node, doesItExist := l.lookup[key]
	if !doesItExist {
		return nil
	}
	l.detach(node)
	l.prepend(node)

	return &node.Value

}
func (l *LRU[K, V]) detach(node *Node[V]) {
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	}

	if l.head == node {
		l.head = l.head.Next
	}

	if l.tail == node {
		l.tail = l.tail.Prev
	}
	node.Next = nil
	node.Prev = nil

}
func (l *LRU[K, V]) prepend(node *Node[V]) {
	if l.head == nil {
		l.head = node
		l.tail = node
		return
	}

	node.Next = l.head
	l.head.Prev = node
	l.head = node

}
func (l *LRU[K, V]) trimCache() {
	if l.length <= l.Capacity {
		return
	}

	tail := l.tail
	l.detach(l.tail)

	key := l.reverseLookup[tail]
	delete(l.lookup, key)
	delete(l.reverseLookup, tail)
	l.length--
}
