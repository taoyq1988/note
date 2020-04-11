package _1_nQueens

import (
	"fmt"
	"testing"
)

func TestQS(t *testing.T) {
	rs := solveNQueens(4)
	for _, r1 := range rs {
		for _, r2 := range r1 {
			fmt.Println(r2)
		}
		fmt.Printf("\n")
	}
}
