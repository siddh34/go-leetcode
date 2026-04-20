package regexmatch

// https://leetcode.com/problems/regular-expression-matching/description/

func isMatch(s string, p string) bool {
	memo := make(map[string]bool)
	return dp(s, p, 0, 0, memo)
}

func dp(s string, p string, i int, j int, memo map[string]bool) bool {
	if j == len(p) {
		return i == len(s)
	}

	key := string(rune(i)) + "," + string(rune(j))
	if val, exists := memo[key]; exists {
		return val
	}

	match := i < len(s) && (p[j] == '.' || s[i] == p[j])

	var result bool
	if j+1 < len(p) && p[j+1] == '*' {
		result = dp(s, p, i, j+2, memo) || (match && dp(s, p, i+1, j, memo))
	} else {
		result = match && dp(s, p, i+1, j+1, memo)
	}

	memo[key] = result
	return result
}
