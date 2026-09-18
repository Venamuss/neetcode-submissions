func maxArea(heights []int) int {
	left := 0
	right := len(heights) - 1
	res := 0

	for left <= right {
		res = max(res, min(heights[left],
		heights[right]) * (right - left))

		if heights[left] < heights[right] {
			left++
		} else {
			right--
		}
	}

	return res
}
