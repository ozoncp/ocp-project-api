package api_test

import (
	"context"
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-project-api/internal/api"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"
)

var _ = Describe("Api", func() {
	var (
		ctx context.Context

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		projectStorage storage.ProjectStorage
		grpcApi        desc.OcpProjectApiServer

		createRequest  *desc.CreateProjectRequest
		createResponse *desc.CreateProjectResponse

		err error

		chunkSize int
	)
	BeforeEach(func() {
		ctx = context.Background()

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")
	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
		err = db.Close()
	})
	Context("create project simple", func() {
		var projectId uint64 = 1

		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(projectId)
			mock.ExpectQuery("INSERT INTO projects").WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).WillReturnRows(rows)

			createRequest = &desc.CreateProjectRequest{CourseId: 1, Name: "1"}
			createResponse, err = grpcApi.CreateProject(ctx, createRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.ProjectId).Should(Equal(projectId))
		})
	})
})
