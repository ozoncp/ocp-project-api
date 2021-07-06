package ocp_repo_api_test

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
	repoApi "github.com/ozoncp/ocp-project-api/internal/api/ocp-repo-api"
	"github.com/ozoncp/ocp-project-api/internal/mocks"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/storage"
	desc "github.com/ozoncp/ocp-project-api/pkg/ocp-repo-api"
)

var _ = Describe("Api", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		db     *sql.DB
		sqlxDB *sqlx.DB
		mock   sqlmock.Sqlmock

		repos = []models.Repo{
			{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			{Id: 2, ProjectId: 2, UserId: 2, Link: "2"},
		}

		logProducer *mocks.MockProducer
		repoStorage storage.RepoStorage
		grpcApi     desc.OcpRepoApiServer

		createRequest  *desc.CreateRepoRequest
		createResponse *desc.CreateRepoResponse

		updateRequest  *desc.UpdateRepoRequest
		updateResponse *desc.UpdateRepoResponse

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
		ctrl = gomock.NewController(GinkgoT())

		db, mock, err = sqlmock.New()
		Expect(err).Should(BeNil())

		sqlxDB = sqlx.NewDb(db, "sqlmock")

		repoStorage = storage.NewRepoStorage(sqlxDB, 2)
		logProducer = mocks.NewMockProducer(ctrl)
		grpcApi = repoApi.NewOcpRepoApi(repoStorage, logProducer)
	})

	JustBeforeEach(func() {
	})

	AfterEach(func() {
		mock.ExpectClose()
		err = db.Close()
		Expect(err).Should(BeNil())
		ctrl.Finish()
	})

	Context("version checking", func() {
		var versionResponse *desc.VersionResponse
		BeforeEach(func() {
			versionRequest := &desc.VersionRequest{}

			versionResponse, err = grpcApi.Version(ctx, versionRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(versionResponse.Version).Should(MatchRegexp("\\d+.\\d+.\\d+"))
		})
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

			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any())

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.RepoId).Should(Equal(repoId))
		})
	})

	Context("create repo: producer is not available", func() {
		BeforeEach(func() {
			createRequest = &desc.CreateRepoRequest{UserId: 1, ProjectId: 1, Link: "1"}

			logProducer.EXPECT().IsAvailable().Return(false)

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("create repo: producer error", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			createRequest = &desc.CreateRepoRequest{ProjectId: 1, UserId: 1, Link: "1"}

			rows := sqlmock.NewRows([]string{"id"}).
				AddRow(repoId)
			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(createRequest.ProjectId, createRequest.UserId, createRequest.Link).
				WillReturnRows(rows)

			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any()).Return(errors.New("i am bad producer"))

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(createResponse.RepoId).Should(Equal(repoId))
		})
	})

	Context("create repo: invalid argument", func() {
		BeforeEach(func() {
			createRequest = &desc.CreateRepoRequest{ProjectId: 0, UserId: 0, Link: "1"}

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("create repo: sql query returns error", func() {
		BeforeEach(func() {
			createRequest = &desc.CreateRepoRequest{ProjectId: 1, UserId: 1, Link: "1"}

			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(createRequest.ProjectId, createRequest.UserId, createRequest.Link).
				WillReturnError(errors.New("i am bad database"))
			logProducer.EXPECT().IsAvailable().Return(true)

			createResponse, err = grpcApi.CreateRepo(ctx, createRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("update repo simple", func() {
		BeforeEach(func() {
			updateRequest = &desc.UpdateRepoRequest{
				Repo: &desc.Repo{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mock.ExpectExec("UPDATE repos SET").
				WithArgs(
					updateRequest.Repo.ProjectId,
					updateRequest.Repo.UserId,
					updateRequest.Repo.Link,
					updateRequest.Repo.Id,
					false).
				WillReturnResult(sqlmock.NewResult(0, 1))

			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any())

			updateResponse, err = grpcApi.UpdateRepo(ctx, updateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(updateResponse.Found).Should(Equal(true))
		})
	})

	Context("update repo: producer error", func() {
		BeforeEach(func() {
			updateRequest = &desc.UpdateRepoRequest{
				Repo: &desc.Repo{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mock.ExpectExec("UPDATE repos SET").
				WithArgs(
					updateRequest.Repo.ProjectId,
					updateRequest.Repo.UserId,
					updateRequest.Repo.Link,
					updateRequest.Repo.Id,
					false).
				WillReturnResult(sqlmock.NewResult(0, 1))

			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any()).Return(errors.New("i am bad producer"))

			updateResponse, err = grpcApi.UpdateRepo(ctx, updateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(updateResponse.Found).Should(Equal(true))
		})
	})

	Context("update repo: not updated", func() {
		BeforeEach(func() {
			updateRequest = &desc.UpdateRepoRequest{
				Repo: &desc.Repo{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mock.ExpectExec("UPDATE repos SET").
				WithArgs(
					updateRequest.Repo.ProjectId,
					updateRequest.Repo.UserId,
					updateRequest.Repo.Link,
					updateRequest.Repo.Id,
					false).
				WillReturnResult(sqlmock.NewResult(0, 0))
			logProducer.EXPECT().IsAvailable().Return(true)

			updateResponse, err = grpcApi.UpdateRepo(ctx, updateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(updateResponse.Found).Should(Equal(false))
		})
	})

	Context("update repo: invalid argument", func() {
		BeforeEach(func() {
			updateRequest = &desc.UpdateRepoRequest{
				Repo: &desc.Repo{Id: 0, ProjectId: 0, UserId: 1, Link: "1"},
			}

			updateResponse, err = grpcApi.UpdateRepo(ctx, updateRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(updateResponse).Should(BeNil())
		})
	})

	Context("update repo: sql query returns error", func() {
		BeforeEach(func() {
			updateRequest = &desc.UpdateRepoRequest{
				Repo: &desc.Repo{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mock.ExpectExec("UPDATE repos SET").
				WithArgs(
					updateRequest.Repo.ProjectId,
					updateRequest.Repo.UserId,
					updateRequest.Repo.Link,
					updateRequest.Repo.Id,
					false).
				WillReturnError(errors.New("i am bad database"))
			logProducer.EXPECT().IsAvailable().Return(true)

			updateResponse, err = grpcApi.UpdateRepo(ctx, updateRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(updateResponse).Should(BeNil())
		})
	})

	Context("describe repo simple", func() {
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
				WithArgs(describeRequest.RepoId, false).
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

	Context("describe repo: not found", func() {

		BeforeEach(func() {
			describeRequest = &desc.DescribeRepoRequest{RepoId: 1}

			rows := sqlmock.NewRows([]string{"id", "course_id", "name"})
			mock.ExpectQuery("SELECT (.+) FROM repos WHERE").
				WithArgs(describeRequest.RepoId, false).
				WillReturnRows(rows)

			describeResponse, err = grpcApi.DescribeRepo(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(describeResponse).Should(BeNil())
		})
	})

	Context("describe repo: invalid argument", func() {
		BeforeEach(func() {
			describeRequest = &desc.DescribeRepoRequest{RepoId: 0}

			describeResponse, err = grpcApi.DescribeRepo(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("describe repo: sql query returns error", func() {
		BeforeEach(func() {
			describeRequest = &desc.DescribeRepoRequest{RepoId: 1}

			mock.ExpectQuery("SELECT (.+) FROM repos WHERE").
				WithArgs(describeRequest.RepoId, false).
				WillReturnError(errors.New("i am bad database"))

			describeResponse, err = grpcApi.DescribeRepo(ctx, describeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(createResponse).Should(BeNil())
		})
	})

	Context("remove repo simple", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}

			mock.ExpectExec("DELETE FROM repos").
				WithArgs(removeRequest.RepoId, false).WillReturnResult(sqlmock.NewResult(0, 1))

			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any())

			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResponse.Found).Should(Equal(true))
		})
	})

	Context("remove repo: producer error", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}

			mock.ExpectExec("DELETE FROM repos").
				WithArgs(removeRequest.RepoId, false).WillReturnResult(sqlmock.NewResult(0, 1))

			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any()).Return(errors.New("i am bad producer"))

			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResponse.Found).Should(Equal(true))
		})
	})

	Context("remove repo: not found", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}

			mock.ExpectExec("DELETE FROM repos").
				WithArgs(removeRequest.RepoId, false).WillReturnResult(sqlmock.NewResult(0, 0))
			logProducer.EXPECT().IsAvailable().Return(true)

			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(removeResponse.Found).Should(Equal(false))
		})
	})

	Context("remove repo: invalid argument", func() {
		var repoId uint64 = 0

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}
			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("remove repo: sql query returns error", func() {
		var repoId uint64 = 1

		BeforeEach(func() {
			removeRequest = &desc.RemoveRepoRequest{RepoId: repoId}

			mock.ExpectExec("DELETE FROM repos").
				WithArgs(removeRequest.RepoId, false).WillReturnError(errors.New("i am bad database"))
			logProducer.EXPECT().IsAvailable().Return(true)

			removeResponse, err = grpcApi.RemoveRepo(ctx, removeRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
		})
	})

	Context("list repo simple", func() {
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
				fmt.Sprintf("SELECT id, project_id, user_id, link FROM repos WHERE deleted = \\$1 LIMIT %d OFFSET %d", limit, offset)).
				WithArgs(false).
				WillReturnRows(rows)

			listResponse, err = grpcApi.ListRepos(ctx, listRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(len(listResponse.Repos)).Should(Equal(len(repos)))
		})
	})

	Context("list repo: sql query returns error", func() {
		var (
			limit  uint64 = 10
			offset uint64 = 0
		)

		BeforeEach(func() {
			listRequest = &desc.ListReposRequest{Limit: limit, Offset: offset}

			mock.ExpectQuery(
				fmt.Sprintf("SELECT id, project_id, user_id, link FROM repos WHERE deleted = \\$1 LIMIT %d OFFSET %d", limit, offset)).
				WithArgs(false).
				WillReturnError(errors.New("i am bad database"))

			listResponse, err = grpcApi.ListRepos(ctx, listRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(listResponse).Should(BeNil())
		})
	})

	Context("multi create repo simple", func() {
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

			rows := sqlmock.NewRows([]string{"id"}).AddRow(1).AddRow(2)
			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(
					repos[0].ProjectId, repos[0].UserId, repos[0].Link,
					repos[1].ProjectId, repos[1].UserId, repos[1].Link).
				WillReturnRows(rows)
			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any())

			multiCreateResponse, err = grpcApi.MultiCreateRepo(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateResponse.CountOfCreated).Should(Equal(int64(len(multiCreateRequest.Repos))))
		})
	})

	Context("multi create repo: invalid argument", func() {
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

	Context("multi create repo: sql query returns error", func() {
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

			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg()).
				WillReturnError(errors.New("i am bad database"))
			logProducer.EXPECT().IsAvailable().Return(true)

			multiCreateResponse, err = grpcApi.MultiCreateRepo(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).ShouldNot(BeNil())
			Expect(multiCreateResponse).Should(BeNil())
		})
	})

	Context("multi create repo: split to bulks", func() {
		BeforeEach(func() {
			repoStorage = storage.NewRepoStorage(sqlxDB, 1)
			grpcApi = repoApi.NewOcpRepoApi(repoStorage, logProducer)

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

			rows := sqlmock.NewRows([]string{"id"}).AddRow(1)
			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(repos[0].ProjectId, repos[0].UserId, repos[0].Link).
				WillReturnRows(rows)

			rows = sqlmock.NewRows([]string{"id"}).AddRow(2)
			mock.ExpectQuery("INSERT INTO repos").
				WithArgs(repos[1].ProjectId, repos[1].UserId, repos[1].Link).
				WillReturnRows(rows)
			logProducer.EXPECT().IsAvailable().Return(true)
			logProducer.EXPECT().SendMessage(gomock.Any())

			multiCreateResponse, err = grpcApi.MultiCreateRepo(ctx, multiCreateRequest)
		})

		It("", func() {
			Expect(err).Should(BeNil())
			Expect(multiCreateResponse.CountOfCreated).Should(Equal(int64(len(multiCreateRequest.Repos))))
		})
	})

})
