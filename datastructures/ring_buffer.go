package datastructures

import (
	"errors"
)

type RingBuffer struct {
	Head    int
	Tail    int
	Length  int
	Content []int
}

// needs work
func (r *RingBuffer) New(head, tail, length int) {
	r.Head = head % length
	r.Tail = tail % length
	r.Length = length
	r.Content = make([]int, length)
}
func (r *RingBuffer) InsertAt(position, value int) error {
	if position < 0 {
		return errors.New("Invalid position")
	}
	r.Content[r.Tail%r.Length] = value
	return nil
}
