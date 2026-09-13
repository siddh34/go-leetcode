package imageoverlap

func largestOverlap(img1 [][]int, img2 [][]int) int {
    result := 0

	record_a := make([][2]int, 0)
	record_b := make([][2]int, 0)

	for i := range img1 {
		for j := range img2 {
			if img1[i][j] == 1 {
				record_a = append(record_a, [2]int{i, j})
			}
			if img2[i][j] == 1 {
				record_b = append(record_b, [2]int{i, j})
			}
		}
	}

	frequency := make(map[[2]int]int)
	for _, a := range record_a {
		for _, b := range record_b {
			diff := [2]int{a[0] - b[0], a[1] - b[1]}
			frequency[diff]++
			if frequency[diff] > result {
				result = frequency[diff]
			}
		}
	}

    return result
}