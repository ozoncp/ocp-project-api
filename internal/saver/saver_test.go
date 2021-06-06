package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-project-api/internal/mocks"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/saver"
)

var _ = Describe("Saver", func() {
	var (
		ctrl *gomock.Controller

		mockFlusher *mocks.MockFlusher
		mockAlarm   *mocks.MockAlarm

		project models.Project
		repo    models.Repo
		s       saver.Saver

		alarms chan struct{}
	)

	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockAlarm = mocks.NewMockAlarm(ctrl)
		mockFlusher = mocks.NewMockFlusher(ctrl)

		alarms = make(chan struct{})
		mockAlarm.EXPECT().Alarms().Return(alarms).AnyTimes()
	})

	AfterEach(func() {
		ctrl.Finish()
	})

	Context("simple alarm imitation", func() {

		BeforeEach(func() {
			s = saver.NewSaver(2, mockAlarm, mockFlusher, saver.CleanOne)

			s.SaveProject(project)
			s.SaveRepo(repo)
			mockFlusher.EXPECT().FlushProjects(gomock.Any()).Return(nil).MinTimes(1)
			mockFlusher.EXPECT().FlushRepos(gomock.Any()).Return(nil).MinTimes(1)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			s.Close()
		})

		It("", func() {
		})
	})

	Context("clean all policy", func() {

		BeforeEach(func() {
			s = saver.NewSaver(2, mockAlarm, mockFlusher, saver.CleanAll)

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			mockFlusher.EXPECT().FlushRepos(gomock.Len(1)).Return(nil).Times(1)
			mockFlusher.EXPECT().FlushProjects(gomock.Len(1)).Return(nil).Times(1)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			s.Close()
		})

		It("", func() {
		})
	})

	Context("clean one policy", func() {

		BeforeEach(func() {
			s = saver.NewSaver(2, mockAlarm, mockFlusher, saver.CleanOne)

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			mockFlusher.EXPECT().FlushRepos(gomock.Len(2)).Return(nil).Times(1)
			mockFlusher.EXPECT().FlushProjects(gomock.Len(2)).Return(nil).Times(1)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			s.Close()
		})

		It("", func() {
		})
	})

	Context("flushing on close without alarm", func() {

		BeforeEach(func() {
			s = saver.NewSaver(2, mockAlarm, mockFlusher, saver.CleanOne)

			s.SaveProject(project)
			s.SaveRepo(repo)
			mockFlusher.EXPECT().FlushProjects(gomock.Any()).Return(nil).MinTimes(1)
			mockFlusher.EXPECT().FlushRepos(gomock.Any()).Return(nil).MinTimes(1)
		})

		JustBeforeEach(func() {
			s.Close()
		})

		It("", func() {
		})
	})

	Context("not all flushed", func() {

		BeforeEach(func() {
			s = saver.NewSaver(2, mockAlarm, mockFlusher, saver.CleanOne)
			s.SaveProject(project)
			s.SaveRepo(repo)

			mockFlusher.EXPECT().FlushRepos(gomock.Any()).Return([]models.Repo{{}}).Times(1)
			mockFlusher.EXPECT().FlushProjects(gomock.Any()).Return([]models.Project{{}}).Times(1)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			s.Close()
		})

		It("", func() {
		})
	})

})
