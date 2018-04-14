package segmentator

func min(values ...int) int {
	a := values[0]
	for i := range values {
		if i < a {
			a = i
		}
	}
	return a
}

func max(values ...int) int {
	a := values[0]
	for i := range values {
		if i > a {
			a = i
		}
	}
	return a
}

func abs(value int) int {
	if value > 0 {
		return value
	}
	return -value
}
