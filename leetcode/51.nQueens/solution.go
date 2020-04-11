package _1_nQueens

var (
	rowList []bool
	colList []bool
	x1List  []bool
	x2Map   map[int]bool
)

func solveNQueens(n int) [][]string {
	results := make([][]string, 0)
	rowList = make([]bool, n+1)
	colList = make([]bool, n+1)
	x1List = make([]bool, 2*n)
	x2Map = make(map[int]bool)
	qs := make([][]bool, n)
	for i := 0; i < n; i++ {
		qs[i] = make([]bool, n)
	}
	solve(&results, 0, n, qs)
	return results
}

func solve(results *[][]string, row, n int, qs [][]bool) {
	if row == n {
		*results = append(*results, convertQS(qs))
	}
	for i := 0; i < n; i++ {
		if isSafe(qs, row, i, n) {
			qs[row][i] = true
			rowList[row] = true
			colList[i] = true
			x1List[row+i] = true
			x2Map[row-i] = true
			solve(results, row+1, n, qs)
			qs[row][i] = false
			rowList[row] = false
			colList[i] = false
			x1List[row+i] = false
			x2Map[row-i] = false
		}
	}
}

func isSafe(qs [][]bool, r, c, n int) bool {
	if rowList[r] || colList[c] || x1List[r+c] || x2Map[r-c] {
		return false
	}
	return true
}

func convertQS(qs [][]bool) []string {
	result := make([]string, len(qs))
	for i, r := range qs {
		str := ""
		for _, q := range r {
			if q {
				str += "Q"
			} else {
				str += "."
			}
		}
		result[i] = str
	}
	return result
}
