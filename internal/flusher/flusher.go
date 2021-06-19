package flusher

import (
	"context"
	"log"

	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
)

type Flusher interface {
	FlushRepos(ctx context.Context, repos []models.Repo) []models.Repo
	FlushProjects(ctx context.Context, projects []models.Project) []models.Project
}

type flusher struct {
	repoStorage    storage.RepoStorage
	projectStorage storage.ProjectStorage
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	repoStorage storage.RepoStorage,
	projectStorage storage.ProjectStorage,
) Flusher {
	return &flusher{
		repoStorage:    repoStorage,
		projectStorage: projectStorage,
	}
}

// FlushRepos flush slice repos in storage
func (f *flusher) FlushRepos(ctx context.Context, repos []models.Repo) []models.Repo {
	indexes, err := f.repoStorage.MultiAddRepo(ctx, repos)
	if err != nil {
		log.Printf("Flushing warning: %v\n", err)
		return repos[len(indexes):]
	}

	return nil
}

// FlushProjects flush slice projects in storage
func (f *flusher) FlushProjects(ctx context.Context, projects []models.Project) []models.Project {
	indexes, err := f.projectStorage.MultiAddProject(ctx, projects)
	if err != nil {
		log.Printf("Flushing warning: %v\n", err)
		return projects[len(indexes):]
	}

	return nil
}
