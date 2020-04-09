package _2_brackets

var left = "("
var right = ")"

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	gen(&result, "", 0, 0, n)
	return result
}

func gen(results *[]string, str string, ln, rn, n int) {
	if ln < rn || ln > n || rn > n {
		return
	}
	if ln == n && rn == n {
		*results = append(*results, str)
		return
	}
	if ln < n {
		gen(results, str+left, ln+1, rn, n)
	}
	if rn < n {
		gen(results, str+right, ln, rn+1, n)
	}
}
