package ocp_repo_api_test

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	projectApi "github.com/ozoncp/ocp-project-api/internal/api/ocp-repo-api"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api"
)

var _ = Describe("Api", func() {
	var (
		ctx context.Context

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		repos = []models.Repo{
			{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
		}

		repoStorage storage.RepoStorage
		grpcApi     desc.OcpRepoApiServer

		createRequest  *desc.CreateRepoRequest
		createResponse *desc.CreateRepoResponse

		describeRequest  *desc.DescribeRepoRequest
		describeResponse *desc.DescribeRepoResponse

		removeRequest  *desc.RemoveRepoRequest
		removeResponse *desc.RemoveRepoResponse

		listRequest  *desc.ListReposRequest
		listResponse *desc.ListReposResponse

		multiCreateRequest  *desc.MultiCreateRepoRequest
		multiCreateResponse *desc.MultiCreateRepoResponse

		err error
	)
	BeforeEach(func() {
		ctx = context.Background()

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")

		repoStorage = storage.NewRepoStorage(sqlxDB, 2)
		grpcApi = projectApi.NewOcpRepoApi(repoStorage)
	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
	})

	Context("create repo simple", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			createRequest = &desc.CreateRepoRequest{ProjectId: 1, UserId: 1, Link: "1"}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(repoId)
			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(createRequest.ProjectId, createRequest.UserId, createRequest.Link).
				WillReturnRows(rows)

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.RepoId).Should(Equal(repoId))
		})
	})

	Context("create project: invalid argument", func() {
		BeforeEach(func() {
			createRequest = &desc.CreateRepoRequest{ProjectId: 0, UserId: 0, Link: "1"}

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("create project: sql query returns error", func() {
		BeforeEach(func() {
			createRequest = &desc.CreateRepoRequest{ProjectId: 1, UserId: 1, Link: "1"}

			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(createRequest.ProjectId, createRequest.UserId, createRequest.Link).
				WillReturnError(errors.New("i am bad database"))

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("describe project simple", func() {
		var (
			repoId    uint64 = 1
			projectId uint64 = 1
			userId    uint64 = 1
			link             = "1"
		)

		BeforeEach(func() {
			describeRequest = &desc.DescribeRepoRequest{RepoId: repoId}

			rows := sqlmock.NewRows([]string{"id", "project_id", "user_id", "link"}).
				AddRow(repoId, projectId, userId, link)
			mock.ExpectQuery("SELECT (.+) FROM repos WHERE").
				WithArgs(describeRequest.RepoId).
				WillReturnRows(rows)

			describeResponse, err = grpcApi.DescribeRepo(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(describeResponse.Repo.Id).Should(Equal(repoId))
			Expect(describeResponse.Repo.ProjectId).Should(Equal(projectId))
			Expect(describeResponse.Repo.UserId).Should(Equal(userId))
			Expect(describeResponse.Repo.Link).Should(Equal(link))
		})
	})

	Context("describe project: invalid argument", func() {
		BeforeEach(func() {
			describeRequest = &desc.DescribeRepoRequest{RepoId: 0}

			describeResponse, err = grpcApi.DescribeRepo(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("describe project: sql query returns error", func() {
		BeforeEach(func() {
			describeRequest = &desc.DescribeRepoRequest{RepoId: 1}

			mock.ExpectQuery("SELECT (.+) FROM repos WHERE").
				WithArgs(describeRequest.RepoId).
				WillReturnError(errors.New("i am bad database"))

			describeResponse, err = grpcApi.DescribeRepo(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("remove project simple", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}

			mock.ExpectExec("DELETE FROM repos").
				WithArgs(removeRequest.RepoId).WillReturnResult(sqlmock.NewResult(0, 1))

			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResponse.Found).Should(Equal(true))
		})
	})
	Context("remove project: not found", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}

			mock.ExpectExec("DELETE FROM repos").
				WithArgs(removeRequest.RepoId).WillReturnResult(sqlmock.NewResult(0, 0))

			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResponse.Found).Should(Equal(false))
		})
	})

	Context("remove project: invalid argument", func() {
		var repoId uint64 = 0

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}
			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("remove project: sql query returns error", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}

			mock.ExpectExec("DELETE FROM repos").
				WithArgs(removeRequest.RepoId).WillReturnError(errors.New("i am bad database"))

			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
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
			listRequest = &desc.ListReposRequest{Limit: limit, Offset: offset}

			rows := sqlmock.NewRows([]string{"id", "project_id", "user_id", "link"}).
				AddRow(repos[0].Id, repos[0].ProjectId, repos[0].UserId, repos[0].Link).
				AddRow(repos[1].Id, repos[1].ProjectId, repos[1].UserId, repos[1].Link)

			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, project_id, user_id, link FROM repos LIMIT %d OFFSET %d", limit, offset)).
				WillReturnRows(rows)

			listResponse, err = grpcApi.ListRepos(ctx, listRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(len(listResponse.Repos)).Should(Equal(len(repos)))
		})
	})

	Context("list project: sql query returns error", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0
		)

		BeforeEach(func() {
			listRequest = &desc.ListReposRequest{Limit: limit, Offset: offset}

			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, project_id, user_id, link FROM repos LIMIT %d OFFSET %d", limit, offset)).
				WillReturnError(errors.New("i am bad database"))

			listResponse, err = grpcApi.ListRepos(ctx, listRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(listResponse).Should(BeNil())
		})
	})

	Context("multi create project simple", func() {
		BeforeEach(func() {
			multiCreateRequest = &desc.MultiCreateRepoRequest{
				Repos: []*desc.NewRepo{
					{
						ProjectId: repos[0].ProjectId,
						UserId:    repos[0].UserId,
						Link:      repos[0].Link,
					},
					{
						ProjectId: repos[1].ProjectId,
						UserId:    repos[1].UserId,
						Link:      repos[1].Link,
					},
				},
			}

			mock.ExpectExec("INSERT INTO repos").
				WithArgs(
					repos[0].ProjectId, repos[0].UserId, repos[0].Link,
					repos[1].ProjectId, repos[1].UserId, repos[1].Link).
				WillReturnResult(sqlmock.NewResult(2, 2))

			multiCreateResponse, err = grpcApi.MultiCreateRepo(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateResponse.CountOfCreated).Should(Equal(int64(len(multiCreateRequest.Repos))))
		})
	})

	Context("multi create project: invalid argument", func() {
		BeforeEach(func() {
			multiCreateRequest = &desc.MultiCreateRepoRequest{
				Repos: []*desc.NewRepo{
					{
						ProjectId: 0,
						UserId:    0,
						Link:      "",
					},
				},
			}

			multiCreateResponse, err = grpcApi.MultiCreateRepo(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateResponse).Should(BeNil())
		})
	})

	Context("multi create project: sql query returns error", func() {
		BeforeEach(func() {
			multiCreateRequest = &desc.MultiCreateRepoRequest{
				Repos: []*desc.NewRepo{
					{
						ProjectId: 1,
						UserId:    1,
						Link:      "1",
					},
				},
			}

			mock.ExpectExec("INSERT INTO repos").
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("i am bad database"))

			multiCreateResponse, err = grpcApi.MultiCreateRepo(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateResponse).Should(BeNil())
		})
	})

	Context("multi create project: split to bulks", func() {
		BeforeEach(func() {
			repoStorage = storage.NewRepoStorage(sqlxDB, 1)
			grpcApi = projectApi.NewOcpRepoApi(repoStorage)

			multiCreateRequest = &desc.MultiCreateRepoRequest{
				Repos: []*desc.NewRepo{
					{
						ProjectId: repos[0].ProjectId,
						UserId:    repos[0].UserId,
						Link:      repos[0].Link,
					},
					{
						ProjectId: repos[1].ProjectId,
						UserId:    repos[1].UserId,
						Link:      repos[1].Link,
					},
				},
			}

			mock.ExpectExec("INSERT INTO repos").
				WithArgs(repos[0].ProjectId, repos[0].UserId, repos[0].Link).
				WillReturnResult(sqlmock.NewResult(1, 1))

			mock.ExpectExec("INSERT INTO repos").
				WithArgs(repos[1].ProjectId, repos[1].UserId, repos[1].Link).
				WillReturnResult(sqlmock.NewResult(1, 1))

			multiCreateResponse, err = grpcApi.MultiCreateRepo(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateResponse.CountOfCreated).Should(Equal(int64(len(multiCreateRequest.Repos))))
		})
	})

})
