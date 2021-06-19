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
	)

	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockRepoStorage = mocks.NewMockRepoStorage(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(mockRepoStorage, &mocks.MockProjectStorage{})
		rest = f.FlushRepos(ctx, repos)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("storage saves all objects", func() {

		BeforeEach(func() {
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Len(len(repos))).Return([]uint64{1, 2}, nil).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("storage didn't save all objects", func() {

		BeforeEach(func() {
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
			}

			mockRepoStorage.EXPECT().MultiAddRepo(ctx, gomock.Len(len(repos))).Return([]uint64{}, errors.New("some error")).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(repos))
		})
	})

	Context("storage didn't save part of objects", func() {

		BeforeEach(func() {
			repos = []models.Repo{
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 1, ProjectId: 1, UserId: 1, Link: "1"},
				{Id: 4, ProjectId: 2, UserId: 2, Link: "1"},
				{Id: 5, ProjectId: 3, UserId: 3, Link: "1"},
			}

			mockRepoStorage.EXPECT().
				MultiAddRepo(ctx, gomock.Len(len(repos))).
				Return([]uint64{1, 2, 3}, errors.New("some error")).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo([]models.Repo{
				{Id: 4, ProjectId: 2, UserId: 2, Link: "1"},
				{Id: 5, ProjectId: 3, UserId: 3, Link: "1"},
			}))
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
	)
	BeforeEach(func() {
		ctx = context.Background()
		ctrl = gomock.NewController(GinkgoT())

		mockProjectStorage = mocks.NewMockProjectStorage(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(&mocks.MockRepoStorage{}, mockProjectStorage)
		rest = f.FlushProjects(ctx, projects)
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("storage saves all objects", func() {

		BeforeEach(func() {
			projects = []models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
			}

			mockProjectStorage.EXPECT().MultiAddProject(ctx, gomock.Len(len(projects))).Return([]uint64{1, 2}, nil).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("storage didn't save all objects", func() {

		BeforeEach(func() {
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

	Context("storage didn't save part of objects", func() {

		BeforeEach(func() {
			projects = []models.Project{
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 1, CourseId: 1, Name: "1"},
				{Id: 4, CourseId: 2, Name: "2"},
				{Id: 5, CourseId: 3, Name: "3"},
			}

			mockProjectStorage.EXPECT().
				MultiAddProject(ctx, gomock.Len(len(projects))).
				Return([]uint64{1, 2, 3}, errors.New("some error")).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo([]models.Project{
				{Id: 4, CourseId: 2, Name: "2"},
				{Id: 5, CourseId: 3, Name: "3"},
			}))
		})
	})
})
