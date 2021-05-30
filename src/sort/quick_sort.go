package sort

func Quicksort(nums []int, start, end int) {
	if start >= end {
		return
	}
	p := partition(nums, start, end)
	Quicksort(nums, start, p-1)
	Quicksort(nums, p+1, end)
}

func partition(nums []int, start, end int) int {
	base := nums[end]
	i := start
	for j := start; j < end; j++ {
		if nums[j] < base {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}
	nums[i], nums[end] = nums[end], nums[i]
	return i
}

func Mergesort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	mid := len(nums) / 2
	left := Mergesort(nums[:mid])
	right := Mergesort(nums[mid:])

	return merge(left, right)
}

func merge(nums1, nums2 []int) []int {
	res := make([]int, 0)
	i, j := 0, 0
	for i < len(nums1) && j < len(nums2) {
		if nums1[i] <= nums2[j] {
			res = append(res, nums1[i])
			i++
		} else {
			res = append(res, nums2[j])
			j++
		}
	}

	if i < len(nums1) {
		res = append(res, nums1[i:]...)
	}
	if j < len(nums2) {
		res = append(res, nums2[j:]...)
	}

	return res
}
