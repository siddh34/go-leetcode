package maxactivesectionsaftertrade

import "math"

func maxActiveSectionsAfterTrade(sections string) int {
	record := []int{}
	count := 0
	maxSum := 0
	posSum := 0

	for i := range sections {
		if i != 0 && sections[i] != sections[i-1] {
			record = append(record, count)
			count = 0
		}
		
		if sections[i] == '0' {
			count--
		} else {
			count++
		}
	}

	record = append(record, count)

	for _, val := range record {
		if val > 0 {
			posSum += val
		}
	}

	maxSum = posSum
	m := len(record)

	for i := 0; i < m - 1; i++ {
		if i > 0 && record[i-1] < 0 && record[i+1] < 0 && record[i] > 0 {
			tempSum := posSum + int(math.Abs(float64(record[i-1] + record[i+1])))
			maxSum = int(math.Max(float64(maxSum), float64(tempSum)))
		}
	}

	return maxSum
}