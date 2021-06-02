package storage

import "github.com/ozoncp/ocp-project-api/internal/models"

type RepoStorage interface {
	AddRepos(task []models.Repo) error
	RemoveRepo(repoId uint64) error
	DescribeRepo(repoId uint64) (*models.Repo, error)
	ListRepos(limit, offset uint64) ([]models.Repo, error)
}
