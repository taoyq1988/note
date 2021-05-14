package _7_solveSudoku

import (
	"fmt"
	"strconv"
)

var p = byte('.')

// 原题描述太智障，没有说明输入是字符串数字，函数类型却用byte来描述，导致误以为数字就用byte表示，在结果输出上会不过
func solveSudoku(board [][]byte) {
	rowMap := make(map[int]map[int]bool)
	colMap := make(map[int]map[int]bool)
	cellMap := make(map[int]map[int]bool)
	for i := 0; i < 9; i++ {
		rowMap[i] = make(map[int]bool)
		colMap[i] = make(map[int]bool)
		cellMap[i] = make(map[int]bool)
	}
	for i, b := range board {
		for j, c := range b {
			if c != p {
				n := int(c)
				rowMap[i][n] = true
				colMap[j][n] = true
				cell := (i / 3) + (j / 3 * 3)
				cellMap[cell][n] = true
			}
		}
	}
	solve(board, 0, 0, rowMap, colMap, cellMap)
}

func solve(board [][]byte, r, c int, rowMap, colMap, cellMap map[int]map[int]bool) bool {
	if r == 9 && c == 0 {
		return true
	}
	if board[r][c] == p {
		for i := 1; i <= 9; i++ {
			cell := (r / 3) + (c / 3 * 3)
			if (!rowMap[r][i]) && (!colMap[c][i]) && (!cellMap[cell][i]) {
				board[r][c] = byte(i)
				rowMap[r][i] = true
				colMap[c][i] = true
				cellMap[cell][i] = true
				if c == 8 {
					if solve(board, r+1, 0, rowMap, colMap, cellMap) {
						return true
					}
				} else {
					if solve(board, r, c+1, rowMap, colMap, cellMap) {
						return true
					}
				}
				board[r][c] = p
				rowMap[r][i] = false
				colMap[c][i] = false
				cellMap[cell][i] = false
			}
		}
		return false
	} else {
		if c == 8 {
			return solve(board, r+1, 0, rowMap, colMap, cellMap)
		} else {
			return solve(board, r, c+1, rowMap, colMap, cellMap)
		}
	}
}

func sudoPrint(board [][]byte) {
	result := ""
	for i, row := range board {
		r := "|"
		for j, col := range row {
			if col == p {
				r += "."
			} else {
				r += strconv.FormatInt(int64(col), 10)
			}
			if (j+1)%3 == 0 && j != 8 {
				r += "|"
			}
			if j < len(row)-1 {
				r += " "
			}
		}
		r += "|"
		result += r + "\n"
		if (i+1)%3 == 0 && i != 8 {
			result += "---------------------\n"
		}
	}
	fmt.Println(result)
}
