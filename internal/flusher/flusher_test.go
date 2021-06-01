package flusher_test

import (
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
		ctrl *gomock.Controller

		mockRepoStorage *mocks.MockRepoStorage

		repos []models.Repo
		rest  []models.Repo

		f flusher.Flusher

		chunkSize int
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockRepoStorage = mocks.NewMockRepoStorage(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockRepoStorage, &mocks.MockProjectStorage{})
		rest = f.FlushRepos(repos)
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

			mockRepoStorage.EXPECT().AddRepos(gomock.Any()).Return(nil).Times(0)
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

			mockRepoStorage.EXPECT().AddRepos(gomock.Len(chunkSize)).Return(nil).Times(1)
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

			mockRepoStorage.EXPECT().AddRepos(gomock.Len(len(repos))).Return(errors.New("some error")).Times(1)
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
				mockRepoStorage.EXPECT().AddRepos(gomock.Len(chunkSize)).Return(nil).Times(1),
				mockRepoStorage.EXPECT().AddRepos(
					gomock.Len(len(repos)-chunkSize)).Return(errors.New("some error")).Times(1),
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
				mockRepoStorage.EXPECT().AddRepos(gomock.Len(chunkSize)).Return(nil).Times(1),
				mockRepoStorage.EXPECT().AddRepos(gomock.Len(len(repos)-chunkSize)).Return(nil).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})
})

var _ = Describe("Flush into ProjectStorage", func() {
	var (
		ctrl *gomock.Controller

		mockProjectStorage *mocks.MockProjectStorage

		projects []models.Project
		rest     []models.Project

		f flusher.Flusher

		chunkSize int
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockProjectStorage = mocks.NewMockProjectStorage(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, &mocks.MockRepoStorage{}, mockProjectStorage)
		rest = f.FlushProjects(projects)
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

			mockProjectStorage.EXPECT().AddProjects(gomock.Any()).Return(nil).Times(0)
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

			mockProjectStorage.EXPECT().AddProjects(gomock.Len(chunkSize)).Return(nil).Times(1)
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

			mockProjectStorage.EXPECT().AddProjects(gomock.Len(len(projects))).Return(errors.New("some error")).Times(1)
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
				mockProjectStorage.EXPECT().AddProjects(gomock.Len(chunkSize)).Return(nil).Times(1),
				mockProjectStorage.EXPECT().AddProjects(
					gomock.Len(len(projects)-chunkSize)).Return(errors.New("some error")).Times(1),
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
				mockProjectStorage.EXPECT().AddProjects(gomock.Len(chunkSize)).Return(nil).Times(1),
				mockProjectStorage.EXPECT().AddProjects(gomock.Len(len(projects)-chunkSize)).Return(nil).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})
})
