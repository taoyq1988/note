package _7_solveSudoku

import (
	"fmt"
	"testing"
)

func TestSolveSudoku(t *testing.T) {
	board := [][]byte{
		[]byte{5,3,p,p,7,p,p,p,p},
		[]byte{6,p,p,1,9,5,p,p,p},
		[]byte{p,9,8,p,p,p,p,6,p},
		[]byte{8,p,p,p,6,p,p,p,3},
		[]byte{4,p,p,8,p,3,p,p,1},
		[]byte{7,p,p,p,2,p,p,p,6},
		[]byte{p,6,p,p,p,p,2,8,p},
		[]byte{p,p,p,4,1,9,p,p,5},
		[]byte{p,p,p,p,8,p,p,7,9},
	}
	sudoPrint(board)

	solveSudoku(board)
	sudoPrint(board)
	fmt.Println(board)
}
