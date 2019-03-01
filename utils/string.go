package utils

func Substr(str string, start int, end int) string {
	strlen := len(str)
	if start < 0 {
		start = 0
	}
	if start > strlen {
		start = strlen - 1
	}
	if end < 0 {
		if -end > strlen {
			end = 0
		} else {
			end = len(str) + end
		}
	} else {
		if end > strlen {
			end = strlen
		}
	}
	if (start > end) {
		start, end = end, start
	}
	r := string([]byte(str)[start:end])
	return r
}
