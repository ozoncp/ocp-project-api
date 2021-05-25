package models_test

import (
	"github.com/ozoncp/ocp-project-api/internal/models"
	"testing"
)

func TestProjectInterface(t *testing.T) {
	var m models.Artifact = models.NewProject().Init(1, 2, "a")

	if m, ok := m.(*models.Project); ok {
		if m.Id() != 1 {
			t.Errorf("Project object has not that value of id")
			return
		}
		if m.String() == "" {
			t.Errorf("Project object returns invalid output by String()")
			return
		}
		if m.CourseId != 2 {
			t.Errorf("Project object has not that value of CourseId")
			return
		}
		if m.Name != "a" {
			t.Errorf("Project object has not that value of Name")
			return
		}
	} else {
		t.Errorf("Failed Project type declaration")
		return
	}
}

func TestProjectCreation(t *testing.T) {
	var m = models.NewProject()
	if m == nil {
		t.Errorf("Failed creation Project object")
		return
	}
	m.Init(1, 2, "a")
	if m.Id() != 1 {
		t.Errorf("Project object has not that value of id")
		return
	}
	if m.CourseId != 2 {
		t.Errorf("Project object has not that value of CourseId")
		return
	}
	if m.Name != "a" {
		t.Errorf("Project object has not that value of Name")
		return
	}
}
