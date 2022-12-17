package health_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	gohealth "github.com/InVisionApp/go-health"
	"github.com/go-seidon/provider/health"
	"github.com/go-seidon/provider/health/job"
	mock_health "github.com/go-seidon/provider/health/mock"
	mock_logging "github.com/go-seidon/provider/logging/mock"
	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestHealthCheck(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HealthCheck Package")
}

var _ = Describe("Health Check Job", func() {

	Context("WithLogger function", Label("unit"), func() {
		var (
			logger *mock_logging.MockLogger
		)

		BeforeEach(func() {
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			logger = mock_logging.NewMockLogger(ctrl)
		})

		When("logger is specified", func() {
			It("should append logger", func() {
				opt := health.WithLogger(logger)

				var option health.HealthCheckParam
				opt(&option)

				Expect(option.Logger).To(Equal(logger))
				Expect(option.Jobs).To(BeNil())
			})
		})
	})

	Context("AddJob function", Label("unit"), func() {
		var (
			j *job.HealthJob
		)

		BeforeEach(func() {
			j = &job.HealthJob{
				Name:     "mock-name",
				Interval: 1 * time.Second,
			}
		})

		When("jobs are empty", func() {
			It("should append job", func() {
				opt := health.AddJob(j)

				var option health.HealthCheckParam
				opt(&option)

				Expect(option.Logger).To(BeNil())
				Expect(len(option.Jobs)).To(Equal(1))
			})
		})

		When("jobs are not empty", func() {
			It("should append job", func() {
				opt := health.AddJob(j)

				var option health.HealthCheckParam
				option.Jobs = append(option.Jobs, j)
				opt(&option)

				Expect(option.Logger).To(BeNil())
				Expect(len(option.Jobs)).To(Equal(2))
			})
		})
	})

	Context("WithJobs function", Label("unit"), func() {
		var (
			j    *job.HealthJob
			jobs []*job.HealthJob
		)

		BeforeEach(func() {
			j = &job.HealthJob{
				Name:     "mock-name",
				Interval: 1 * time.Second,
			}
			jobs = []*job.HealthJob{j, j}
		})

		When("jobs are empty", func() {
			It("should add jobs", func() {
				opt := health.WithJobs(jobs)

				var option health.HealthCheckParam

				opt(&option)

				Expect(option.Logger).To(BeNil())
				Expect(len(option.Jobs)).To(Equal(2))
			})
		})

		When("jobs are not empty", func() {
			It("should replace jobs", func() {
				opt := health.WithJobs(jobs)

				var option health.HealthCheckParam
				option.Jobs = []*job.HealthJob{j}

				opt(&option)

				Expect(option.Logger).To(BeNil())
				Expect(len(option.Jobs)).To(Equal(2))
			})
		})
	})

	Context("Start function", Label("unit"), func() {
		var (
			ctx    context.Context
			client *mock_health.MockHealthClient
			s      health.HealthCheck
			logger *mock_logging.MockLogger
		)

		BeforeEach(func() {
			ctx = context.Background()
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			client = mock_health.NewMockHealthClient(ctrl)
			jobs := []*job.HealthJob{
				{
					Name:     "mock-job",
					Checker:  nil,
					Interval: 1,
				},
			}
			logger = mock_logging.NewMockLogger(ctrl)
			s = health.NewHealthCheck(
				health.WithJobs(jobs),
				health.WithLogger(logger),
				health.WithClient(client),
			)
		})

		When("failed add checkers", func() {
			It("should return error", func() {
				client.
					EXPECT().
					AddChecks(gomock.Any()).
					Return(fmt.Errorf("failed add checkers")).
					Times(1)

				err := s.Start(ctx)

				Expect(err).To(Equal(fmt.Errorf("failed add checkers")))
			})
		})

		When("failed start app", func() {
			It("should return error", func() {
				client.
					EXPECT().
					AddChecks(gomock.Any()).
					Return(nil).
					Times(1)

				client.
					EXPECT().
					Start().
					Return(fmt.Errorf("failed start app")).
					Times(1)

				err := s.Start(ctx)

				Expect(err).To(Equal(fmt.Errorf("failed start app")))
			})
		})

		When("success start app", func() {
			It("should return result", func() {
				client.
					EXPECT().
					AddChecks(gomock.Any()).
					Return(nil).
					Times(1)

				client.
					EXPECT().
					Start().
					Return(nil).
					Times(1)

				err := s.Start(ctx)

				Expect(err).To(BeNil())
			})
		})

		When("app is already started", func() {
			It("should return result", func() {
				client.
					EXPECT().
					AddChecks(gomock.Any()).
					Return(nil).
					Times(1)

				client.
					EXPECT().
					Start().
					Return(nil).
					Times(1)

				err1 := s.Start(ctx)
				err2 := s.Start(ctx)

				Expect(err1).To(BeNil())
				Expect(err2).To(BeNil())
			})
		})
	})

	Context("Stop function", Label("unit"), func() {
		var (
			ctx    context.Context
			client *mock_health.MockHealthClient
			s      health.HealthCheck
			logger *mock_logging.MockLogger
		)

		BeforeEach(func() {
			ctx = context.Background()
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			client = mock_health.NewMockHealthClient(ctrl)
			jobs := []*job.HealthJob{
				{
					Name:     "mock-job",
					Checker:  nil,
					Interval: 1,
				},
			}
			logger = mock_logging.NewMockLogger(ctrl)
			s = health.NewHealthCheck(
				health.WithJobs(jobs),
				health.WithLogger(logger),
				health.WithClient(client),
			)
		})

		When("failed stop app", func() {
			It("should return error", func() {
				client.
					EXPECT().
					Stop().
					Return(fmt.Errorf("failed stop app")).
					Times(1)

				err := s.Stop(ctx)

				Expect(err).To(Equal(fmt.Errorf("failed stop app")))
			})
		})

		When("success stop app", func() {
			It("should return result", func() {
				client.
					EXPECT().
					Stop().
					Return(nil).
					Times(1)

				err := s.Stop(ctx)

				Expect(err).To(BeNil())
			})
		})
	})

	Context("Check function", Label("unit"), func() {
		var (
			ctx              context.Context
			client           *mock_health.MockHealthClient
			s                health.HealthCheck
			currentTimestamp time.Time
			logger           *mock_logging.MockLogger
		)

		BeforeEach(func() {
			ctx = context.Background()
			t := GinkgoT()
			ctrl := gomock.NewController(t)
			client = mock_health.NewMockHealthClient(ctrl)
			jobs := []*job.HealthJob{
				{
					Name:     "mock-job",
					Checker:  nil,
					Interval: 1,
				},
			}
			logger = mock_logging.NewMockLogger(ctrl)
			s = health.NewHealthCheck(
				health.WithJobs(jobs),
				health.WithLogger(logger),
				health.WithClient(client),
			)
			currentTimestamp = time.Now()
		})

		When("error occured", func() {
			It("should return error", func() {
				client.
					EXPECT().
					State().
					Return(nil, true, fmt.Errorf("network error")).
					Times(1)

				res, err := s.Check(ctx)

				Expect(res).To(BeNil())
				Expect(err).To(Equal(fmt.Errorf("network error")))
			})
		})

		When("failed get state", func() {
			It("should return result", func() {
				client.
					EXPECT().
					State().
					Return(nil, true, nil).
					Times(1)

				res, err := s.Check(ctx)

				expected := &health.CheckResult{
					Status: "FAILED",
					Items:  make(map[string]health.CheckResultItem),
				}
				Expect(res).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})

		When("all check is ok", func() {
			It("should return result", func() {
				states := map[string]gohealth.State{
					"mock-job": {
						Name:      "mock-job",
						Status:    "ok",
						Err:       "",
						Fatal:     false,
						Details:   nil,
						CheckTime: currentTimestamp,
					},
				}

				client.
					EXPECT().
					State().
					Return(states, false, nil).
					Times(1)

				res, err := s.Check(ctx)

				expected := &health.CheckResult{
					Status: "OK",
					Items: map[string]health.CheckResultItem{
						"mock-job": {
							Name:      "mock-job",
							Status:    "OK",
							Error:     "",
							Fatal:     false,
							CheckedAt: currentTimestamp.UTC(),
						},
					},
				}
				Expect(res).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})

		When("all check is failed", func() {
			It("should return result", func() {
				states := map[string]gohealth.State{
					"mock-job": {
						Name:      "mock-job",
						Status:    "failed",
						Err:       "some error",
						Fatal:     false,
						Details:   nil,
						CheckTime: currentTimestamp,
					},
				}

				client.
					EXPECT().
					State().
					Return(states, false, nil).
					Times(1)

				res, err := s.Check(ctx)

				expected := &health.CheckResult{
					Status: "FAILED",
					Items: map[string]health.CheckResultItem{
						"mock-job": {
							Name:      "mock-job",
							Status:    "FAILED",
							Error:     "some error",
							Fatal:     false,
							CheckedAt: currentTimestamp.UTC(),
						},
					},
				}
				Expect(res).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})

		When("some check is failed", func() {
			It("should return result", func() {
				states := map[string]gohealth.State{
					"mock-job": {
						Name:      "mock-job",
						Status:    "failed",
						Err:       "some error",
						Fatal:     false,
						Details:   nil,
						CheckTime: currentTimestamp,
					},
					"mock-job-2": {
						Name:      "mock-job-2",
						Status:    "ok",
						Err:       "",
						Fatal:     false,
						Details:   nil,
						CheckTime: currentTimestamp,
					},
				}

				client.
					EXPECT().
					State().
					Return(states, false, nil).
					Times(1)

				res, err := s.Check(ctx)

				expected := &health.CheckResult{
					Status: "WARNING",
					Items: map[string]health.CheckResultItem{
						"mock-job": {
							Name:      "mock-job",
							Status:    "FAILED",
							Error:     "some error",
							Fatal:     false,
							CheckedAt: currentTimestamp.UTC(),
						},
						"mock-job-2": {
							Name:      "mock-job-2",
							Status:    "OK",
							Error:     "",
							Fatal:     false,
							CheckedAt: currentTimestamp.UTC(),
						},
					},
				}
				Expect(res).To(Equal(expected))
				Expect(err).To(BeNil())
			})
		})
	})

})
