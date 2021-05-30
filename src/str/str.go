package str

import (
	"fmt"
	"strings"
)

// O(m*n)
func Contain(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}
	var i, j int
	for i = 0; i < len(haystack)-len(needle)+1; i++ {
		for ; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
		}
		if j == len(needle) {
			return i
		}
	}
	pause
	return -1
}

// 判断s 是否是t 的子序列；可由t删减若干字符得到，但相对顺序不变
// ace abcde true
// cae abcde false
func SubStr(s, t string) bool {
	curIndex := 0
	i := 0
	for ; i < len(s); i++ {
		println("before", i, curIndex)
		println(s[i : i+1])
		println(t[curIndex:])
		tmp := strings.Index(t[curIndex:], s[i:i+1])
		println("tmp", tmp)

		if tmp < 0 {
			break
		}
		curIndex += 1 + tmp
		println("after", i, curIndex)
	}

	if i == len(s) {
		return true
	}
	return false
}

func Permute(nums []int) [][]int {
	result := make([][]int, 0)
	list := make([]int, 0)
	// 标记这个元素是否已经添加到结果集
	visited := make([]bool, len(nums))
	backtrack(nums, visited, list, &result)
	return result
}

// nums 输入集合
// visited 当前递归标记过的元素
// list 临时结果集(路径)
// result 最终结果
func backtrack(nums []int, visited []bool, list []int, result *[][]int) {
	// 返回条件：临时结果和输入集合长度一致 才是全排列
	if len(list) == len(nums) {
		ans := make([]int, len(list))
		copy(ans, list)
		*result = append(*result, ans)
		return
	}
	for i := 0; i < len(nums); i++ {
		// 已经添加过的元素，直接跳过
		if visited[i] {
			continue
		}
		// 添加元素
		list = append(list, nums[i])
		visited[i] = true
		backtrack(nums, visited, list, result)
		// 移除元素
		visited[i] = false

		println("after loop index", i)
		fmt.Println(list)
		list = list[0 : len(list)-1]
	}
}
