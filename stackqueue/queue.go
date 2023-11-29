package stackqueue

// Queue represents the slice of values that will hold the data for FIFO
type Queue struct {
	items []int
}

// Enqueue
// add new value to end of queue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue
// Remove last value in queue, shift all values over and return the removed value
func (q *Queue) Dequeue() int {
	first := q.items[0]
	q.items = q.items[1:len(q.items)]
	return first
}
