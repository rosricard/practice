package bst

type BST struct {
	Value int

	Left  *BST
	Right *BST
}

// implement itterative searching because its more efficient than recursive
// Average: O(log(n)) time | O(1) space
// Worst: O(n) time | O(1) space

func (tree *BST) Insert(value int) *BST {
	newNode := &BST{Value: value}
	current := tree

	parentEdgeLink := &tree // pointer to the parent node
	for {
		if value < current.Value {
			parentEdgeLink = &current.Left
			current = current.Left
		} else {
			parentEdgeLink = &current.Right
			current = current.Right
		}

		if current == nil {
			*parentEdgeLink = newNode
			break
		}
		return tree
	}
	return tree
}

func (tree *BST) Contains(value int) bool {
	current := tree

	for current != nil {
		if value < current.Value {
			current = current.Left
		} else if value > current.Value {
			current = current.Right
		} else {
			return true
		}
	}
	return false
}

func (tree *BST) Remove(value int) (t *BST) {
	// find the node to remove
	// replace with the number in the right subtree that is the smallest
	// if there is no right subtree, replace with the left subtree that has the largest value

	current := tree
	parentEdgeLink := &tree
	var nodeToRemove *BST
	for {
		if value == current.Value {
			nodeToRemove = current
			break
		} else if value < current.Value {
			parentEdgeLink = &current.Left
			current = current.Left
		} else {
			parentEdgeLink = &current.Right
			current = current.Right
		}
	}

	if nodeToRemove == nil {
		return tree
	}

	if nodeToRemove.Left == nil && nodeToRemove.Right == nil {
		*parentEdgeLink = nil

		// on the right side, the most left value needs to be found
		// on the left side, the most right value needs to be found
	} else if nodeToRemove.Right != nil {
		current = nodeToRemove.Right
		parentEdgeLink = &nodeToRemove.Right
		for current.Left != nil {
			parentEdgeLink = &current.Left
			current = current.Left
		}
		nodeToRemove.Value = current.Value
		*parentEdgeLink = current.Right
	} else if nodeToRemove.Left != nil {
		current = nodeToRemove.Left
		parentEdgeLink = &nodeToRemove.Left
		for current.Right != nil {
			parentEdgeLink = &current.Right
			current = current.Right
		}
		nodeToRemove.Value = current.Value
		*parentEdgeLink = current.Left
	}

	return tree
}
