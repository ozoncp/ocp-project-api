package utils

import (
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"os"
)

// SplitSlice converts slice sl to slice of slices with n-size chunks
func SplitSlice(sl []interface{}, n int) [][]interface{} {
	if n <= 0 || sl == nil {
		return nil
	}
	if n >= len(sl) {
		return [][]interface{}{sl}
	}

	count := len(sl) / n
	addition := 0
	if len(sl)%n != 0 {
		addition = 1
	}
	var result = make([][]interface{}, 0, count+addition)

	for i := 0; i < count; i += 1 {
		result = append(result, sl[i*n:i*n+n])
	}

	if addition != 0 {
		result = append(result, sl[count*n:])
	}

	return result
}

// ReverseKeyValue converts map m which maps key to value to map which maps value to key
func ReverseKeyValue(m map[string]int) map[int]string {
	if m == nil {
		return nil
	}

	var result = make(map[int]string, len(m))

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

	var blackMap = map[interface{}]struct{}{}
	for _, item := range blackList {
		blackMap[item] = struct{}{}
	}

	for _, item := range sl {
		if _, found := blackMap[item]; !found {
			result = append(result, item)
		}
	}

	return result
}

// LoopOpenClose open file with name fileName and make some magic with it in loop (usage functor and defer in loop)
func LoopOpenClose(fileName string, msg string, count int) {
	for i := 0; i < count; i++ {
		func() {
			f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
			if err != nil {
				return
			}
			defer f.Close()
			// make something with file
			f.WriteString(fmt.Sprintf("%s: loop count %d\n", msg, i))
		}()
	}
}

// SplitToBulks converts slice sl to slice of slices with butchSize-size chunks
func SplitToBulks(sl []models.Artifact, butchSize uint) [][]models.Artifact {
	if butchSize <= 0 || sl == nil {
		return nil
	}
	if int(butchSize) >= len(sl) {
		return [][]models.Artifact{sl}
	}

	count := len(sl) / int(butchSize)
	addition := 0
	if len(sl)%int(butchSize) != 0 {
		addition = 1
	}
	var result = make([][]models.Artifact, 0, count+addition)

	for i := 0; i < count; i += 1 {
		result = append(result, sl[i*int(butchSize):i*int(butchSize)+int(butchSize)])
	}

	if addition != 0 {
		result = append(result, sl[count*int(butchSize):])
	}

	return result
}

// SliceToMap convert slice of structs sl to map with struct id as key and struct as value
func SliceToMap(sl []models.Artifact) map[uint64]models.Artifact {
	if sl == nil {
		return nil
	}

	result := make(map[uint64]models.Artifact, len(sl))
	for _, item := range sl {
		if _, found := result[item.Id()]; found {
			panic(fmt.Sprintf("Invalid slice of models: model with id %d occurs more than once", item.Id()))
		}

		result[item.Id()] = item
	}

	return result
}
