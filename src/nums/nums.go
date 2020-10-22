package nums

import "sort"

func SubSets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}

	cur := nums[len(nums)-1]
	res := SubSets(nums[:len(nums)-1])
	l := len(res)
	for i := 0; i < l; i++ {
		res = append(res, append(res[i], cur))
	}
	return res
}

func SubSetsV2(nums []int) [][]int {
	var (
		backtrack func(nums []int, track []int, res *[][]int)

		res   = make([][]int, 0)
		track = make([]int, 0)
	)

	backtrack = func(nums []int, track []int, res *[][]int) {
		tmp := make([]int, len(track))
		copy(tmp, track)
		*res = append(*res, tmp)

		for i := 0; i < len(nums); i++ {
			track = append(track, nums[i])
			backtrack(nums[i+1:], track, res)
			track = track[:len(track)-1]
		}
	}

	backtrack(nums, track, &res)

	return res
}

func Combine(n, k int) [][]int {
	var (
		backtrack func(n, pos, k int, track []int, res *[][]int)

		res   = make([][]int, 0)
		track = make([]int, 0)
	)

	backtrack = func(n, pos, k int, track []int, res *[][]int) {
		if len(track) == k {
			ans := make([]int, k)
			copy(ans, track)
			*res = append(*res, ans)
			return
		}

		for i := pos; i <= n; i++ {
			track = append(track, i)
			backtrack(n, i+1, k, track, res)
			track = track[:len(track)-1]
		}
	}

	backtrack(n, 1, k, track, &res)

	return res
}
