package maxproductoftwodigits

import (
	"slices"
)

func maxProduct(n int) int {
	numRecords := []int{}
	temp := n
	for temp > 0 {
		numRecords = append(numRecords, temp%10)
		temp = temp / 10
	}

	slices.Sort(numRecords)
	return numRecords[len(numRecords)-1] * numRecords[len(numRecords)-2]
}
