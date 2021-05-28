package utils_test

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"reflect"
	"syscall"
	"testing"
)

func TestSplitSlice(t *testing.T) {
	type TestCase struct {
		InputSlice []interface{}
		InputN     int
		Output     [][]interface{}
	}

	testCases := []TestCase{
		{[]interface{}{1, 2, 3, 4, 5, 6, 7}, 2, [][]interface{}{{1, 2}, {3, 4}, {5, 6}, {7}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, 1, [][]interface{}{{1}, {2}, {3}, {4}, {5}, {6}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, 2, [][]interface{}{{1, 2}, {3, 4}, {5, 6}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, 3, [][]interface{}{{1, 2, 3}, {4, 5, 6}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, 4, [][]interface{}{{1, 2, 3, 4}, {5, 6}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, 5, [][]interface{}{{1, 2, 3, 4, 5}, {6}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, 6, [][]interface{}{{1, 2, 3, 4, 5, 6}}},
		{[]interface{}{1, 2, 3, 4, 5, 6}, 7, [][]interface{}{{1, 2, 3, 4, 5, 6}}},
		{[]interface{}{}, 2, [][]interface{}{{}}},
		{[]interface{}{1, 2, 3}, 0, nil},
		{nil, 2, nil},
	}

	for i, c := range testCases {
		res := utils.SplitSlice(c.InputSlice, c.InputN)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestReverseKeyValue(t *testing.T) {
	type TestCase struct {
		Input  map[string]int
		Output map[int]string
	}

	var testCases = []TestCase{
		{map[string]int{"a": 1, "b": 2}, map[int]string{1: "a", 2: "b"}},
		{map[string]int{}, map[int]string{}},
		{nil, nil},
	}

	for i, c := range testCases {
		res := utils.ReverseKeyValue(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}

	assert.Panics(t, func() { utils.ReverseKeyValue(map[string]int{"a": 1, "b": 1}) }, "Fail panic test: The code did not panic")
}

func TestFilterSlice(t *testing.T) {
	type TestCase struct {
		InputSlice     []interface{}
		InputBlackList []interface{}
		Output         []interface{}
	}

	var testCases = []TestCase{
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
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestLoopOpenClose(t *testing.T) {
	var (
		fileName string = "test_looping.txt"
		msg      string = "Test LoopOpenClose function"
		count    int    = 10
	)
	err := utils.LoopOpenClose("", msg, count)
	if err == nil {
		t.Errorf("Fail test: function does not return error")
		return
	}
	if err := os.Remove(fileName); err != nil {
		var pathError *fs.PathError
		if !errors.As(err, &pathError) || pathError.Err != syscall.ENOENT {
			t.Errorf("Fail test: can't remove artefact file")
			return
		}
	}
	err = utils.LoopOpenClose(fileName, msg, count)
	if err != nil {
		t.Errorf("Fail test: function returns error")
	}

	var f *os.File
	f, err = os.Open(fileName)
	if err != nil {
		t.Errorf("Fail test: no file exists")
		return
	}

	input := bufio.NewScanner(f)
	for i := 0; i < count; i++ {
		isOk := input.Scan()
		if !isOk {
			t.Errorf("Fail test: file scan error")
			return
		}

		if !reflect.DeepEqual(input.Text(), fmt.Sprintf("%s: loop count %d", msg, i)) {
			t.Errorf("Faile test: file content is invalid")
			return
		}
	}
	if os.Remove(fileName) != nil {
		t.Errorf("Fail test: can't remove artefact file")
		return
	}
}

func TestSplitToBulksProject(t *testing.T) {
	type TestCase struct {
		InputSlice []models.Artifact
		InputN     uint
		Output     [][]models.Artifact
	}

	testCases := []TestCase{
		{
			[]models.Artifact{
				models.NewProject(1, 1, "1"),
				models.NewProject(2, 2, "2"),
				models.NewProject(3, 3, "3"),
				models.NewProject(4, 4, "4"),
				models.NewProject(5, 5, "5"),
				models.NewProject(6, 6, "6"),
			},
			2,
			[][]models.Artifact{
				{
					models.NewProject(1, 1, "1"),
					models.NewProject(2, 2, "2"),
				},
				{
					models.NewProject(3, 3, "3"),
					models.NewProject(4, 4, "4"),
				},
				{
					models.NewProject(5, 5, "5"),
					models.NewProject(6, 6, "6"),
				},
			},
		},
		{[]models.Artifact{}, 2, [][]models.Artifact{{}}},
		{
			[]models.Artifact{
				models.NewProject(1, 1, "1"),
				models.NewProject(2, 2, "2"),
				models.NewProject(3, 3, "3"),
			},
			0,
			nil,
		},
		{nil, 2, nil},
	}

	for i, c := range testCases {
		res := utils.SplitToBulks(c.InputSlice, c.InputN)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestSplitToBulksRepo(t *testing.T) {
	type TestCase struct {
		InputSlice []models.Artifact
		InputN     uint
		Output     [][]models.Artifact
	}

	testCases := []TestCase{
		{
			[]models.Artifact{
				models.NewRepo(1, 1, 1, "1"),
				models.NewRepo(2, 2, 2, "2"),
				models.NewRepo(3, 3, 3, "3"),
				models.NewRepo(4, 4, 4, "4"),
				models.NewRepo(5, 5, 5, "5"),
				models.NewRepo(6, 6, 6, "6"),
			},
			5,
			[][]models.Artifact{
				{
					models.NewRepo(1, 1, 1, "1"),
					models.NewRepo(2, 2, 2, "2"),
					models.NewRepo(3, 3, 3, "3"),
					models.NewRepo(4, 4, 4, "4"),
					models.NewRepo(5, 5, 5, "5"),
				},
				{
					models.NewRepo(6, 6, 6, "6"),
				},
			},
		},
	}

	for i, c := range testCases {
		res := utils.SplitToBulks(c.InputSlice, c.InputN)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestSliceToMap(t *testing.T) {
	type TestCase struct {
		Input  []models.Artifact
		Output map[uint64]models.Artifact
	}

	var testCases = []TestCase{
		{
			[]models.Artifact{
				models.NewRepo(1, 1, 1, "1"),
				models.NewRepo(2, 2, 2, "2"),
				models.NewRepo(3, 3, 3, "3"),
				models.NewRepo(4, 4, 4, "4"),
				models.NewRepo(5, 5, 5, "5"),
				models.NewRepo(6, 6, 6, "6"),
			},
			map[uint64]models.Artifact{
				1: models.NewRepo(1, 1, 1, "1"),
				2: models.NewRepo(2, 2, 2, "2"),
				3: models.NewRepo(3, 3, 3, "3"),
				4: models.NewRepo(4, 4, 4, "4"),
				5: models.NewRepo(5, 5, 5, "5"),
				6: models.NewRepo(6, 6, 6, "6"),
			},
		},
		{nil, nil},
	}

	for i, c := range testCases {
		res := utils.SliceToMap(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}

	assert.Panics(t,
		func() {
			var failSlice = []models.Artifact{
				models.NewRepo(2, 1, 1, "1"),
				models.NewRepo(2, 2, 2, "2"),
			}
			utils.SliceToMap(failSlice)
		},
		"Fail panic test: The code did not panic")
}
