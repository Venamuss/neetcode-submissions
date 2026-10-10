
func removeElement(nums []int, val int) int {
    l := 0
	r := 0

	for r < len(nums) {
		if nums[r] != val {
			nums[l] = nums[r]
			l++
		}
		r++
	}

	return l
}
