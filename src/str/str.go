package str

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
	return -1
}
