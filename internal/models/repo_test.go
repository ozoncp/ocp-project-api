package models_test

import (
	"github.com/ozoncp/ocp-project-api/internal/models"
	"testing"
)

func TestRepoInterface(t *testing.T) {
	var m = models.Repo{Id: 1, ProjectId: 2, UserId: 3, Link: "a"}

	if len(m.String()) == 0 {
		t.Errorf("Repo object function String() returns empty string")
	}
}
