package hashtable

import "fmt"

// note that maps in go are hashtables

// size of the hashtable array
const ArraySize = 7

// HashTableStructure (Array)
type HashTable struct {
	array [ArraySize]*Bucket
}

// bucket structure (linkedlist)
type Bucket struct {
	head *bucketNode
}

// bucketNode structure
type bucketNode struct {
	key  string
	next *bucketNode
}

// for Hashtable
// Insert will take in a key and insert it into the hashtable
func (h *HashTable) Insert(key string) {
	index := Hash(key)
	h.array[index].Insert(key)
}

// Search
func (h *HashTable) Search(key string) bool {
	index := Hash(key)
	return h.array[index].Search(key)
}

// Delete
func (h *HashTable) Delete(key string) {
	index := Hash(key)
	h.array[index].Delete(key)
}

// for buckets (linked list)
// insert will take in a key, create a node with the key and insert the node in the bucket
func (b *Bucket) Insert(k string) {
	if !b.Search(k) {
		newNode := &bucketNode{key: k}
		newNode.next = b.head
		b.head = newNode
	} else {
		fmt.Println(k, "already exists")
	}
}

// search will take a key, create a node with the key and insert the node in the bucket
func (b *Bucket) Search(k string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == k {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

// delete will take a key and delete the node
func (b *Bucket) Delete(k string) {
	if b.head.key == k {
		b.head = b.head.next
		return
	}

	previousNode := b.head
	for previousNode.next != nil {
		if previousNode.next.key == k {
			// delete
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
}

// hash
func Hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// init will create a bucket in each slot of the hash table
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &Bucket{}
	}
	return result
}
