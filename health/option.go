package health

import (
	"github.com/go-seidon/provider/health/job"
	"github.com/go-seidon/provider/logging"
)

type HealthCheckParam struct {
	Jobs   []*job.HealthJob
	Logger logging.Logger
	Client HealthClient
}

type HealthOption func(*HealthCheckParam)

func WithLogger(logger logging.Logger) HealthOption {
	return func(p *HealthCheckParam) {
		p.Logger = logger
	}
}

func AddJob(job *job.HealthJob) HealthOption {
	return func(p *HealthCheckParam) {
		p.Jobs = append(p.Jobs, job)
	}
}

func WithJobs(jobs []*job.HealthJob) HealthOption {
	return func(p *HealthCheckParam) {
		p.Jobs = jobs
	}
}

func WithClient(client HealthClient) HealthOption {
	return func(p *HealthCheckParam) {
		p.Client = client
	}
}
