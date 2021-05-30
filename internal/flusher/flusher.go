package flusher

import (
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	"github.com/ozoncp/ocp-project-api/internal/utils"
	"log"
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
	chunks, err := utils.SplitToBulks(objects, f.chunkSize)
	if err != nil {
		log.Printf("Flushing warning: %v\n", err)
		return nil
	}

	for i := 0; i < len(chunks); i++ {
		if err := f.objectRepo.AddObjects(chunks[i]); err != nil {
			return objects[i*f.chunkSize:]
		}
	}

	return nil
}
