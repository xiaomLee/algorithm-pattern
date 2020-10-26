package nums

import (
	"fmt"
	"testing"
)

func TestSubSets(t *testing.T) {
	fmt.Println(SubSets([]int{1, 2, 3}))
	fmt.Println(SubSetsV2([]int{1, 2, 3}))
	fmt.Println(SubsetsDup([]int{1, 2, 2}))
}

func TestCombine(t *testing.T) {
	fmt.Println(Combine(4, 2))
}

func TestPermutation(t *testing.T) {
	fmt.Println(Permutation([]int{1, 2, 3}))
	fmt.Println(PermutationV2([]int{1, 2, 2}))
}

func TestRestoreIpAddresses(t *testing.T) {
	fmt.Println(RestoreIpAddresses("25525511135"))
}
