package utils

func SplitSlice(sl []int, n int) [][]int {
	var result [][]int

	if n <= 0 || sl == nil {
		return make([][]int, 0)
	}

	if n >= len(sl) {
		return append(result, sl)
	}

	count := len(sl)/n
	for i := 0; i < count; i +=1 {
		result = append(result, sl[i*n:i*n + n])
	}

	if len(sl) % n != 0 {
		result = append(result, sl[count*n:])
	}

	return result
}

func ReverseKeyValue(m map[string]int) map[int]string {
	var result = map[int]string{}

	for key, value := range m {
		result[value] = key
	}

	return result
}

func FilterSlice(sl []string, blackList []string) []string {
	var result []string

	for _,str := range sl {
		found := false
		for _,blackStr := range blackList {
			if (str == blackStr) {
				found = true;
				break;
			}
		}

		if !found {
			result = append(result, str)
		}
	}

	return result
}
