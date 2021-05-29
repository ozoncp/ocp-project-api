package flusher

import (
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
)

type Flusher interface {
	Flush(objects []models.Artifact) []models.Artifact
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(
	chunkSize int,
	objectRepo storage.Storage,
) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		objectRepo: objectRepo,
	}
}

type flusher struct {
	chunkSize  int
	objectRepo storage.Storage
}

func (f *flusher) Flush(objects []models.Artifact) []models.Artifact {

	for i := 0; i < len(objects); i += f.chunkSize {
		j := i + f.chunkSize
		if j >= len(objects) {
			j = len(objects)
		}

		chunk := objects[i:j]
		if err := f.objectRepo.AddObjects(chunk); err != nil {
			return objects[i:]
		}
	}
	return nil
}
