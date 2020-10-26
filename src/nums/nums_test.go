package nums

import (
	"fmt"
	"testing"
)

func TestSubSets(t *testing.T) {
	fmt.Println(SubSets([]int{1, 2, 3}))
	fmt.Println(SubSetsV2([]int{1, 2, 3}))
}

func TestCombine(t *testing.T) {
	fmt.Println(Combine(4, 2))
}

func TestDeleteItem(t *testing.T) {
	fmt.Println(deleteItem([]int{1, 2, 3, 4}, 2))
}
