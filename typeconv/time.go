package typeconv

import (
	"time"
)

func Time(i time.Time) *time.Time {
	t := i.UTC()
	return &t
}

func TimeVal(i *time.Time) time.Time {
	if i == nil {
		return time.Time{}.UTC()
	}
	return i.UTC()
}

func UnixMilli(i *int64) *time.Time {
	if i == nil {
		return nil
	}
	ms := time.UnixMilli(Int64Val(i))
	return Time(ms)
}

func TimeMilli(i *time.Time) *int64 {
	if i == nil {
		return nil
	}
	t := i.UnixMilli()
	return Int64(t)
}
