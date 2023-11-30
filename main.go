package main

import (
	"fmt"

	hashtable "github.com/rosricard/practice/maps"
)

func main() {
	// myQueue := stackqueue.Queue{}

	// myQueue.Enqueue(100)
	// myQueue.Enqueue(200)
	// myQueue.Enqueue(300)
	// myQueue.Dequeue()
	// fmt.Print(myQueue)

	// HEAPS
	// m := &heap.MaxHeap{}
	// fmt.Println(m)
	// buildHeap := []int{10, 20, 30, 5, 7, 9, 11, 13, 15, 17}
	// for _, v := range buildHeap {
	// 	m.Insert(v)
	// 	fmt.Println(m)
	// }

	// // always extracting and the largest key is always in the front
	// for i := 0; i < 5; i++ {
	// 	m.Extract()
	// 	fmt.Println(m)
	// }

	// HASHTABLE
	testHashTable := hashtable.Init()
	list := []string{
		"ERIC",
		"KENNY",
		"KYLE",
		"STAN",
		"RANDY",
		"BUTTERS",
		"TOKEN",
	}
	for _, v := range list {
		testHashTable.Insert(v)
	}
	testHashTable.Delete("STAN")
	fmt.Println(testHashTable.Search("STAN"))
	fmt.Println(testHashTable.Search("KENNY"))

}
