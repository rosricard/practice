package heap

// A heap is a complete binary tree

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	slice []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.slice = append(h.slice, key)
	h.maxHeapifyUp(len(h.slice) - 1)
}

// Extract returns the largest key., and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.slice[0]
	l := len(h.slice) - 1
	// check if array is empty
	if len(h.slice) == 0 {
		fmt.Println("cannot extract because array length is 0")
		return -1
	}
	// take out the last index and put it in the root
	h.slice[0] = h.slice[l]
	h.slice = h.slice[:]

	return extracted
}

// maxHeapifyUp will heapify from bottom top
func (h *MaxHeap) maxHeapifyUp(index int) {
	for h.slice[parent(index)] < h.slice[index] {
		h.swap(parent(index), index)
		index = parent(index)
	}
}

// maxHeapifyDown will heapify top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.slice) - 1
	l, r := left(index), right(index)
	childToCompare := 0
	// loop while index has at least one child
	for l <= lastIndex {
		if l == lastIndex { // when left child is the only child
			childToCompare = 1
		} else if h.slice[l] > h.slice[r] { // when left child is larger
			childToCompare = 1
		} else { // when right child is larger
			childToCompare = r
		}
		// compare array value of current index to larger child and swap if smaller
		if h.slice[index] < h.slice[childToCompare] {
			h.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}
}

// get the parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get the left child index
func left(i int) int {
	return 2*i + 1
}

// get the right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i1, i2 int) {
	h.slice[i1], h.slice[i2] = h.slice[i2], h.slice[i1]

}
