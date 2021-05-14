package _107_rotate

import (
	"testing"
)

func TestRotate(t *testing.T) {
	//matrix := [][]int{
	//	{1,2,3},
	//	{4,5,6},
	//	{7,8,9},
	//}
	//rotate(matrix)
	//print(matrix)

	matrix1 := [][]int{
		{5, 1, 9, 11},
		{2, 4, 8, 10},
		{13, 3, 6, 7},
		{15, 14, 12, 16},
	}
	print(matrix1)
	rotate(matrix1)
	print(matrix1)
}
