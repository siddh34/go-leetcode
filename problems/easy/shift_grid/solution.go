package shiftgrid

func shiftGrid(grid [][]int, k int) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	totalElements := rows * cols

	flatGrid := make([]int, totalElements)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			flatGrid[i*cols+j] = grid[i][j]
		}
	}

	k = k % totalElements

	result := make([][]int, rows)
	for i := range result {
		result[i] = make([]int, cols)
	}

	for i := 0; i < totalElements; i++ {
		newIndex := (i + k) % totalElements
		result[newIndex/cols][newIndex%cols] = flatGrid[i]
	}

	return result
}