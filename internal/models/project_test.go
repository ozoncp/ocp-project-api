package models_test

import (
	"github.com/ozoncp/ocp-project-api/internal/models"
	"testing"
)

func TestProjectInterface(t *testing.T) {
	var m = models.Project{Id: 1, CourseId: 2, Name: "a"}

	if len(m.String()) == 0 {
		t.Errorf("Project object function String() returns empty string")
	}
}
