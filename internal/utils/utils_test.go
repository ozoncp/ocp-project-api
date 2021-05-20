package utils_test

import (
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/utils"
	"reflect"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	type TestCase struct {
		InputSlice []interface{}
		InputN int
		Output [][]interface{}
	}

	testCases := []TestCase{
		{ []interface{}{1,2,3,4,5,6}, 2, [][]interface{}{{1,2},{3,4},{5,6}} },
		{ []interface{}{1,2,3,4,5,6}, 4, [][]interface{}{{1,2,3,4},{5,6}} },
		{ []interface{}{1,2,3,4,5,6}, 7, [][]interface{}{{1,2,3,4,5,6}} },
		{ []interface{}{}, 2, [][]interface{}{{}} },
		{ []interface{}{1,2,3}, 0, nil },
		{ nil, 2, nil },
	}

	for i, c := range testCases {
		res := utils.SplitSlice(c.InputSlice, c.InputN)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i + 1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestReverseKeyValue(t *testing.T) {
	type TestCase struct {
		Input map[string]int
		Output map[int]string
	}

	var testCases = []TestCase {
		{map[string]int{"a": 1, "b": 2}, map[int]string{1: "a", 2: "b"}},
		{map[string]int{}, map[int]string{}},
		{nil, nil},
	}

	for i, c := range testCases {
		res := utils.ReverseKeyValue(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i + 1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestFilterSlice(t *testing.T) {
	type TestCase struct {
		InputSlice []interface{}
		InputBlackList []interface{}
		Output []interface{}
	}

	var testCases = []TestCase {
		{[]interface{}{"a", "b", "c"}, []interface{}{"b", "c"}, []interface{}{"a"}},
		{[]interface{}{"a", "b", "c"}, []interface{}{}, []interface{}{"a", "b", "c"}},
		{[]interface{}{"a", "b", "c"}, nil, []interface{}{"a", "b", "c"}},
		{nil, []interface{}{"b", "c"}, nil},
		{nil, nil, nil},
	}

	for i, c := range testCases {
		res := utils.FilterSlice(c.InputSlice, c.InputBlackList)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i + 1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}
