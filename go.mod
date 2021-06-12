module github.com/ozoncp/ocp-project-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/Masterminds/squirrel v1.5.0
	github.com/golang/mock v1.5.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.2.0
	github.com/onsi/ginkgo v1.16.2
	github.com/onsi/gomega v1.13.0
	github.com/ozoncp/ocp-project-api/pkg/ocp-project-api v0.0.0-00010101000000-000000000000
	github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api v0.0.0-00010101000000-000000000000
	github.com/rs/zerolog v1.22.0
	github.com/stretchr/testify v1.7.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.38.0
)

replace github.com/ozoncp/ocp-project-api/pkg/ocp-project-api => ./pkg/ocp-project-api

replace github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api => ./pkg/ocp-repo-api
