package utils

import "fmt"

// SplitSlice converts slice sl to slice of slices with n-size chunks
func SplitSlice(sl []interface{}, n int) [][]interface{} {
	var result [][]interface{}

	if n <= 0 || sl == nil {
		return nil
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

// ReverseKeyValue converts map m which maps key to value to map which maps value to key
func ReverseKeyValue(m map[string]int) map[int]string {
	var result = map[int]string{}

	if m == nil {
		return nil
	}

	for key, value := range m {
		if _, found := result[value]; found {
			panic(fmt.Sprintf("Unsupported value: %v occurs more than once", value))
		}
		result[value] = key
	}

	return result
}

// FilterSlice returns slice with elements from sl which do not appear on the slice blackList
func FilterSlice(sl []interface{}, blackList []interface{}) []interface{} {
	var result []interface{}

	var blackMap = map[interface{}]bool{}
	for _,item := range blackList {
		blackMap[item] = true
	}

	for _,item := range sl {
		if _,found := blackMap[item]; !found {
			result = append(result, item)
		}
	}

	return result
}
