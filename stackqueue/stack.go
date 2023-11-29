package stackqueue

// Stack represents the stack that holds a slice for the LIFO
type Stack struct {
	items []int
}

// Push
// add new value to end of queue
func (s *Stack) Push(i int) {
	s.items = append(s.items, i)
}

// Pop
// Remove first value in queue, shift all values over and return the removed value
func (s *Stack) Pop() int {
	l := len(s.items) - 1
	end := s.items[l]
	s.items = s.items[:l]
	return end
}
