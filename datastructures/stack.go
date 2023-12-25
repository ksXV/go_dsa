package datastructures

type Stack struct {
	head   *Node
	Length int
}

func (s *Stack) Push(value int) {
	node := &Node{Value: value}
	if s.head == nil {
		s.head = node
		s.Length = 1
		return
	}
	s.Length++
	node.Next = s.head
	s.head = node
}
func (s *Stack) Pop() *int {
	if s.head == nil {
		return nil
	}
	s.Length--
	temp := s.head
	s.head = s.head.Next
	temp.Next = nil
	return &temp.Value

}
func (s *Stack) Peek() *int {
	if s.head == nil {
		return nil
	}
	return &s.head.Value
}
