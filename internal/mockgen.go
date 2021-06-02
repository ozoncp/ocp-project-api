package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-project-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/storage_mock.go -package=mocks github.com/ozoncp/ocp-project-api/internal/storage RepoStorage,ProjectStorage
