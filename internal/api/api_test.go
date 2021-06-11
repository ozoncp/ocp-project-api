package api_test

import (
	"context"
	"database/sql"
	"errors"
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
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})
	Context("create project simple", func() {
		var projectId uint64 = 1

		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			createRequest = &desc.CreateProjectRequest{CourseId: 1, Name: "1"}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(projectId)
			mock.ExpectQuery("INSERT INTO projects").
				WithArgs(createRequest.CourseId, createRequest.Name).
				WillReturnRows(rows)

			createResponse, err = grpcApi.CreateProject(ctx, createRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.ProjectId).Should(Equal(projectId))
		})
	})

	Context("create project: invalid argument", func() {
		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			createRequest = &desc.CreateProjectRequest{CourseId: 0, Name: "1"}

			createResponse, err = grpcApi.CreateProject(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("create project: sql query returns error", func() {
		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			createRequest = &desc.CreateProjectRequest{CourseId: 1, Name: "1"}

			mock.ExpectQuery("INSERT INTO projects").
				WithArgs(createRequest.CourseId, createRequest.Name).WillReturnError(errors.New("i am bad database"))

			createResponse, err = grpcApi.CreateProject(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

})
