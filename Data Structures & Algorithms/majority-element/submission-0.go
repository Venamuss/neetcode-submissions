func majorityElement(nums []int) int {
    hash := map[int]int{}
	majority := -1
	mostValue := -1

	for _, v := range nums {
		hash[v]++
	}

	for k, v := range hash {
		if v > mostValue {
			majority = k
			mostValue = v
		}
	}	

	return majority
}
