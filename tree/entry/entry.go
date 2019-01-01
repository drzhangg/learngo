package main

import (
	"fmt"
	"learngo/tree"
)

type myTreeNode struct {
	node *tree.TreeNode
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrder()
	right.postOrder()
	myNode.node.Print()
}

func main()  {
	var root tree.TreeNode

	root = tree.TreeNode{Value:3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5,nil,nil}
	root.Right.Left = 	new(tree.TreeNode)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)

	root.Traverse()
	fmt.Println()
	myRoot := myTreeNode{&root}
	myRoot.postOrder()
	fmt.Println()
}
