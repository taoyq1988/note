package _1_nextPermutation

import (
	"fmt"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	r := []int{5, 4, 7, 5, 3, 2}
	nextPermutation(r)
	fmt.Println(r)

	r = []int{2, 3, 1}
	nextPermutation(r)
	fmt.Println(r)
}
