package checker

import (
	"context"
	"time"
)

type Checker interface {
	Status() (interface{}, error)
}

type DataSource interface {
	Ping(ctx context.Context) error
}

type repoPingChecker struct {
	dataSource DataSource
}

type RepoPingResult struct {
	CheckedAt time.Time
}

func (p *repoPingChecker) Status() (interface{}, error) {
	err := p.dataSource.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	ts := time.Now()
	res := &RepoPingResult{ts}
	return res, nil
}

func NewRepoPing(dataSource DataSource) *repoPingChecker {
	return &repoPingChecker{dataSource}
}
