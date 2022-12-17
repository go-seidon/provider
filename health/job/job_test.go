package job_test

import (
	"fmt"
	"net/url"
	"testing"
	"time"

	"github.com/go-seidon/provider/health/job"
	mock_health "github.com/go-seidon/provider/health/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestJob(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Job Package")
}

var _ = Describe("Health Check Job", func() {

	Context("NewHttpPingJob function", Label("unit"), func() {
		var (
			p job.HttpPingParam
		)

		BeforeEach(func() {
			p = job.HttpPingParam{
				Name:     "internet-checker",
				Interval: 30 * time.Second,
				Url:      "https://google.com",
			}
		})

		When("name is invalid", func() {
			It("should return error", func() {
				p.Name = " "
				res, err := job.NewHttpPing(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid name")))
			})
		})

		When("url is invalid", func() {
			It("should return error", func() {
				p.Url = "http:// "
				res, err := job.NewHttpPing(p)

				expectedErr := &url.Error{
					Op:  "parse",
					URL: "http:// ",
					Err: url.InvalidHostError(" "),
				}
				Expect(res).To(BeNil())
				Expect(err).To(Equal(expectedErr))
			})
		})

		When("parameter are valid", func() {
			It("should return result", func() {
				res, err := job.NewHttpPing(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("NewDiskUsageJob function", Label("unit"), func() {
		var (
			p job.DiskUsageParam
		)

		BeforeEach(func() {
			p = job.DiskUsageParam{
				Name:      "app-disk",
				Interval:  60 * time.Second,
				Directory: "/usr/bin",
			}
		})

		When("name is invalid", func() {
			It("should return error", func() {
				p.Name = " "
				res, err := job.NewDiskUsage(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid name")))
			})
		})

		When("directory is invalid", func() {
			It("should return error", func() {
				p.Directory = " "
				res, err := job.NewDiskUsage(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid directory")))
			})
		})

		When("parameter are valid", func() {
			It("should return result", func() {
				res, err := job.NewDiskUsage(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

	Context("NewRepoPingJob function", Label("unit"), func() {
		var (
			p job.RepoPingParam
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			datasource := mock_health.NewMockDataSource(ctrl)

			p = job.RepoPingParam{
				Name:       "repo-ping",
				Interval:   60 * time.Second,
				DataSource: datasource,
			}
		})

		When("name is invalid", func() {
			It("should return error", func() {
				p.Name = " "
				res, err := job.NewRepoPing(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid name")))
			})
		})

		When("datasource is invalid", func() {
			It("should return error", func() {
				p.DataSource = nil
				res, err := job.NewRepoPing(p)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("invalid data source")))
			})
		})

		When("parameter are valid", func() {
			It("should return result", func() {
				res, err := job.NewRepoPing(p)

				Expect(res).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})

})
