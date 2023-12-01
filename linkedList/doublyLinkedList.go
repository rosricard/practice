package linkedlist

import (
	"fmt"
)

// PlateNode represents a node in the doubly linked list
type PlateNode struct {
	plateType string
	prev      *PlateNode
	next      *PlateNode
}

// PlateLinkedList represents the doubly linked list of plates
type PlateLinkedList struct {
	head       *PlateNode
	tail       *PlateNode
	lastOfType map[string]*PlateNode
}

// NewPlateLinkedList creates a new empty PlateLinkedList
func NewPlateLinkedList() *PlateLinkedList {
	return &PlateLinkedList{
		lastOfType: make(map[string]*PlateNode),
	}
}

// InsertPlates inserts one or two plates into the linked list
func (ll *PlateLinkedList) InsertPlates(plates ...string) {
	for _, plate := range plates {
		newNode := &PlateNode{plateType: plate}
		lastNode, exists := ll.lastOfType[plate]

		if !exists { // Insert at the end if this plate type is new
			if ll.tail != nil {
				ll.tail.next = newNode
				newNode.prev = ll.tail
				ll.tail = newNode
			} else { // List is empty
				ll.head = newNode
				ll.tail = newNode
			}
		} else { // Insert after the last node of the same type
			newNode.next = lastNode.next
			newNode.prev = lastNode
			if lastNode.next != nil {
				lastNode.next.prev = newNode
			} else {
				ll.tail = newNode
			}
			lastNode.next = newNode
		}

		ll.lastOfType[plate] = newNode
	}
}

// PrintList prints the plates in the linked list
func (ll *PlateLinkedList) PrintList() {
	for node := ll.head; node != nil; node = node.next {
		fmt.Printf("%s -> ", node.plateType)
	}
	fmt.Println("End")
}

func main() {
	ll := NewPlateLinkedList()
	ll.InsertPlates("flower-decorated")
	ll.InsertPlates("light green", "light green")
	ll.InsertPlates("big blue", "big blue")
	ll.PrintList() // Prints the initial setup

	ll.InsertPlates("light green", "light green")
	ll.PrintList() // Prints after inserting two light green plates

	ll.InsertPlates("big blue", "light green")
	ll.PrintList() // Prints after inserting a big blue and a light green plate
}
