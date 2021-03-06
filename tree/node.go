package tree

import "fmt"

type TreeNode struct {
	Value int
	Left,Right *TreeNode
}

func (node TreeNode) Print() {
	fmt.Print(node.Value," ")
}

func (node *TreeNode) SetValue(value int) {
	if node == nil {
		fmt.Println("Setting Value to nil" +
			"node. Ignored.")
		return
	}
	node.Value = value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}

func CreateNode(value int) *TreeNode {
	return &TreeNode{Value:value}
}