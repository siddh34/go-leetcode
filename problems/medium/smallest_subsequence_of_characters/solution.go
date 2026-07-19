package smallestsubsequenceofcharacters

func smallestSubsequence(s string) string {
	counts := make(map[rune]int)
	for _, char := range s {
		counts[char]++
	}

	visited := make(map[rune]bool)
	stack := []rune{}

	for _, char := range s {
		counts[char]--
		if visited[char] {
			continue
		}

		for len(stack) > 0 && char < stack[len(stack)-1] && counts[stack[len(stack) - 1]] > 0 {
			visited[stack[len(stack)-1]] = false
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, char)
		visited[char] = true
	}

	return string(stack)
}