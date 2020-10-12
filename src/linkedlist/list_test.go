package linkedlist

import "testing"

func defaultSortedList() *ListNode {
	return NewIntList([]int{1, 2, 2, 2, 2, 3, 3, 4, 5, 5})
}

func TestNewIntList(t *testing.T) {
	list := defaultSortedList()
	list.Print()
	println(list.String())
}

func TestDeleteDuplicates(t *testing.T) {
	head := defaultSortedList()
	head.Print()
	DeleteDuplicates(head)
	head.Print()
}

func TestDeleteDupKeepOnce(t *testing.T) {
	head := defaultSortedList()
	head.Print()
	DeleteDupKeepOnce(head)
	head.Print()
}

func TestReverseList(t *testing.T) {
	head := defaultSortedList()
	head.Print()
	ReverseList(head).Print()
}
