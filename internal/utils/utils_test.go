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

func TestProjectsSplitToBulks(t *testing.T) {
	type TestCase struct {
		InputSlice []models.Project
		InputN     int
		Output     [][]models.Project
	}

	testCases := []TestCase{
		{
			[]models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 2, CourseId: 2, Name: "1"},
				{Id: 3, CourseId: 3, Name: "1"},
				{Id: 4, CourseId: 4, Name: "1"},
				{Id: 5, CourseId: 5, Name: "1"},
				{Id: 6, CourseId: 6, Name: "1"},
			},
			2,
			[][]models.Project{
				{
					{Id: 1, CourseId: 1, Name: "1"},
					{Id: 2, CourseId: 2, Name: "1"},
				},
				{
					{Id: 3, CourseId: 3, Name: "1"},
					{Id: 4, CourseId: 4, Name: "1"},
				},
				{
					{Id: 5, CourseId: 5, Name: "1"},
					{Id: 6, CourseId: 6, Name: "1"},
				},
			},
		},
		{
			[]models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 2, CourseId: 2, Name: "1"},
				{Id: 3, CourseId: 3, Name: "1"},
				{Id: 4, CourseId: 4, Name: "1"},
				{Id: 5, CourseId: 5, Name: "1"},
				{Id: 6, CourseId: 6, Name: "1"},
			},
			4,
			[][]models.Project{
				{
					{Id: 1, CourseId: 1, Name: "1"},
					{Id: 2, CourseId: 2, Name: "1"},
					{Id: 3, CourseId: 3, Name: "1"},
					{Id: 4, CourseId: 4, Name: "1"},
				},
				{
					{Id: 5, CourseId: 5, Name: "1"},
					{Id: 6, CourseId: 6, Name: "1"},
				},
			},
		},

		{[]models.Project{}, 2, [][]models.Project{{}}},
		{
			[]models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 2, CourseId: 2, Name: "1"},
				{Id: 3, CourseId: 3, Name: "1"},
			},
			0,
			nil,
		},
		{nil, 2, nil},
	}

	for i, c := range testCases {
		res, _ := utils.ProjectsSplitToBulks(c.InputSlice, c.InputN)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestReposSplitToBulks(t *testing.T) {
	type TestCase struct {
		InputSlice []models.Repo
		InputN     int
		Output     [][]models.Repo
	}

	testCases := []TestCase{
		{
			[]models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
				{Id: 3, ProjectId: 3, UserId: 3, Link: "3"},
				{Id: 4, ProjectId: 4, UserId: 4, Link: "4"},
				{Id: 5, ProjectId: 5, UserId: 5, Link: "5"},
				{Id: 6, ProjectId: 6, UserId: 6, Link: "6"},
			},
			2,
			[][]models.Repo{
				{
					{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
					{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
				},
				{
					{Id: 3, ProjectId: 3, UserId: 3, Link: "3"},
					{Id: 4, ProjectId: 4, UserId: 4, Link: "4"},
				},
				{
					{Id: 5, ProjectId: 5, UserId: 5, Link: "5"},
					{Id: 6, ProjectId: 6, UserId: 6, Link: "6"},
				},
			},
		},
		{
			[]models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
				{Id: 3, ProjectId: 3, UserId: 3, Link: "3"},
				{Id: 4, ProjectId: 4, UserId: 4, Link: "4"},
				{Id: 5, ProjectId: 5, UserId: 5, Link: "5"},
				{Id: 6, ProjectId: 6, UserId: 6, Link: "6"},
			},
			4,
			[][]models.Repo{
				{
					{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
					{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
					{Id: 3, ProjectId: 3, UserId: 3, Link: "3"},
					{Id: 4, ProjectId: 4, UserId: 4, Link: "4"},
				},
				{
					{Id: 5, ProjectId: 5, UserId: 5, Link: "5"},
					{Id: 6, ProjectId: 6, UserId: 6, Link: "6"},
				},
			},
		},

		{[]models.Repo{}, 2, [][]models.Repo{{}}},
		{
			[]models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
				{Id: 3, ProjectId: 3, UserId: 3, Link: "3"},
			},
			0,
			nil,
		},
		{nil, 2, nil},
	}

	for i, c := range testCases {
		res, _ := utils.ReposSplitToBulks(c.InputSlice, c.InputN)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}
}

func TestReposSliceToMap(t *testing.T) {
	type TestCase struct {
		Input  []models.Repo
		Output map[uint64]models.Repo
	}

	var testCases = []TestCase{
		{
			[]models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
				{Id: 3, ProjectId: 3, UserId: 3, Link: "3"},
				{Id: 4, ProjectId: 4, UserId: 4, Link: "4"},
				{Id: 5, ProjectId: 5, UserId: 5, Link: "5"},
				{Id: 6, ProjectId: 6, UserId: 6, Link: "6"},
			},
			map[uint64]models.Repo{
				1: {Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				2: {Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
				3: {Id: 3, ProjectId: 3, UserId: 3, Link: "3"},
				4: {Id: 4, ProjectId: 4, UserId: 4, Link: "4"},
				5: {Id: 5, ProjectId: 5, UserId: 5, Link: "5"},
				6: {Id: 6, ProjectId: 6, UserId: 6, Link: "6"},
			},
		},
		{nil, nil},
	}

	for i, c := range testCases {
		res := utils.ReposSliceToMap(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}

	assert.Panics(t,
		func() {
			var failSlice = []models.Repo{
				{Id: 2, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
			}
			utils.ReposSliceToMap(failSlice)
		},
		"Fail panic test: The code did not panic")
}

func TestProjectsSliceToMap(t *testing.T) {
	type TestCase struct {
		Input  []models.Project
		Output map[uint64]models.Project
	}

	var testCases = []TestCase{
		{
			[]models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 2, CourseId: 2, Name: "1"},
				{Id: 3, CourseId: 3, Name: "1"},
				{Id: 4, CourseId: 4, Name: "1"},
				{Id: 5, CourseId: 5, Name: "1"},
				{Id: 6, CourseId: 6, Name: "1"},
			},
			map[uint64]models.Project{
				1: {Id: 1, CourseId: 1, Name: "1"},
				2: {Id: 2, CourseId: 2, Name: "1"},
				3: {Id: 3, CourseId: 3, Name: "1"},
				4: {Id: 4, CourseId: 4, Name: "1"},
				5: {Id: 5, CourseId: 5, Name: "1"},
				6: {Id: 6, CourseId: 6, Name: "1"},
			},
		},
		{nil, nil},
	}

	for i, c := range testCases {
		res := utils.ProjectsSliceToMap(c.Input)
		if !reflect.DeepEqual(res, c.Output) {
			fmt.Println("Fail result: ", res)
			t.Errorf("Fail test case %d\n", i+1)
			return
		}
		fmt.Println("Good result: ", res)
	}

	assert.Panics(t,
		func() {
			var failSlice = []models.Project{
				{Id: 2, CourseId: 1, Name: "1"},
				{Id: 2, CourseId: 2, Name: "1"},
			}
			utils.ProjectsSliceToMap(failSlice)
		},
		"Fail panic test: The code did not panic")
}
