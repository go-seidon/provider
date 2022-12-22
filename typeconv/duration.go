package typeconv

import (
	"time"
)

func Duration(i time.Duration) *time.Duration {
	return &i
}

func DurationVal(i *time.Duration) time.Duration {
	if i == nil {
		return 0
	}
	return *i
}
