package storage

import "github.com/ozoncp/ocp-project-api/internal/models"

type ProjectStorage interface {
	AddProjects(task []models.Project) error
	RemoveProject(projectId uint64) error
	DescribeProject(projectId uint64) (*models.Project, error)
	ListProjects(limit, offset uint64) ([]models.Project, error)
}
