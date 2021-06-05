package saver_test

import (
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/ozoncp/ocp-project-api/internal/mocks"
	"github.com/ozoncp/ocp-project-api/internal/models"

	"github.com/ozoncp/ocp-project-api/internal/saver"
)

const (
	capacity = 10
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

		s = saver.NewSaver(capacity, mockAlarm, mockFlusher, saver.CleanOne)
	})

	JustBeforeEach(func() {
		s.SaveProject(project)
		s.SaveRepo(repo)
	})

	AfterEach(func() {
		s.Close()
		ctrl.Finish()
	})

	Context("alarm imitation", func() {

		BeforeEach(func() {
			mockFlusher.EXPECT().FlushRepos(gomock.Any()).Return(nil).MinTimes(1).Times(1)
			mockFlusher.EXPECT().FlushProjects(gomock.Any()).Return(nil).MinTimes(1).Times(1)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		It("", func() {
			Expect(project).Should(Equal(project))
		})
	})
})
