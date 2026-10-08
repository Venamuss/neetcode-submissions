type TimeValue struct {
	values     []string
	timestamps []int
}

type TimeMap struct {
	data map[string]TimeValue
}

func Constructor() TimeMap {
	return TimeMap{
		data: make(map[string]TimeValue),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	possiblePrev, ok := this.data[key]
	var prevTimestamps []int
	var prevValues []string

	if ok {
		prevTimestamps = possiblePrev.timestamps
		prevValues = possiblePrev.values
	} else {
		prevTimestamps = make([]int, 0)
		prevValues = make([]string, 0)
	}

	timeValue := TimeValue{
		values:     append(prevValues, value),
		timestamps: append(prevTimestamps, timestamp),
	}

	this.data[key] = timeValue
}

func (this *TimeMap) Get(key string, timestamp int) string {
	data := this.data[key]

	foundedTimestamp := searchTimestampIndex(data.timestamps, timestamp)
	if foundedTimestamp == -1 {
		//fmt.Println("Returned \"\"")
		return ""
	}

	res := ""

	for foundedTimestamp >= 0 {
		res = data.values[foundedTimestamp]

		if res == "" {
			foundedTimestamp--
		} else {
			break
		}
	}

	return res
}

func searchTimestampIndex(timestamps []int, target int) int {
	if len(timestamps) == 0 {
		//fmt.Println("return -1 by len")
		return -1
	}

	lastElement := timestamps[len(timestamps)-1]
	if lastElement < target {
		return len(timestamps) - 1
	}
	if timestamps[0] > target {
		return -1
	}

	l := 0
	r := len(timestamps) - 1

	for l <= r {
		mid := l + (r-l)/2

		if timestamps[mid] == target {
			return mid
		}
		if timestamps[mid] > target {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	return r
}