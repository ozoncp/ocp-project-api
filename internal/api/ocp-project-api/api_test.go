package ocp_project_api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/golang/mock/gomock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	projectApi "github.com/ozoncp/ocp-project-api/internal/api/ocp-project-api"
	"github.com/ozoncp/ocp-project-api/internal/mocks"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-project-api"
)

var _ = Describe("Api", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		projects = []models.Project{
			{Id: 1, CourseId: 1, Name: "1"},
			{Id: 2, CourseId: 2, Name: "2"},
		}

		logProducer    *mocks.MockProducer
		projectStorage storage.ProjectStorage
		grpcApi        desc.OcpProjectApiServer

		createRequest  *desc.CreateProjectRequest
		createResponse *desc.CreateProjectResponse

		describeRequest  *desc.DescribeProjectRequest
		describeResponse *desc.DescribeProjectResponse

		removeRequest  *desc.RemoveProjectRequest
		removeResponse *desc.RemoveProjectResponse

		listRequest  *desc.ListProjectsRequest
		listResponse *desc.ListProjectsResponse

		multiCreateRequest  *desc.MultiCreateProjectRequest
		multiCreateResponse *desc.MultiCreateProjectResponse

		err error
	)
	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")

		projectStorage = storage.NewProjectStorage(sqlxDB, 2)

		logProducer = mocks.NewMockProducer(ctrl)
		grpcApi = projectApi.NewOcpProjectApi(projectStorage, logProducer)
	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
		ctrl.Finish()
	})

	Context("create project simple", func() {
		var projectId uint64 = 1

		BeforeEach(func() {
			createRequest = &desc.CreateProjectRequest{CourseId: 1, Name: "1"}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(projectId)
			mock.ExpectQuery("INSERT INTO projects").
				WithArgs(createRequest.CourseId, createRequest.Name).
				WillReturnRows(rows)

			logProducer.EXPECT().SendMessage(gomock.Any())

			createResponse, err = grpcApi.CreateProject(ctx, createRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.ProjectId).Should(Equal(projectId))
		})
	})

	Context("create project: invalid argument", func() {
		BeforeEach(func() {
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
			name             = "1"
		)

		BeforeEach(func() {
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
			describeRequest = &desc.DescribeProjectRequest{ProjectId: 0}

			describeResponse, err = grpcApi.DescribeProject(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("describe project: sql query returns error", func() {
		BeforeEach(func() {
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
			removeRequest = &desc.RemoveProjectRequest{ProjectId: projectId}

			mock.ExpectExec("DELETE FROM projects").
				WithArgs(removeRequest.ProjectId).WillReturnResult(sqlmock.NewResult(0, 1))

			logProducer.EXPECT().SendMessage(gomock.Any())

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
			removeRequest = &desc.RemoveProjectRequest{ProjectId: projectId}

			mock.ExpectExec("DELETE FROM projects").
				WithArgs(removeRequest.ProjectId).WillReturnError(errors.New("i am bad database"))

			removeResponse, err = grpcApi.RemoveProject(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("list project simple", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0
		)

		BeforeEach(func() {
			listRequest = &desc.ListProjectsRequest{Limit: limit, Offset: offset}

			rows := sqlmock.NewRows([]string{"id", "course_id", "name"}).
				AddRow(projects[0].Id, projects[0].CourseId, projects[0].Name).
				AddRow(projects[1].Id, projects[1].CourseId, projects[1].Name)

			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, course_id, name FROM projects LIMIT %d OFFSET %d", limit, offset)).
				WillReturnRows(rows)

			listResponse, err = grpcApi.ListProjects(ctx, listRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(len(listResponse.Projects)).Should(Equal(len(projects)))
		})
	})

	Context("list project: sql query returns error", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0
		)

		BeforeEach(func() {
			listRequest = &desc.ListProjectsRequest{Limit: limit, Offset: offset}

			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, course_id, name FROM projects LIMIT %d OFFSET %d", limit, offset)).
				WillReturnError(errors.New("i am bad database"))

			listResponse, err = grpcApi.ListProjects(ctx, listRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(listResponse).Should(BeNil())
		})
	})

	Context("multi create project simple", func() {
		BeforeEach(func() {
			multiCreateRequest = &desc.MultiCreateProjectRequest{
				Projects: []*desc.NewProject{
					{
						CourseId: projects[0].CourseId,
						Name:     projects[0].Name,
					},
					{
						CourseId: projects[1].CourseId,
						Name:     projects[1].Name,
					},
				},
			}

			mock.ExpectExec("INSERT INTO projects").
				WithArgs(projects[0].CourseId, projects[0].Name, projects[1].CourseId, projects[1].Name).
				WillReturnResult(sqlmock.NewResult(2, 2))

			multiCreateResponse, err = grpcApi.MultiCreateProject(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateResponse.CountOfCreated).Should(Equal(int64(len(multiCreateRequest.Projects))))
		})
	})

	Context("multi create project: invalid argument", func() {
		BeforeEach(func() {
			multiCreateRequest = &desc.MultiCreateProjectRequest{
				Projects: []*desc.NewProject{
					{
						CourseId: 0,
						Name:     "",
					},
				},
			}

			multiCreateResponse, err = grpcApi.MultiCreateProject(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateResponse).Should(BeNil())
		})
	})

	Context("multi create project: sql query returns error", func() {
		BeforeEach(func() {
			multiCreateRequest = &desc.MultiCreateProjectRequest{
				Projects: []*desc.NewProject{
					{
						CourseId: 1,
						Name:     "1",
					},
				},
			}

			mock.ExpectExec("INSERT INTO projects").
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("i am bad database"))

			multiCreateResponse, err = grpcApi.MultiCreateProject(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateResponse).Should(BeNil())
		})
	})

	Context("multi create project: split to bulks", func() {
		BeforeEach(func() {
			projectStorage = storage.NewProjectStorage(sqlxDB, 1)
			logProducer = mocks.NewMockProducer(ctrl)
			grpcApi = projectApi.NewOcpProjectApi(projectStorage, logProducer)

			multiCreateRequest = &desc.MultiCreateProjectRequest{
				Projects: []*desc.NewProject{
					{
						CourseId: projects[0].CourseId,
						Name:     projects[0].Name,
					},
					{
						CourseId: projects[1].CourseId,
						Name:     projects[1].Name,
					},
				},
			}

			mock.ExpectExec("INSERT INTO projects").
				WithArgs(projects[0].CourseId, projects[0].Name).
				WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectExec("INSERT INTO projects").
				WithArgs(projects[1].CourseId, projects[1].Name).
				WillReturnResult(sqlmock.NewResult(1, 1))

			multiCreateResponse, err = grpcApi.MultiCreateProject(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateResponse.CountOfCreated).Should(Equal(int64(len(multiCreateRequest.Projects))))
		})
	})

})
