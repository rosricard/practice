package main

import (
	"fmt"

	"github.com/rosricard/practice/stackqueue"
)

func main() {
	myQueue := stackqueue.Queue{}

	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Dequeue()
	fmt.Print(myQueue)
}
