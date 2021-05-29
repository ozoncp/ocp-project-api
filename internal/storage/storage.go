package storage

import "github.com/ozoncp/ocp-project-api/internal/models"

type Storage interface {
	AddObjects(task []models.Artifact) error
	RemoveObject(objectId uint64) error
	DescribeObject(objectId uint64) (*models.Artifact, error)
	ListObject(limit, offset uint64) ([]models.Artifact, error)
}
