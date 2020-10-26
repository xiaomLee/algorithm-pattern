package tree

import (
	"fmt"
)

const (
	NULL = -1
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Traversal
func PreTraversal(root *TreeNode) {
	if root == nil {
		fmt.Printf("%d,", NULL)
		return
	}
	fmt.Printf("%d,", root.Val)
	PreTraversal(root.Left)
	PreTraversal(root.Right)
}

func MidTraversal(root *TreeNode) {
	if root == nil {
		fmt.Printf("%d,", NULL)
		return
	}
	PreTraversal(root.Left)
	fmt.Printf("%d,", root.Val)
	PreTraversal(root.Right)
}

// 思路：通过stack 保存已经访问的元素，用于原路返回
func InorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	for len(stack) > 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left // 一直向左
		}
		// 弹出
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right
	}
	return result
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		fmt.Printf("%d,", NULL)
		return
	}
	PreTraversal(root.Left)

	PreTraversal(root.Right)

	fmt.Printf("%d,", root.Val)
}

func PreOderDeserialize(data []int) *TreeNode {
	var deserialize func(data []int, index *int) *TreeNode

	deserialize = func(data []int, index *int) *TreeNode {
		if *index >= len(data) {
			return nil
		}

		val := data[*index]
		*index++
		if val == NULL {
			return nil
		}

		root := &TreeNode{Val: val}
		root.Left = deserialize(data, index)
		root.Right = deserialize(data, index)

		return root
	}

	index := 0
	return deserialize(data, &index)
}

// 后序遍历 反序列化， 最后一位是root
func PostOderDeserialize(data []int) *TreeNode {
	var deserialize func(data []int, index *int) *TreeNode

	deserialize = func(data []int, index *int) *TreeNode {
		if *index < 0 {
			return nil
		}

		val := data[*index]
		*index--
		if val == NULL {
			return nil
		}

		root := &TreeNode{Val: val}
		root.Left = deserialize(data, index)
		root.Right = deserialize(data, index)

		return root
	}

	index := len(data) - 1
	return deserialize(data, &index)
}

func (root *TreeNode) PreOrderSerialize() []int {
	var (
		traversal func(node *TreeNode, res *[]int)
		res       = make([]int, 0)
	)

	traversal = func(node *TreeNode, res *[]int) {
		if node == nil {
			*res = append(*res, NULL)
			return
		}

		*res = append(*res, node.Val)
		traversal(node.Left, res)
		traversal(node.Right, res)
	}

	traversal(root, &res)

	return res
}

func (root *TreeNode) PostOrderSerialize() []int {
	var (
		traversal func(node *TreeNode, res *[]int)
		res       = make([]int, 0)
	)

	traversal = func(node *TreeNode, res *[]int) {
		if node == nil {
			*res = append(*res, NULL)
			return
		}

		traversal(node.Left, res)
		traversal(node.Right, res)
		*res = append(*res, node.Val)
	}

	traversal(root, &res)

	return res
}

func MaxDepth(root *TreeNode) int {
	var (
		queue = make([]*TreeNode, 0)
		depth = 0
	)

	queue = append(queue, root)
	for len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// 判断target, 此题全部遍历完即可
		fmt.Print(cur.Val, ",")

		// 将所有相邻节点加入队列
		if cur.Left != nil {
			queue = append(queue, cur.Left)
		}
		if cur.Right != nil {
			queue = append(queue, cur.Right)
		}

		// step +1
		depth++
	}

	return depth
}
