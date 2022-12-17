package health

import (
	"context"
	"time"

	"github.com/InVisionApp/go-health"
	"github.com/go-seidon/provider/health/job"
	"github.com/go-seidon/provider/logging/logrus"
)

const (
	STATUS_OK      = "OK"
	STATUS_WARNING = "WARNING"
	STATUS_FAILED  = "FAILED"
)

type HealthCheck interface {
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Check(ctx context.Context) (*CheckResult, error)
}

type CheckResult struct {
	Status string
	Items  map[string]CheckResultItem
}

type CheckResultItem struct {
	Name      string
	Status    string
	Error     string
	Fatal     bool
	CheckedAt time.Time
}

type HealthClient interface {
	AddChecks(cfgs []*health.Config) error
	Start() error
	Stop() error
	State() (map[string]health.State, bool, error)
}

type healthCheck struct {
	client    HealthClient
	jobs      []*job.HealthJob
	runStatus bool
}

func (s *healthCheck) Start(ctx context.Context) error {
	if s.runStatus {
		return nil
	}

	cfgs := []*health.Config{}
	for _, job := range s.jobs {
		cfgs = append(cfgs, &health.Config{
			Name:     job.Name,
			Checker:  job.Checker,
			Interval: job.Interval,
		})
	}

	err := s.client.AddChecks(cfgs)
	if err != nil {
		return err
	}

	err = s.client.Start()
	if err != nil {
		return err
	}

	s.runStatus = true
	return nil
}

func (s *healthCheck) Stop(ctx context.Context) error {
	return s.client.Stop()
}

func (s *healthCheck) Check(ctx context.Context) (*CheckResult, error) {
	states, isFailed, err := s.client.State()
	if err != nil {
		return nil, err
	}

	res := &CheckResult{
		Status: STATUS_FAILED,
		Items:  make(map[string]CheckResultItem),
	}
	if isFailed {
		return res, nil
	}

	totalFailed := 0
	for key, state := range states {

		status := STATUS_OK
		if state.Status == "failed" {
			status = STATUS_FAILED
			totalFailed++
		}

		res.Items[key] = CheckResultItem{
			Name:      state.Name,
			Status:    status,
			Error:     state.Err,
			CheckedAt: state.CheckTime.UTC(),
		}
	}

	if totalFailed == 0 {
		res.Status = STATUS_OK
	} else if totalFailed != len(states) {
		res.Status = STATUS_WARNING
	}

	return res, nil
}

func NewHealthCheck(opts ...HealthOption) *healthCheck {
	p := HealthCheckParam{
		Jobs: []*job.HealthJob{},
	}
	for _, opt := range opts {
		opt(&p)
	}

	logger := p.Logger
	if logger == nil {
		logger = logrus.NewLogger()
	}

	client := p.Client
	if client == nil {
		h := health.New()
		h.Logger = &healthLog{client: logger}

		client = h
	}

	return &healthCheck{
		client: client,
		jobs:   p.Jobs,
	}
}
