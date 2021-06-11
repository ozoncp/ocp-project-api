package flusher

import (
	"context"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	"github.com/ozoncp/ocp-project-api/internal/utils"
	"log"
)

type Flusher interface {
	FlushRepos(ctx context.Context, repos []models.Repo) []models.Repo
	FlushProjects(ctx context.Context, projects []models.Project) []models.Project
}

type flusher struct {
	chunkSize      int
	repoStorage    storage.RepoStorage
	projectStorage storage.ProjectStorage
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	repoStorage storage.RepoStorage,
	projectStorage storage.ProjectStorage,
) Flusher {
	return &flusher{
		chunkSize:      chunkSize,
		repoStorage:    repoStorage,
		projectStorage: projectStorage,
	}
}

// FlushRepos flush slice repos in storage
func (f *flusher) FlushRepos(ctx context.Context, repos []models.Repo) []models.Repo {
	chunks, err := utils.ReposSplitToBulks(repos, f.chunkSize)
	if err != nil {
		log.Printf("Flushing warning: %v\n", err)
		return repos
	}

	for i := 0; i < len(chunks); i++ {
		if err := f.repoStorage.AddRepos(chunks[i]); err != nil {
			return repos[i*f.chunkSize:]
		}
	}

	return nil
}

// FlushProjects flush slice projects in storage
func (f *flusher) FlushProjects(ctx context.Context, projects []models.Project) []models.Project {
	chunks, err := utils.ProjectsSplitToBulks(projects, f.chunkSize)
	if err != nil {
		log.Printf("Flushing warning: %v\n", err)
		return projects
	}

	for i := 0; i < len(chunks); i++ {
		if _, err := f.projectStorage.MultiAddProject(ctx, chunks[i]); err != nil {
			return projects[i*f.chunkSize:]
		}
	}

	return nil
}
