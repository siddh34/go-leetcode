package maximumlengthsubstringwithtwooccurrences

func maxLengthBetweenEqualCharacters(s string) int {
	maxLength := 0
	for i := 0; i < len(s); i++ {
		occurrences := make(map[rune]int)
		for j := i; j < len(s); j++ {
			char := rune(s[j])
			occurrences[char]++
			if occurrences[char] > 2 {
				break
			}
			length := j - i + 1
			if length > maxLength {
				maxLength = length
			}
		}
	}
	return maxLength
}
