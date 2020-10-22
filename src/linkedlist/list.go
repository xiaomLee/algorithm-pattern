package linkedlist

import (
	"fmt"
)

type ListNode struct {
	val  interface{}
	pre  *ListNode
	next *ListNode
}

func NewIntList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := &ListNode{}
	head := dummy
	for _, val := range nums {
		node := &ListNode{
			val: val,
		}
		node.pre = head
		head.next = node
		head = node
	}

	return dummy.next
}

func (head *ListNode) Print() {
	curr := head
	for curr != nil {
		fmt.Printf("%v->", curr.val)
		curr = curr.next
	}
	println("NULL")
}

func (head *ListNode) String() string {
	s := ""
	curr := head
	for curr != nil {
		s += fmt.Sprintf("%v->", curr.val)
		curr = curr.next
	}
	return s + "NULL"
}

// 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
func DeleteDuplicates(head *ListNode) *ListNode {
	current := head
	for current != nil {
		for current.next != nil && current.next.val == current.val {
			current.next = current.next.next
		}
		current = current.next
	}

	return head
}

// 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中   没有重复出现的数字。
// 只出现一次的节点
func DeleteDupKeepOnce(head *ListNode) *ListNode {
	dummy := &ListNode{}
	dummy.next = head
	current := dummy

	var rmVal interface{}
	for current.next != nil && current.next.next != nil {
		if current.next.val == current.next.next.val {
			rmVal = current.next.val
			// 指针直接跳到第三个节点
			current.next = current.next.next.next
			// 顺序删除
			for current.next != nil && current.next.val == rmVal {
				current.next = current.next.next
			}
		} else {
			current = current.next
		}
	}

	return dummy.next
}

// 反转一个单链表。
// 思路：用一个 prev 节点保存向前指针，temp 保存向后的临时指针
func ReverseList(head *ListNode) *ListNode {
	var pre *ListNode
	for head != nil {
		temp := head.next
		head.next = pre
		pre = head
		head = temp
	}

	return pre
}

// 反转从位置  *m*  到  *n*  的链表。请使用一趟扫描完成反转。
func ReverseListMN(head *ListNode, m, n int) *ListNode {
	if m < 0 || m > n {
		return head
	}

	dummy := &ListNode{}
	dummy.next = head

	i := 0
	current := dummy
	// 遍历跳过
	nodeMPre := current
	for current != nil && i < m {
		nodeMPre = current
		current = current.next
		i++
	}

	// 遍历
	var pre *ListNode
	nodeM := current
	for current != nil && i <= n {
		temp := current.next
		current.next = pre
		pre = current
		current = temp
		i++
	}

	nodeMPre.next = pre
	nodeM.next = current

	return dummy.next
}

// 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于  *x*  的节点都在大于或等于  *x*  的节点之前。
// 思路：将大于 x 的节点，放到另外一个链表，最后连接这两个链表, 无需排序
func PartitionX(head *ListNode, x int) *ListNode {
	headDummy := &ListNode{}
	tailDummy := &ListNode{}

	headDummy.next = head
	current := headDummy
	tail := tailDummy
	for current.next != nil {
		if current.next.val.(int) < x {
			current = current.next
		} else {
			tail.next = current.next
			tail = tail.next
			current.next = current.next.next
		}
	}

	tail.next = nil
	current.next = tailDummy.next

	return headDummy.next
}

func Xxx() {
	type A struct {
		x int
		y int
		z []int
	}
	a := &A{
		x: 0,
		y: 1,
		z: make([]int, 0),
	}
	a.z = append(a.z, 1)

	fmt.Printf("a.ptr:%p a.x:%d a.y:%d, a.z:%v \n", a, a.x, a.y, a.z)

	c := a
	a.x = 10
	a.y = 11
	c.z = append(c.z, 10)
	fmt.Printf("c.ptr:%p c.x:%d c.y:%d, c.z:%v \n", c, c.x, c.y, c.z)
	fmt.Printf("a.ptr:%p a.x:%d a.y:%d, a.z:%v \n", a, a.x, a.y, a.z)
}
