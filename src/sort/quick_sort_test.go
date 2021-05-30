package sort

import (
	"fmt"
	"testing"
)

func TestQuicksort(t *testing.T) {
	nums := []int{1, 3, 2, 5, 6, 4}
	Quicksort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func TestMergesort(t *testing.T) {
	nums := []int{1, 3, 2, 5, 6, 4}
	res := Mergesort(nums)
	fmt.Println(res)
}
