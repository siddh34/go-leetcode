package generate_parentheses

func backtrack(result *[]string, current string, open int, close int, max int) {
	if len(current) == max*2 {
		*result = append(*result, current)
		return
	}

	if open < max {
		backtrack(result, current + "(", open + 1, close, max)
	}

	if close < open {
		backtrack(result, current + ")", open, close + 1, max)
	}
}


func generateParenthesis(n int) []string {
	result := []string{}
	backtrack(&result, "", 0, 0, n)
    return result
}