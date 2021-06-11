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

		describeRequest  *desc.DescribeProjectRequest
		describeResponse *desc.DescribeProjectResponse

		removeRequest  *desc.RemoveProjectRequest
		removeResponse *desc.RemoveProjectResponse

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
				WithArgs(createRequest.CourseId, createRequest.Name).
				WillReturnError(errors.New("i am bad database"))

			createResponse, err = grpcApi.CreateProject(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("describe project simple", func() {
		var (
			projectId uint64 = 1
			courseId  uint64 = 1
			name      string = "1"
		)

		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			describeRequest = &desc.DescribeProjectRequest{ProjectId: 1}

			rows := sqlmock.NewRows([]string{"id", "course_id", "name"}).
				AddRow(projectId, courseId, name)
			mock.ExpectQuery("SELECT (.+) FROM projects WHERE").
				WithArgs(describeRequest.ProjectId).
				WillReturnRows(rows)

			describeResponse, err = grpcApi.DescribeProject(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(describeResponse.Project.Id).Should(Equal(projectId))
			Expect(describeResponse.Project.CourseId).Should(Equal(courseId))
			Expect(describeResponse.Project.Name).Should(Equal(name))
		})
	})

	Context("describe project: invalid argument", func() {
		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			describeRequest = &desc.DescribeProjectRequest{ProjectId: 0}

			describeResponse, err = grpcApi.DescribeProject(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("describe project: sql query returns error", func() {
		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			describeRequest = &desc.DescribeProjectRequest{ProjectId: 1}

			mock.ExpectQuery("SELECT (.+) FROM projects WHERE").
				WithArgs(describeRequest.ProjectId).
				WillReturnError(errors.New("i am bad database"))

			describeResponse, err = grpcApi.DescribeProject(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("remove project simple", func() {
		var projectId uint64 = 1

		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			removeRequest = &desc.RemoveProjectRequest{ProjectId: projectId}

			mock.ExpectExec("DELETE FROM projects").
				WithArgs(removeRequest.ProjectId).WillReturnResult(sqlmock.NewResult(0, 1))

			removeResponse, err = grpcApi.RemoveProject(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResponse.Found).Should(Equal(true))
		})
	})
	Context("remove project: not found", func() {
		var projectId uint64 = 1

		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			removeRequest = &desc.RemoveProjectRequest{ProjectId: projectId}

			mock.ExpectExec("DELETE FROM projects").
				WithArgs(removeRequest.ProjectId).WillReturnResult(sqlmock.NewResult(0, 0))

			removeResponse, err = grpcApi.RemoveProject(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResponse.Found).Should(Equal(false))
		})
	})

	Context("remove project: invalid argument", func() {
		var projectId uint64 = 0

		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			removeRequest = &desc.RemoveProjectRequest{ProjectId: projectId}

			removeResponse, err = grpcApi.RemoveProject(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("remove project: sql query returns error", func() {
		var projectId uint64 = 1

		BeforeEach(func() {
			chunkSize = 1
			projectStorage = storage.NewProjectStorage(sqlxDB, chunkSize)
			grpcApi = api.NewOcpProjectApi(projectStorage)

			removeRequest = &desc.RemoveProjectRequest{ProjectId: projectId}

			mock.ExpectExec("DELETE FROM projects").
				WithArgs(removeRequest.ProjectId).WillReturnError(errors.New("i am bad database"))

			removeResponse, err = grpcApi.RemoveProject(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

})
