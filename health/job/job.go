package job

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/InVisionApp/go-health/checkers"
	diskchk "github.com/InVisionApp/go-health/checkers/disk"
	"github.com/go-seidon/provider/health/checker"
)

type HealthJob struct {
	Name     string
	Checker  checker.Checker
	Interval time.Duration
}

type HttpPingParam struct {
	Name     string
	Interval time.Duration
	Url      string
}

func NewHttpPing(p HttpPingParam) (*HealthJob, error) {
	if strings.TrimSpace(p.Name) == "" {
		return nil, fmt.Errorf("invalid name")
	}

	pingUrl, err := url.Parse(p.Url)
	if err != nil {
		return nil, err
	}

	internetConnection, err := checkers.NewHTTP(&checkers.HTTPConfig{
		URL: pingUrl,
	})
	if err != nil {
		return nil, err
	}

	job := &HealthJob{
		Name:     p.Name,
		Interval: p.Interval,
		Checker:  internetConnection,
	}
	return job, nil
}

type DiskUsageParam struct {
	Name      string
	Interval  time.Duration
	Directory string
}

func NewDiskUsage(p DiskUsageParam) (*HealthJob, error) {
	if strings.TrimSpace(p.Name) == "" {
		return nil, fmt.Errorf("invalid name")
	}
	if strings.TrimSpace(p.Directory) == "" {
		return nil, fmt.Errorf("invalid directory")
	}

	appDiskChecker, err := diskchk.NewDiskUsage(&diskchk.DiskUsageConfig{
		Path:              p.Directory,
		WarningThreshold:  50,
		CriticalThreshold: 20,
	})
	if err != nil {
		return nil, err
	}

	job := &HealthJob{
		Name:     p.Name,
		Interval: p.Interval,
		Checker:  appDiskChecker,
	}
	return job, nil
}

type RepoPingParam struct {
	Name       string
	Interval   time.Duration
	DataSource checker.DataSource
}

func NewRepoPing(p RepoPingParam) (*HealthJob, error) {
	if strings.TrimSpace(p.Name) == "" {
		return nil, fmt.Errorf("invalid name")
	}
	if p.DataSource == nil {
		return nil, fmt.Errorf("invalid data source")
	}

	pingChecker := checker.NewRepoPing(p.DataSource)

	job := &HealthJob{
		Name:     p.Name,
		Interval: p.Interval,
		Checker:  pingChecker,
	}
	return job, nil
}
