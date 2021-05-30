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

var _ = Describe("Flusher", func() {
	var (
		ctrl *gomock.Controller

		mockStorage *mocks.MockStorage

		objects []models.Artifact
		rest    []models.Artifact

		f flusher.Flusher

		chunkSize int
	)
	BeforeEach(func() {
		ctrl = gomock.NewController(GinkgoT())

		mockStorage = mocks.NewMockStorage(ctrl)
	})

	JustBeforeEach(func() {
		f = flusher.NewFlusher(chunkSize, mockStorage)
		rest = f.Flush(objects)
	})

	AfterEach(func() {
		ctrl.Finish()
	})
	Context("checking incorrect input", func() {

		BeforeEach(func() {
			chunkSize = -1
			objects = []models.Artifact{
				models.NewProject(1, 1, "1"),
				models.NewProject(1, 1, "1"),
			}

			mockStorage.EXPECT().AddObjects(gomock.Any()).Return(nil).Times(0)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("storage saves all objects", func() {

		BeforeEach(func() {
			chunkSize = 2
			objects = []models.Artifact{
				models.NewProject(1, 1, "1"),
				models.NewProject(1, 1, "1"),
			}

			mockStorage.EXPECT().AddObjects(gomock.Len(chunkSize)).Return(nil).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})

	Context("storage doesn't save all objects", func() {

		BeforeEach(func() {
			chunkSize = 2
			objects = []models.Artifact{
				models.NewProject(1, 1, "1"),
			}

			mockStorage.EXPECT().AddObjects(gomock.Len(len(objects))).Return(errors.New("some error")).Times(1)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo(objects))
		})
	})

	Context("storage doesn't save part of objects", func() {

		BeforeEach(func() {
			chunkSize = 3
			objects = []models.Artifact{
				models.NewProject(1, 1, "1"),
				models.NewProject(1, 1, "1"),
				models.NewProject(1, 1, "1"),
				models.NewProject(2, 2, "2"),
				models.NewProject(3, 3, "3"),
			}

			gomock.InOrder(
				mockStorage.EXPECT().AddObjects(gomock.Len(chunkSize)).Return(nil).Times(1),
				mockStorage.EXPECT().AddObjects(
					gomock.Len(len(objects)-chunkSize)).Return(errors.New("some error")).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeEquivalentTo([]models.Artifact{
				models.NewProject(2, 2, "2"),
				models.NewProject(3, 3, "3"),
			}))
		})
	})

	Context("checking storage calls order", func() {

		BeforeEach(func() {
			chunkSize = 3
			objects = []models.Artifact{
				models.NewProject(1, 1, "1"),
				models.NewProject(1, 1, "1"),
				models.NewProject(1, 1, "1"),
				models.NewProject(1, 1, "1"),
			}

			gomock.InOrder(
				mockStorage.EXPECT().AddObjects(gomock.Len(chunkSize)).Return(nil).Times(1),
				mockStorage.EXPECT().AddObjects(gomock.Len(len(objects)-chunkSize)).Return(nil).Times(1),
			)
		})

		It("", func() {
			Expect(rest).Should(BeNil())
		})
	})
})
