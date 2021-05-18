package utils_test

import (
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/utils"
	"reflect"
	"testing"
)

func TestUtils(t *testing.T) {
	array := []int{1,2,3,4,5,6}

	n := 2
	correctAns := [][]int{{1,2},{3,4},{5,6}}

	res := utils.SplitSlice(array, n)
	if !reflect.DeepEqual(res, correctAns) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)

	n = 4
	correctAns = [][]int{{1,2,3,4},{5,6}}

	res = utils.SplitSlice(array, n)
	if !reflect.DeepEqual(res, correctAns) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)

	n = 7
	correctAns = [][]int{{1,2,3,4,5,6}}

	res = utils.SplitSlice(array, n)
	if !reflect.DeepEqual(res, correctAns) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)

	array = []int{}
	n = 2
	correctAns = [][]int{{}}

	res = utils.SplitSlice(array, n)
	if !reflect.DeepEqual(res, correctAns) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)

	var nilArray []int
	n = 0

	res = utils.SplitSlice(nilArray, n)
	if !reflect.DeepEqual(res, make([][]int, 0)) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)

	n = 2

	res = utils.SplitSlice(nilArray, n)
	if !reflect.DeepEqual(res, make([][]int, 0)) {
		fmt.Println("Fail result: ", res)
		t.Error("fail")
		return
	}
	fmt.Println("Good result: ", res)
}
