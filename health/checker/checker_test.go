package checker_test

import (
	"fmt"
	"testing"

	"github.com/go-seidon/provider/health/checker"
	mock_health "github.com/go-seidon/provider/health/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestChecker(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Checker Package")
}

var _ = Describe("Health Check Checker", func() {

	Context("RepoPing Checker", Label("unit"), func() {
		var (
			check      checker.Checker
			dataSource *mock_health.MockDataSource
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			dataSource = mock_health.NewMockDataSource(ctrl)
			check = checker.NewRepoPing(dataSource)
		})

		When("failed ping repository", func() {
			It("should return error", func() {
				dataSource.
					EXPECT().
					Ping(gomock.Any()).
					Return(fmt.Errorf("db error")).
					Times(1)

				res, err := check.Status()

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("db error")))
			})
		})

		When("success ping repository", func() {
			It("should return result", func() {
				dataSource.
					EXPECT().
					Ping(gomock.Any()).
					Return(nil).
					Times(1)

				res, err := check.Status()

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

})
