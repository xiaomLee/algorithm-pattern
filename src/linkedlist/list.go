package linkedlist

import "fmt"

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
