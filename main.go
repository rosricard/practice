package main

import (
	"fmt"

	"github.com/rosricard/practice/interfaces"
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
	// testHashTable := hashtable.Init()
	// list := []string{
	// 	"ERIC",
	// 	"KENNY",
	// 	"KYLE",
	// 	"STAN",
	// 	"RANDY",
	// 	"BUTTERS",
	// 	"TOKEN",
	// }
	// for _, v := range list {
	// 	testHashTable.Insert(v)
	// }
	// testHashTable.Delete("STAN")
	// fmt.Println(testHashTable.Search("STAN"))
	// fmt.Println(testHashTable.Search("KENNY"))

	/// LINKED LIST (SINGLY LINKED)
	// mylist := linkedlist.Linkedlist{}
	// node1 := &linkedlist.Node{Data: 48}
	// node2 := &linkedlist.Node{Data: 18}
	// node3 := &linkedlist.Node{Data: 16}
	// node4 := &linkedlist.Node{Data: 11}
	// node5 := &linkedlist.Node{Data: 7}
	// node6 := &linkedlist.Node{Data: 2}
	// mylist.Prepend(node1)
	// mylist.Prepend(node2)
	// mylist.Prepend(node3)
	// mylist.Prepend(node4)
	// mylist.Prepend(node5)
	// mylist.Prepend(node6)
	// mylist.PrintListData()
	// mylist.DeleteWithValue(100)
	// mylist.DeleteWithValue(2)
	// mylist.PrintListData()
	// emptyList := linkedlist.Linkedlist{}
	// emptyList.DeleteWithValue(10)

	// INTERFACES
	c1 := interfaces.Circle{4.5}
	r1 := interfaces.Rect{5, 7}
	shapes := []interfaces.Shape{&c1, &r1}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
