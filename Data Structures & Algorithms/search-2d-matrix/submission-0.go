func searchMatrix(matrix [][]int, target int) bool {
	totalLen := len(matrix) * len(matrix[0])
	matrixLen := len(matrix[0])

	start := 0
	end := totalLen - 1

	for start <= end {
		mid := start + (end - start) / 2
		midNum := matrix[mid / matrixLen][mid % matrixLen]

		if midNum == target {
			return true
		}
		if midNum > target {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return false
}
