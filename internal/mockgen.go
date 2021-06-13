package internal

//go:generate mockgen -destination=./mocks/flusher_mock.go -package=mocks github.com/ozoncp/ocp-project-api/internal/flusher Flusher
//go:generate mockgen -destination=./mocks/storage_mock.go -package=mocks github.com/ozoncp/ocp-project-api/internal/storage RepoStorage,ProjectStorage
//go:generate mockgen -destination=./mocks/saver_mock.go -package=mocks github.com/ozoncp/ocp-project-api/internal/saver Saver
//go:generate mockgen -destination=./mocks/alarm_mock.go -package=mocks github.com/ozoncp/ocp-project-api/internal/alarm Alarm
//go:generate mockgen -destination=./mocks/producer_mock.go -package=mocks github.com/ozoncp/ocp-project-api/internal/producer Producer
