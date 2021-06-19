package flusher_test

import (
	"context"
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-project-api/internal/flusher"
	"github.com/ozoncp/ocp-project-api/internal/mocks"
	"github.com/ozoncp/ocp-project-api/internal/models"
)

var _ = Describe("Flush into RepoStorage", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		mockRepoStorage *mocks.MockRepoStorage

		repos []models.Repo
		rest  []models.Repo

		f flusher.Flusher

		chunkSize int
	)
	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockRepoStorage = mocks.NewMockRepoStorage(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepoStorage, &mocks.MockProjectStorage{})
		rest = f.FlushRepos(ctx, repos)
	})

	AfterEach(func() {
		ctrl.Finish()
	})
	Context("checking incorrect input", func() {

		BeforeEach(func() {
			chunkSize = -1
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Any()).Return([]uint64{}, nil).Times(0)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(repos))
		})
	})

	Context("storage saves all objects", func() {

		BeforeEach(func() {
			chunkSize = 2
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Len(chunkSize)).Return([]uint64{1, 2}, nil).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("storage doesn't save all objects", func() {

		BeforeEach(func() {
			chunkSize = 2
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Len(len(repos))).Return([]uint64{}, errors.New("some error")).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(repos))
		})
	})

	Context("storage doesn't save part of objects", func() {

		BeforeEach(func() {
			chunkSize = 3
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 2, ProjectId: 2, UserId: 2, Link: "1"},
				{Id: 3, ProjectId: 3, UserId: 3, Link: "1"},
			}

			gomock.InOrder(
				mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Len(chunkSize)).Return([]uint64{1, 2, 3}, nil).Times(1),
				mockRepoStorage.EXPECT().MultiAddRepo(
					ctx,
					gomock.Len(len(repos)-chunkSize)).Return([]uint64{}, errors.New("some error")).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo([]models.Repo{
				{Id: 2, ProjectId: 2, UserId: 2, Link: "1"},
				{Id: 3, ProjectId: 3, UserId: 3, Link: "1"},
			}))
		})
	})

	Context("checking storage calls order", func() {

		BeforeEach(func() {
			chunkSize = 3
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			gomock.InOrder(
				mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Len(chunkSize)).Return([]uint64{1, 2, 3}, nil).Times(1),
				mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Len(len(repos)-chunkSize)).Return([]uint64{4}, nil).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})
})

var _ = Describe("Flush into ProjectStorage", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		mockProjectStorage *mocks.MockProjectStorage

		projects []models.Project
		rest     []models.Project

		f flusher.Flusher

		chunkSize int
	)
	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockProjectStorage = mocks.NewMockProjectStorage(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, &mocks.MockRepoStorage{}, mockProjectStorage)
		rest = f.FlushProjects(ctx, projects)
	})

	AfterEach(func() {
		ctrl.Finish()
	})
	Context("checking incorrect input", func() {

		BeforeEach(func() {
			chunkSize = -1
			projects = []models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
			}

			mockProjectStorage.EXPECT().MultiAddProject(ctx, gomock.Any()).Return([]uint64{}, nil).Times(0)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(projects))
		})
	})

	Context("storage saves all objects", func() {

		BeforeEach(func() {
			chunkSize = 2
			projects = []models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
			}

			mockProjectStorage.EXPECT().MultiAddProject(ctx, gomock.Len(chunkSize)).Return([]uint64{1, 2}, nil).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("storage doesn't save all objects", func() {

		BeforeEach(func() {
			chunkSize = 2
			projects = []models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
			}

			mockProjectStorage.EXPECT().MultiAddProject(
				ctx, gomock.Len(len(projects))).Return([]uint64{}, errors.New("some error")).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(projects))
		})
	})

	Context("storage doesn't save part of objects", func() {

		BeforeEach(func() {
			chunkSize = 3
			projects = []models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 2, CourseId: 2, Name: "2"},
				{Id: 3, CourseId: 3, Name: "3"},
			}

			gomock.InOrder(
				mockProjectStorage.EXPECT().MultiAddProject(ctx, gomock.Len(chunkSize)).Return([]uint64{1, 2, 3}, nil).Times(1),
				mockProjectStorage.EXPECT().MultiAddProject(
					ctx,
					gomock.Len(len(projects)-chunkSize)).Return([]uint64{}, errors.New("some error")).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo([]models.Project{
				{Id: 2, CourseId: 2, Name: "2"},
				{Id: 3, CourseId: 3, Name: "3"},
			}))
		})
	})

	Context("checking storage calls order", func() {

		BeforeEach(func() {
			chunkSize = 3
			projects = []models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
			}

			gomock.InOrder(
				mockProjectStorage.EXPECT().MultiAddProject(
					ctx, gomock.Len(chunkSize)).Return([]uint64{1, 2, 3}, nil).Times(1),
				mockProjectStorage.EXPECT().MultiAddProject(
					ctx,
					gomock.Len(len(projects)-chunkSize)).Return([]uint64{4}, nil).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})
})
