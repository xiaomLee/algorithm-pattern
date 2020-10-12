package str

import "testing"

func TestContain(t *testing.T) {
	str := "abcdeqqqqqwww"
	needle := "www"
	println(Contain(str, needle))
}
