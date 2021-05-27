package models_test

import (
	"github.com/ozoncp/ocp-project-api/internal/models"
	"testing"
)

func TestRepoInterface(t *testing.T) {
	var m models.Artifact = models.NewRepo().Init(1, 2, 3, "a")

	if m, ok := m.(*models.Repo); ok {
		if m.Id() != 1 {
			t.Errorf("Repo object has not that value of id")
			return
		}
		if m.String() == "" {
			t.Errorf("Repo object returns invalid output by String()")
			return
		}
		if m.ProjectId != 2 {
			t.Errorf("Repo object has not that value of ProjectId")
		}
		if m.UserId != 3 {
			t.Errorf("Repo object has not that value of UserId")
			return
		}
		if m.Link != "a" {
			t.Errorf("Repo object has not that value of Link")
			return
		}
	} else {
		t.Errorf("Failed Repo type declaration")
		return
	}
}

func TestRepoCreation(t *testing.T) {
	var m = models.NewRepo()
	if m == nil {
		t.Errorf("Failed creation Repo object")
		return
	}
	m.Init(1, 2, 3, "a")
	if m.Id() != 1 {
		t.Errorf("Repo object has not that value of id")
		return
	}
	if m.ProjectId != 2 {
		t.Errorf("Repo object has not that value of ProjectId")
	}
	if m.UserId != 3 {
		t.Errorf("Repo object has not that value of UserId")
		return
	}
	if m.Link != "a" {
		t.Errorf("Repo object has not that value of Link")
		return
	}
}
