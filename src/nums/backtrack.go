package nums

import (
	"sort"
	"strings"
)

// 递归解法
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

// 回溯解法
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

// 回溯法 有重复数字的序列
func SubsetsDup(nums []int) [][]int {
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
			// 去重 同一树层去重
			if i > 0 && nums[i] == nums[i-1] {
				continue
			}

			track = append(track, nums[i])
			backtrack(nums[i+1:], track, res)
			track = track[:len(track)-1]
		}
	}

	sort.Ints(nums)
	backtrack(nums, track, &res)

	return res
}

// 回溯法 组合问题
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

// 非重复序列
func Permutation(nums []int) [][]int {
	var (
		backtrack func(nums []int, visited []bool, track []int, res *[][]int)

		res     = make([][]int, 0)
		visited = make([]bool, len(nums))
		track   = make([]int, 0)
	)

	backtrack = func(nums []int, visited []bool, track []int, res *[][]int) {
		if len(track) == len(nums) {
			ans := make([]int, len(nums))
			copy(ans, track)
			*res = append(*res, ans)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}
			visited[i] = true
			track = append(track, nums[i])
			backtrack(nums, visited, track, res)
			track = track[:len(track)-1]
			visited[i] = false
		}
	}

	backtrack(nums, visited, track, &res)

	return res
}

// 有重复数字的序列
func PermutationV2(nums []int) [][]int {
	var (
		backtrack func(nums []int, visited []bool, track []int, res *[][]int)

		res     = make([][]int, 0)
		visited = make([]bool, len(nums))
		track   = make([]int, 0)
	)

	backtrack = func(nums []int, visited []bool, track []int, res *[][]int) {
		if len(track) == len(nums) {
			ans := make([]int, len(nums))
			copy(ans, track)
			*res = append(*res, ans)
			return
		}

		for i := 0; i < len(nums); i++ {
			if visited[i] {
				continue
			}

			// 去重 同一树层去重 visited[i-1] == false
			if i > 0 && nums[i] == nums[i-1] && !visited[i-1] {
				continue
			}

			//// 去重 同一树枝去重 visited[i-1] == false
			//if i >0 && nums[i] == nums[i-1] && visited[i-1] {
			//	continue
			//}

			visited[i] = true
			track = append(track, nums[i])
			backtrack(nums, visited, track, res)
			track = track[:len(track)-1]
			visited[i] = false
		}
	}

	sort.Ints(nums)
	backtrack(nums, visited, track, &res)

	return res
}

func RestoreIpAddresses(s string) []string {
	var (
		res   = make([]string, 0)
		track = make([]string, 0)
	)
	backtrack(s, len(s), 0, 0, track, &res)
	return res
}

func backtrack(s string, total, lastIndex int, segLen int, track []string, res *[]string) {

	// 每段ip最多3位，
	if total-lastIndex-segLen > (4-len(track))*3 || total-lastIndex-segLen < 4-len(track) {
		return
	}

	if len(track) == 4 {
		ts := ""
		for i := 0; i < len(track); i++ {
			ts += track[i] + "."
		}
		*res = append(*res, strings.TrimSuffix(ts, "."))
		return
	}

	for i := 1; i <= 3; i++ {
		if s[lastIndex:lastIndex+i] > "255" || (i > 1 && s[lastIndex] == '0') {
			continue
		}
		track = append(track, s[lastIndex:lastIndex+i])
		backtrack(s, len(s), lastIndex+i, i, track, res)
		track = track[:len(track)-1]
	}
}
