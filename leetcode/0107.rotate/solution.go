package _107_rotate

import (
	"fmt"
	"strconv"
)

func rotate(matrix [][]int) {
	rotate0(matrix, 0, len(matrix)-1)
}

func rotate0(matrix [][]int, i, index int) {
	if i >= index {
		return
	}
	for k := 0; k < index-i; k++ {
		matrix[i][i+k], matrix[k+i][index], matrix[index][index-k], matrix[index-k][i] = matrix[index-k][i], matrix[i][i+k], matrix[k+i][index], matrix[index][index-k]
	}
	i++
	index--
	rotate0(matrix, i, index)
}

func print(matrix [][]int) {
	result := "[\n"
	for _, a := range matrix {
		r := "\t["
		for i, v := range a {
			r += strconv.FormatInt(int64(v), 10)
			if i < len(a)-1 {
				r += ", "
			}
		}
		r += "],\n"
		result += r
	}
	result += "]"
	fmt.Println(result)
}
