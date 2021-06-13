package saver_test

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	"github.com/ozoncp/ocp-project-api/internal/mocks"
	"github.com/ozoncp/ocp-project-api/internal/models"
	"github.com/ozoncp/ocp-project-api/internal/saver"
	"time"
)

var _ = Describe("Saver", func() {
	var (
		ctx  context.Context
		ctrl *gomock.Controller

		mockFlusher *mocks.MockFlusher
		mockAlarm   *mocks.MockAlarm

		project models.Project
		repo    models.Repo
		s       saver.Saver

		alarms chan struct{}
	)

	BeforeEach(func() {
		ctx = context.Background()
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
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			s = saver.NewSaver(ctx, 2, mockAlarm, mockFlusher, saver.CleanOne, "simple alarm imitation")

			s.SaveProject(project)
			s.SaveRepo(repo)
			mockFlusher.EXPECT().FlushProjects(ctx, gomock.Any()).Return(nil).MinTimes(1)
			mockFlusher.EXPECT().FlushRepos(ctx, gomock.Any()).Return(nil).MinTimes(1)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			cancelFunc()
			s.Close()
		})

		It("", func() {
		})
	})

	Context("clean all policy", func() {
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			s = saver.NewSaver(ctx, 2, mockAlarm, mockFlusher, saver.CleanAll, "clean all policy")

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			mockFlusher.EXPECT().FlushRepos(ctx, gomock.Len(1)).Return(nil).Times(1)
			mockFlusher.EXPECT().FlushProjects(ctx, gomock.Len(1)).Return(nil).Times(1)
			mockFlusher.EXPECT().FlushRepos(ctx, gomock.Len(0)).Return(nil).AnyTimes()
			mockFlusher.EXPECT().FlushProjects(ctx, gomock.Len(0)).Return(nil).AnyTimes()
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			cancelFunc()
			s.Close()
		})

		It("", func() {
		})
	})

	Context("clean one policy", func() {
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			s = saver.NewSaver(ctx, 2, mockAlarm, mockFlusher, saver.CleanOne, "clean one policy")

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			s.SaveProject(project)
			s.SaveRepo(repo)

			mockFlusher.EXPECT().FlushRepos(ctx, gomock.Len(2)).Return(nil).Times(1)
			mockFlusher.EXPECT().FlushProjects(ctx, gomock.Len(2)).Return(nil).Times(1)
			mockFlusher.EXPECT().FlushRepos(ctx, gomock.Len(0)).Return(nil).AnyTimes()
			mockFlusher.EXPECT().FlushProjects(ctx, gomock.Len(0)).Return(nil).AnyTimes()

		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			cancelFunc()
			s.Close()
		})

		It("", func() {
		})
	})

	Context("flushing on close without alarm", func() {
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			s = saver.NewSaver(ctx, 2, mockAlarm, mockFlusher, saver.CleanOne, "flushing on close without alarm")

			s.SaveProject(project)
			s.SaveRepo(repo)
			mockFlusher.EXPECT().FlushProjects(ctx, gomock.Any()).Return(nil).MinTimes(1)
			mockFlusher.EXPECT().FlushRepos(ctx, gomock.Any()).Return(nil).MinTimes(1)
		})

		JustBeforeEach(func() {
			cancelFunc()
			s.Close()
		})

		It("", func() {
			time.Sleep(time.Second * 2)
		})
	})

	Context("not all flushed", func() {
		var (
			cancelFunc context.CancelFunc
		)

		BeforeEach(func() {
			ctx, cancelFunc = context.WithCancel(ctx)
			s = saver.NewSaver(ctx, 2, mockAlarm, mockFlusher, saver.CleanOne, "not all flushed")
			s.SaveProject(project)
			s.SaveRepo(repo)

			mockFlusher.EXPECT().FlushRepos(ctx, gomock.Any()).Return([]models.Repo{{}}).MinTimes(1)
			mockFlusher.EXPECT().FlushProjects(ctx, gomock.Any()).Return([]models.Project{{}}).MinTimes(1)
		})

		JustBeforeEach(func() {
			alarms <- struct{}{}
		})

		AfterEach(func() {
			cancelFunc()
			s.Close()
		})

		It("", func() {
		})
	})

})
