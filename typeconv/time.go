package typeconv

import (
	"time"
)

func Time(i time.Time) *time.Time {
	return &i
}

func TimeVal(i *time.Time) time.Time {
	if i == nil {
		return time.Time{}
	}
	return *i
}

func Duration(i time.Duration) *time.Duration {
	return &i
}

func DurationVal(i *time.Duration) time.Duration {
	if i == nil {
		return 0
	}
	return *i
}
