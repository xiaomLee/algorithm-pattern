package tree

import (
	"fmt"
	"testing"
)

func TestDeserialize(t *testing.T) {
	root := PreOderDeserialize([]int{1, 2, -1, 4, -1, -1, 3, -1, -1})
	PreTraversal(root)
	println()
	fmt.Println(root.PreOrderSerialize())
	fmt.Println(root.PostOrderSerialize())
}

func TestTreeNode_Serialize(t *testing.T) {
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}

	root.Left.Right = &TreeNode{Val: 4}

	PreTraversal(root)
	println()
	fmt.Println(root.PreOrderSerialize())
	fmt.Println(root.PostOrderSerialize())
}

func TestMidTraversal(t *testing.T) {
	root := PreOderDeserialize([]int{1, 2, -1, 4, -1, -1, 3, -1, -1})
	MidTraversal(root)
	fmt.Println(InorderTraversal(root))
}

func TestMaxDepth(t *testing.T) {
	root := PreOderDeserialize([]int{1, 2, -1, 4, -1, -1, 3, -1, -1})
	MaxDepth(root)
}
