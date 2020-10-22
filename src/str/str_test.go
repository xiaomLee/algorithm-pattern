package str

import (
	"fmt"
	"testing"
)

func TestContain(t *testing.T) {
	str := "abcdeqqqqqwww"
	needle := "www"
	println(Contain(str, needle))
}

func TestSubStr(t *testing.T) {
	s := "ccae"
	ss := "abcdeeeca"
	println(SubStr(s, ss))
}

func TestPermute(t *testing.T) {
	fmt.Println(Permute([]int{1, 2, 3}))
}
