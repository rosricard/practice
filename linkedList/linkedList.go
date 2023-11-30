package linkedlist

import "fmt"

type Node struct {
	Data int
	next *Node
}

type Linkedlist struct {
	head   *Node
	length int
}

func (l *Linkedlist) Prepend(n *Node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l Linkedlist) PrintListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.Data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *Linkedlist) DeleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.Data == value {
		l.head = l.head.next
		l.length--
		return
	}

	previousToDelete := l.head
	for previousToDelete.next.Data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}
