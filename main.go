package main

import (
	"encoding/json"
	"fmt"

	"github.com/rosricard/practice/bst"
	"github.com/stretchr/testify/require"
)

func NewBST(value int) *bst.BST {
	return &bst.BST{Value: value}
}

func TestCaseBST(t require.TestingT) {
	root := NewBST(10)
	root.Left = NewBST(5)
	root.Left.Left = NewBST(2)
	root.Left.Left.Left = NewBST(1)
	root.Left.Right = NewBST(5)
	root.Right = NewBST(15)
	root.Right.Left = NewBST(13)
	root.Right.Left.Right = NewBST(14)
	root.Right.Right = NewBST(22)

	root.Insert(12)
	require.True(t, root.Right.Left.Left.Value == 12)

	root.Remove(10)
	require.True(t, root.Contains(10) == false)
	require.True(t, root.Value == 12)

	require.True(t, root.Contains(15))
}

type BSTTest struct {
	Arguments []int  `json: "arguments"`
	Method    string `json:"method"`
}

var tests = []BSTTest{}

func main() {
	testJson := `[{"arguments": [5], "method": "insert"}, {"arguments": [15], "method": "insert"}]`

	// Unmarshal into a slice of maps
	var operations []map[string]interface{}
	err := json.Unmarshal([]byte(testJson), &operations)
	if err != nil {
		panic(err)
	}

	// Now operations is a slice of maps
	fmt.Println(operations)

}
