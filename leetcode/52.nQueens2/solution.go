package _1_nQueens

var (
	rowList []bool
	colList []bool
	x1List  []bool
	x2Map   map[int]bool
	count   int
)

func totalNQueens(n int) int {
	rowList = make([]bool, n+1)
	colList = make([]bool, n+1)
	x1List = make([]bool, 2*n)
	x2Map = make(map[int]bool)
	qs := make([][]bool, n)
	count = 0
	for i := 0; i < n; i++ {
		qs[i] = make([]bool, n)
	}
	solve(0, n, qs)
	return count
}

func solve(row, n int, qs [][]bool) {
	if row == n {
		count++
	}
	for i := 0; i < n; i++ {
		if isSafe(row, i) {
			qs[row][i] = true
			rowList[row] = true
			colList[i] = true
			x1List[row+i] = true
			x2Map[row-i] = true
			solve(row+1, n, qs)
			qs[row][i] = false
			rowList[row] = false
			colList[i] = false
			x1List[row+i] = false
			x2Map[row-i] = false
		}
	}
}

func isSafe(r, c int) bool {
	if rowList[r] || colList[c] || x1List[r+c] || x2Map[r-c] {
		return false
	}
	return true
}
